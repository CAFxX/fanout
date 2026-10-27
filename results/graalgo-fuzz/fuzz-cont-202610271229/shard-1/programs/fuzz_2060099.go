package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint16; F2 string }
type GenS5 struct { F0 int8; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (169 != 203)
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 239
    var v5 uint8 = (v4 - (v4 & 42))
    var v6 float32 = 1.5
    delete(v1, 1)
    fmt.Println(v5, v4, v3)
    for fi7 := 0; fi7 < 16; fi7++ {
        if (-1 < 383) {
            switch "a" {
            case "a":
                v8 := []int{(-(fi7)), (fi7 >> ((518) & 63)), -1, fi7}
                _ = v8
            case "\n\t":
                var v9 int = fi7
                _ = v9
            }
        } else {
            if v2 {
                v5 -= 42
                v3 += 511
            } else {
                _ = len(v1)
                fmt.Println(v4, v4, v2)
            }
            var v10 int8 = -1
            v11 := []int{fi7, (^(fi7))}
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    v12 := GenS1{F0: -1, F1: -574, F2: 364}
    v13 := GenS2{F0: "hello", F1: 0, F2: (v5 >> ((42) & 63))}
    if (true || ("foo bar" >= "")) {
        delete(v1, 65535)
    }
    var v14 int = 128
    fmt.Println(v4, v6, v6)
    v15 := GenS1{F0: 1, F1: 32767, F2: 202}
    if v2 {
        if (131 < 32767) {
            fmt.Println(v5, v2)
        }
        fmt.Println(v2, v5)
        for fi16 := 0; fi16 < 9; fi16++ {
            v4 += 127
            _ = fi16
        }
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 int8 = -128
            switch (v14 & -9223372036854775808) {
            case 737:
                if (v2 || v2) {
                    delete(v1, 0)
                    for fi19 := 0; fi19 < 7; fi19++ {
                        var v20 int64 = (v3 * (v3 / (((^(v3))) | 1)))
                        fmt.Println(v18, v2)
                        if v2 {
                            v6 = (0.1 + 0.1)
                            v21 := GenS1{F0: -716, F1: int(v18), F2: 1}
                            _ = v21
                        } else {
                            switch (fi19 * 32768) {
                            case 32768:
                                v22 := []int{(v14 / ((fi19) | 1))}
                                _ = v22
                            }
                            var v23 int16 = 32767
                            v24 := GenS2{F0: "foo bar", F1: 0, F2: v4}
                            _ = v23
                            _ = v24
                        }
                        _ = fi19
                        _ = v20
                    }
                } else {
                    fmt.Println(v4, v6, v2)
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 int8 = v18
                        switch "" {
                        case "":
                            var v27 uint = 685
                            _ = v27
                        case "hello":
                            v5 = v4
                        case "\n\t":
                            switch -478 {
                            case -699:
                                v28 := []int{}
                                _ = v28
                            case -2147483648:
                                switch (fi17 ^ (^(fi25))) {
                                case 4294967295:
                                    v5 -= (^(v4))
                                case 2147483647:
                                    if (v5 == 127) {
                                        v29 := []int{fi25, (fi25 / (((fi17 >> ((255) & 63))) | 1)), fi17, -9223372036854775808}
                                        fmt.Println(fi25, fi25)
                                        delete(v1, 2147483647)
                                        _ = v29
                                    } else {
                                        if (!(827 > -2147483648)) {
                                            delete(v1, 2)
                                            switch (-(fi17)) {
                                            case -1:
                                                v26 = (v26 | 70)
                                            case -2147483648:
                                                var v30 uint16 = 32767
                                                _ = v30
                                            default:
                                                v31 := v1[749]
                                                _ = v31
                                            }
                                            if true {
                                                v6 = v6
                                            } else {
                                                v32 := GenS3{F0: 9223372036854775807}
                                                v3 -= (v3 - (v3 + (v3 >> ((0) & 63))))
                                                _ = v32
                                            }
                                        }
                                        switch (("a" + ("a" + "foo bar")) + ("hello" + "\n\t")) {
                                        case "foo bar":
                                            fmt.Println(v26)
                                        default:
                                            fmt.Println(v18)
                                        }
                                    }
                                default:
                                    fmt.Println(v6, fi25)
                                }
                            case 32767:
                                if ((-2.25 + (float32(0.0) + float32(0.1))) <= (0.0 * v6)) {
                                    var v33 uint16 = 42
                                    var v34 int = 32768
                                    _ = v33
                                    _ = v34
                                } else {
                                    v14 = -848
                                    v35 := GenS4{F0: 348, F1: 2, F2: "foo bar"}
                                    _ = v35
                                }
                            default:
                                v36 := GenS5{F0: -128, F1: 685, F2: 2147483647}
                                _ = v36
                            }
                        }
                        v37 := GenS1{F0: 42, F1: (fi25 / ((fi25) | 1)), F2: 42}
                        _ = fi25
                        _ = v26
                        _ = v37
                    }
                }
            case 9223372036854775807:
                for fi38 := 0; fi38 < 12; fi38++ {
                    v18 += (v18 >> ((18446744073709551615) & 63))
                    if v2 {
                        fmt.Println(v3)
                        var v39 uint32 = 0
                        _ = v39
                    }
                    v3 -= (v3 / ((int64(v5)) | 1))
                    if (2147483648 == 0) {
                        var v40 int8 = (v18 << ((13) & 63))
                        v41 := GenS3{F0: (v3 << ((4294967295) & 63))}
                        _ = v40
                        _ = v41
                    } else {
                        if v2 {
                            fmt.Println(v4, fi17)
                        }
                    }
                    v18 = (v18 + v18)
                    v42 := []int{fi17, 4294967295, 691, 4294967295}
                    var v43 bool = (19 != 0)
                    _ = fi38
                    _ = v42
                    _ = v43
                }
            }
            _ = fi17
            _ = v18
        }
        if v2 {
            var v44 float64 = ((-2.25 + -2.25) * ((((1e10 * ((((0.1 * 3.14159) - 0.1) + 1.5) / 1.0)) - 1.5) * (-2.25 - 0.1)) * 0.0))
            var v45 float64 = float64(float32(1e10))
            v46 := GenS6{F0: 47, F1: (v14 - v14)}
            _ = v44
            _ = v45
            _ = v46
        }
        switch ("" + (("\n\t" + "") + ("hello" + "hello"))) {
        case "foo bar":
            switch "" {
            case "a":
                var v47 int32 = -1
                _ = v47
            case "\n\t":
                if v2 {
                    var v48 float32 = v6
                    _ = v48
                }
            case "":
                var v49 int16 = -533
                _ = v49
            }
        case "a":
            var v50 uint16 = 979
            _ = v50
        }
    }
    if (4294967295 == 4294967295) {
        v51 := GenS7{F0: -128}
        _ = v51
    } else {
        var v52 int64 = v3
        switch (v14 & 127) {
        case 9223372036854775807:
            v53 := v1[433]
            _ = v53
        }
        if true {
            v54 := GenS1{F0: 268, F1: 4294967295, F2: 19}
            _ = v54
        } else {
            var v55 string = "foo bar"
            _ = v55
        }
        _ = v52
    }
    fmt.Println(v6, v4)
    switch ("\n\t" + ("foo bar" + "\n\t")) {
    case "foo bar":
        v1[4294967295] = 305
    }
    switch (v14 - (v14 + 128)) {
    case 4294967295:
        var v56 int8 = 127
        _ = v56
    default:
        for fi57 := 0; fi57 < 1; fi57++ {
            v6 = 3.14159
            v14 = (v14 | 128)
            for fi58 := 0; fi58 < 8; fi58++ {
                if (186 > (^(v4))) {
                    v59 := GenS1{F0: -2147483648, F1: fi57, F2: 476}
                    var v60 int = (fi58 & (-(fi58)))
                    _ = v59
                    _ = v60
                }
                v61 := GenS7{F0: 1}
                var v62 uint64 = 452
                _ = fi58
                _ = v61
                _ = v62
            }
            _ = fi57
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
