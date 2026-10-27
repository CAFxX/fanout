package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]string)
    v3 := []string{(v1 + (v1 + v1)), v1, ""}
    v4 := []int{128, -1, 32767, 838}
    var v5 uint8 = 1
    v1 = v1
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 string = "a"
        switch 32767 {
        case 1:
            v5 = v5
        }
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: ((v5 - 128) == 239), F1: 9223372036854775807, F2: 128}
    if (0 != v5) {
        v5 += 2
        fmt.Println(v1)
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: false, F1: -1, F2: 661}
        if true {
            switch v1 {
            case "":
                for fi11 := 0; fi11 < 20; fi11++ {
                    fmt.Println(fi9)
                    fmt.Println(fi9)
                    delete(v2, 467)
                    _ = fi11
                }
            case "hello":
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 bool = (false && false)
                    _ = fi12
                    _ = v13
                }
            case "\n\t":
                var v14 uint32 = 2147483648
                _ = v14
            default:
                var v15 uint = 474
                _ = v15
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                v1 = v1
                _ = len(v2)
                _ = len(v2)
                _ = fi16
            }
        } else {
            v3 = append(v3, (v1 + "\n\t"))
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS2{F0: 42}
                if true {
                    if false {
                        v1 = (("\n\t" + "foo bar") + (((("" + "foo bar") + ("hello" + (v1 + "a"))) + ((v1 + ("" + v1)) + "\n\t")) + ("" + (((v1 + "") + v1) + v1))))
                        var v19 float64 = float64(3.14159)
                        _ = v19
                    } else {
                        if true {
                            var v20 int32 = 42
                            switch 128 {
                            case 154:
                                v20 -= 1
                            case 4294967295:
                                if len(v4) > 0 {
                                    v4[0] = 756
                                }
                            case -1:
                                if ((true || (4294967295 > 2147483648)) || true) {
                                    v21 := GenS1{F0: true, F1: -9223372036854775808, F2: (fi17 / ((fi9) | 1))}
                                    _ = v21
                                } else {
                                    var v22 bool = (!(0.0 < (3.14159 - 0.0)))
                                    if (0 <= 2) {
                                        switch "foo bar" {
                                        case "hello":
                                            var v23 string = ("hello" + (v1 + v1))
                                            _ = v23
                                        case "\n\t":
                                            var v24 uint8 = (^(v5))
                                            _ = v24
                                        case "":
                                            _ = len(v2)
                                        default:
                                            if true {
                                                var v25 uint8 = 100
                                                v26 := GenS3{F0: (v20 / (((v20 + (v20 + v20))) | 1)), F1: v1}
                                                if (!false) {
                                                    fmt.Println(v25, v25)
                                                    if v22 {
                                                        switch v1 {
                                                        case "":
                                                            var v27 string = (("\n\t" + v1) + (v1 + v1))
                                                            _ = v27
                                                        case "\n\t":
                                                            var v28 int32 = (v20 * v20)
                                                            _ = v28
                                                        case "a":
                                                            v29 := GenS2{F0: (^(v25))}
                                                            _ = v29
                                                        }
                                                        v30 := GenS4{F0: v5, F1: -243}
                                                        _ = v30
                                                    }
                                                    switch fi9 {
                                                    case 42:
                                                        delete(v2, 2147483648)
                                                    case -9223372036854775808:
                                                        if (("" + v1) < ("" + v1)) {
                                                            v31 := GenS4{F0: v25, F1: (fi17 * fi17)}
                                                            var v32 uint64 = 4294967295
                                                            var v33 float32 = (0.1 + ((((1.5 / 1.0) / 1e10) / 3.14159) / 1.0))
                                                            _ = v31
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                    default:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                } else {
                                                    var v34 float32 = ((3.14159 / 0.1) - float32(3.14159))
                                                    v35 := GenS3{F0: v20, F1: v1}
                                                    v1 = v1
                                                    _ = v34
                                                    _ = v35
                                                }
                                                _ = v25
                                                _ = v26
                                            } else {
                                                var v36 bool = (0.1 <= 1.5)
                                                v37 := GenS2{F0: v5}
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                    } else {
                                        v22 = v22
                                        v38 := v2[65535]
                                        var v39 uint8 = (v5 ^ 159)
                                        _ = v38
                                        _ = v39
                                    }
                                    if (!(-2.25 != 1.5)) {
                                        var v40 uint32 = 805
                                        var v41 bool = v22
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v42 := GenS1{F0: v22, F1: 42, F2: (fi9 % ((fi9) | 1))}
                                        v1 = "foo bar"
                                        fmt.Println(v22, v1)
                                        _ = v42
                                    }
                                    _ = v22
                                }
                            default:
                                if (true || (!(-32768 < -589))) {
                                    var v43 string = v1
                                    v20 = (v20 | int32(v5))
                                    v44 := GenS2{F0: v5}
                                    _ = v43
                                    _ = v44
                                }
                                fmt.Println(v1)
                                v1 = "hello"
                            }
                            _ = v20
                        }
                        if true {
                            v5 -= 226
                        } else {
                            v1 = (v1 + v1)
                            v5 = 103
                            v5 = (^(v5))
                            v1 = v1
                            v1 = v1
                        }
                        if (1e10 != float32(((((0.1 / 1e10) - 1.5) - 0.0) - 1.5))) {
                            var v45 int = fi17
                            v2[4294967295] = "hello"
                            _ = v45
                        }
                        v5 = v5
                    }
                }
                v5 += 37
                _ = fi17
                _ = v18
            }
            v5 += (v5 >> ((4294967295) & 63))
        }
        switch fi9 {
        case 1:
            for fi46 := 0; fi46 < 14; fi46++ {
                switch 4294967295 {
                case 9223372036854775807:
                    var v47 uint8 = 128
                    _ = v47
                case 1:
                    var v48 float32 = ((3.14159 * 0.0) / -2.25)
                    var v49 int8 = -1
                    _ = v48
                    _ = v49
                case 0:
                    v5 = (v5 % ((127) | 1))
                default:
                    v50 := v2[311]
                    _ = v50
                }
                _ = len(v3)
                var v51 int32 = -1
                _ = fi46
                _ = v51
            }
        case 127:
            var v52 int32 = 1
            _ = v52
        default:
            if (int64(fi9) <= -279) {
                for fi53 := 0; fi53 < 10; fi53++ {
                    fmt.Println(fi53, v5, fi53)
                    _ = fi53
                }
                if true {
                    var v54 string = (v1 + "hello")
                    v5 = 2
                    var v55 float64 = (0.0 * (3.14159 * 0.0))
                    _ = v54
                    _ = v55
                }
            } else {
                var v56 int64 = 253
                var v57 int64 = v56
                var v58 int64 = v57
                _ = v56
                _ = v57
                _ = v58
            }
        }
        _ = fi9
        _ = v10
    }
    v2[42] = "hello"
    var v59 float32 = (((((((1.5 * -2.25) - float32(0.0)) + -2.25) - -2.25) - 1e10) / (1e10 * 1e10)) / 3.14159)
    var v60 float32 = v59
    for fi61 := 0; fi61 < 12; fi61++ {
        var v62 float64 = 1.5
        _ = fi61
        _ = v62
    }
    var v63 int8 = -1
    var v64 uint16 = 684
    v1 = (v1 + v1)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v63
    _ = v64
    _ = v8
}
