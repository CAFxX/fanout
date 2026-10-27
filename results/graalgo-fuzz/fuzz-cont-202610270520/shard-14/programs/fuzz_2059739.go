package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := []uint32{1, 2147483648, 65535, 42}
    var v5 uint32 = 2147483647
    var v6 bool = false
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(fi7, v5, v6)
        _ = fi7
    }
    v8 := GenS1{F0: 255, F1: 255}
    if ((v1 + v1) >= (v1 - 0)) {
        var v9 int64 = -9223372036854775808
        var v10 int = 127
        _ = v10
        _ = v9
    } else {
        v5 = (^(v5))
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                if false {
                    var v13 uint64 = 42
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v13
                }
                _ = fi12
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                fmt.Println(fi11)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi14
            }
            switch (fi11 | (fi11 ^ 32767)) {
            case -620:
                switch ("" + "hello") {
                case "\n\t":
                    fmt.Println(fi11)
                case "hello":
                    if v6 {
                        if false {
                            var v15 int16 = -727
                            for fi16 := 0; fi16 < 5; fi16++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v5 += 558
                                var v17 int64 = 324
                                _ = fi16
                                _ = v17
                            }
                            _ = v15
                        }
                        var v18 int16 = 32767
                        for fi19 := 0; fi19 < 3; fi19++ {
                            var v20 uint16 = 42
                            var v21 uint8 = v1
                            v22 := v2[(-(fi19))]
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        _ = v18
                    }
                case "foo bar":
                    for fi23 := 0; fi23 < 3; fi23++ {
                        var v24 int32 = 1
                        fmt.Println(fi23, v1, v24)
                        fmt.Println(v6, v1, v5)
                        var v25 float32 = 0.1
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
            case -171:
                v1 = uint8(v5)
            case 9223372036854775807:
                switch (("a" + (("" + "hello") + ("\n\t" + "a"))) + "") {
                case "a":
                    fmt.Println(v1, v5)
                }
            default:
                v2[(fi11 ^ fi11)] = ("hello" + "")
            }
            _ = fi11
        }
    }
    var v26 int = -2147483648
    fmt.Println(v1, v5)
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 int8 = -67
        _ = fi27
        _ = v28
    }
    v5 = (v5 * v5)
    if ((508 == -848) && (!true)) {
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := v2[fi29]
            _ = fi29
            _ = v30
        }
        switch -9223372036854775808 {
        case 0:
            switch v26 {
            case -938:
                v4 = append(v4, 2)
            }
        case -125:
            v31 := GenS1{F0: 65535, F1: 220}
            _ = v31
        }
    } else {
        fmt.Println(v5, v6, v1)
        var v32 int64 = 9223372036854775807
        _ = v32
    }
    var v33 string = "\n\t"
    fmt.Println(v6, v6, v5)
    if (18446744073709551615 < 0) {
        v33 = (v33 + v33)
        var v34 int = v26
        v35 := GenS1{F0: 1, F1: 42}
        _ = v34
        _ = v35
    }
    var v36 bool = (!(!(-2147483648 < 0)))
    switch v26 {
    case 32767:
        fmt.Println(v33)
    case 1:
        for fi37 := 0; fi37 < 6; fi37++ {
            switch v33 {
            case "foo bar":
                v36 = (v5 > v5)
            case "hello":
                if true {
                    var v38 int16 = 42
                    var v39 uint16 = 65535
                    _ = v38
                    _ = v39
                } else {
                    var v40 float32 = 0.1
                    if v6 {
                        v5 -= v5
                        switch ((v33 + ((v33 + ((v33 + "\n\t") + "hello")) + v33)) + v33) {
                        case "":
                            fmt.Println(v26)
                        case "a":
                            switch fi37 {
                            case 128:
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    v42 := GenS2{F0: (v5 ^ (v5 / ((2147483647) | 1))), F1: 189, F2: 32767}
                                    _ = fi41
                                    _ = v42
                                }
                            default:
                                var v43 float64 = -2.25
                                _ = v43
                            }
                        case "hello":
                            if false {
                                if (2147483647 >= 2147483647) {
                                    for fi44 := 0; fi44 < 3; fi44++ {
                                        switch v26 {
                                        case 4294967295:
                                            switch 127 {
                                            case 9223372036854775807:
                                                v45 := GenS2{F0: (v5 & (v5 + (v5 & v5))), F1: 317, F2: 2}
                                                _ = v45
                                            default:
                                                v1 += v1
                                            }
                                        case 42:
                                            var v46 uint64 = 2
                                            _ = v46
                                        case 2147483647:
                                            delete(v2, 9223372036854775807)
                                        }
                                        _ = fi44
                                    }
                                } else {
                                    fmt.Println(fi37, v36, v26)
                                }
                            }
                        default:
                            v47 := GenS3{F0: (v1 % ((uint8(v5)) | 1)), F1: 42, F2: 42}
                            _ = v47
                        }
                        var v48 float32 = 1.5
                        _ = v48
                    } else {
                        v49 := GenS4{F0: v33, F1: v1}
                        if (!v36) {
                            if true {
                                fmt.Println(v6)
                            } else {
                                var v50 uint32 = v5
                                if (!v36) {
                                    var v51 int32 = 2147483647
                                    if ((1.5 / ((0.0 + 0.1) + (0.1 - 1.5))) >= 0.1) {
                                        var v52 float32 = (v40 / v40)
                                        for fi53 := 0; fi53 < 2; fi53++ {
                                            if v36 {
                                                var v54 uint16 = 255
                                                if (((float32(0.1) * v52) < (1e10 - (v52 / (v52 / (float32(((3.14159 - 3.14159) - -2.25)) / float32(1.5)))))) || ((!true) || (v36 || v36))) {
                                                    v54 = v54
                                                    v50 = (^(v50))
                                                    var v55 int32 = (^(v51))
                                                    _ = v55
                                                }
                                                _ = v54
                                            }
                                            v4 = append(v4, 4294967295)
                                            _ = fi53
                                        }
                                        _ = v52
                                    }
                                    v56 := GenS2{F0: v5, F1: 139, F2: 255}
                                    _ = v51
                                    _ = v56
                                } else {
                                    for fi57 := 0; fi57 < 1; fi57++ {
                                        switch v33 {
                                        case "foo bar":
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(v1)
                                            default:
                                                switch (-(v26)) {
                                                case 2147483647:
                                                    var v58 int64 = 42
                                                    _ = v58
                                                default:
                                                    v6 = false
                                                }
                                            }
                                        default:
                                            v40 = ((0.1 / (v40 * ((v40 / (v40 - (float32(-2.25) - (0.0 / v40)))) / (v40 - float32(-2.25))))) / 0.1)
                                        }
                                        var v59 uint64 = 2
                                        _ = fi57
                                        _ = v59
                                    }
                                    var v60 float32 = float32(1.5)
                                    _ = v60
                                }
                                var v61 int64 = -1
                                _ = v50
                                _ = v61
                            }
                            for fi62 := 0; fi62 < 13; fi62++ {
                                v63 := GenS4{F0: (((v33 + (v33 + (v33 + v33))) + (v33 + v33)) + ((v33 + (((v33 + v33) + "\n\t") + v33)) + v33)), F1: (v1 << ((35) & 63))}
                                _ = fi62
                                _ = v63
                            }
                            switch "" {
                            case "hello":
                                for fi64 := 0; fi64 < 17; fi64++ {
                                    switch "" {
                                    case "\n\t":
                                        v1 = v1
                                    }
                                    _ = fi64
                                }
                            case "\n\t":
                                v26 += -2147483648
                            case "foo bar":
                                for fi65 := 0; fi65 < 8; fi65++ {
                                    var v66 uint64 = 2
                                    var v67 int8 = -128
                                    if (0 < 42) {
                                        v68 := GenS2{F0: 324, F1: 65535, F2: 683}
                                        fmt.Println(v26)
                                        var v69 int32 = -716
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 float64 = 0.0
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                    _ = v70
                                }
                            }
                        }
                        _ = v49
                    }
                    _ = v40
                }
            case "\n\t":
                if v6 {
                    v71 := GenS2{F0: (v5 << ((2) & 63)), F1: 455, F2: 2}
                    v26 += int(v5)
                    v33 = ""
                    _ = v71
                }
            default:
                v5 = v5
            }
            v5 = v5
            _ = fi37
        }
    }
    v72 := GenS2{F0: (v5 - v5), F1: 18446744073709551615, F2: 1}
    var v73 int8 = -128
    v4 = append(v4, 0)
    var v74 uint32 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
    _ = v74
    _ = v8
}
