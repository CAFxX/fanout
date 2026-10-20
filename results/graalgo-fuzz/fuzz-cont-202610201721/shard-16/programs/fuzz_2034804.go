package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 uint; F2 int }
type GenS7 struct { F0 uint64; F1 bool }

func main() {
    var v1 int = 42
    v2 := []string{"hello", "foo bar"}
    var v3 uint = 455
    v4 := make(map[uint32]int)
    switch "a" {
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            v1 -= 2147483647
        case "hello":
            if true {
                if (((230 >= 4294967295) && (false || (1 == 2))) && (!false)) {
                    var v5 int8 = -7
                    _ = v5
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + "")
                }
                v3 += (v3 & v3)
                v6 := GenS1{F0: 863, F1: 255}
                _ = v6
            }
        case "foo bar":
            switch v1 {
            case 2147483647:
                if true {
                    if false {
                        var v7 int64 = 0
                        if ((1000 != 891) && true) {
                            var v8 uint8 = 23
                            _ = v8
                        } else {
                            switch v1 {
                            case -202:
                                var v9 float32 = (1.5 - (3.14159 + (-2.25 * 1e10)))
                                _ = v9
                            case 9223372036854775807:
                                var v10 int64 = (v7 * v7)
                                _ = v10
                            case 32768:
                                for fi11 := 0; fi11 < 6; fi11++ {
                                    _ = len(v4)
                                    var v12 float32 = 3.14159
                                    v7 -= (v7 / (((v7 / ((v7) | 1))) | 1))
                                    _ = fi11
                                    _ = v12
                                }
                            }
                            var v13 uint = 0
                            _ = v13
                        }
                        _ = v7
                    }
                    if (!true) {
                        if ((0.0 + -2.25) > float32((1.5 + (0.0 + 3.14159)))) {
                            var v14 uint32 = 2147483647
                            if false {
                                var v15 int8 = 116
                                if ((!false) && true) {
                                    fmt.Println(v14)
                                    var v16 uint8 = 21
                                    if false {
                                        var v17 uint64 = 4294967295
                                        _ = v17
                                    }
                                    _ = v16
                                }
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    v19 := GenS1{F0: 2, F1: 127}
                                    v20 := GenS1{F0: 624, F1: 127}
                                    var v21 uint = v3
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                _ = v15
                            } else {
                                delete(v4, (^(v14)))
                                switch "a" {
                                case "\n\t":
                                    var v22 uint8 = 127
                                    _ = v22
                                case "hello":
                                    if (!(860 < 4294967295)) {
                                        var v23 float32 = (float32(float64(1.5)) - -2.25)
                                        _ = v23
                                    } else {
                                        v24 := GenS2{F0: 0, F1: 1}
                                        if true {
                                            v25 := GenS1{F0: 255, F1: 42}
                                            var v26 bool = (false && true)
                                            if true {
                                                if true {
                                                    v26 = (!true)
                                                    v27 := GenS2{F0: 15, F1: -32768}
                                                    v28 := GenS3{F0: v1}
                                                    _ = v27
                                                    _ = v28
                                                }
                                                var v29 int8 = -122
                                                if v26 {
                                                    v26 = v26
                                                    fmt.Println(v26, v3)
                                                    if (!v26) {
                                                        for fi30 := 0; fi30 < 8; fi30++ {
                                                            if ("a" < "foo bar") {
                                                                v3 = v3
                                                            } else {
                                                                var v31 int8 = (v29 + v29)
                                                                if v26 {
                                                                    var v32 float32 = ((((((1e10 * 3.14159) - 0.1) * 1.5) + float32((-2.25 - float64((0.0 + 0.0))))) + 0.1) / -2.25)
                                                                    if (42 == 963) {
                                                                        switch v1 {
                                                                        case 128:
                                                                            var v33 uint32 = (v14 & v14)
                                                                            _ = v33
                                                                        case 0:
                                                                            v31 -= -4
                                                                        case 32768:
                                                                            switch fi30 {
                                                                            case 128:
                                                                                var v34 string = ("" + "foo bar")
                                                                                _ = v34
                                                                            default:
                                                                                var v35 int8 = v31
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                        var v36 int = -1
                                                                        _ = v36
                                                                    }
                                                                    _ = v32
                                                                }
                                                                _ = v31
                                                            }
                                                            _ = fi30
                                                        }
                                                    } else {
                                                        var v37 int8 = (v29 + -128)
                                                        var v38 uint8 = 42
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                                _ = v29
                                            }
                                            _ = v25
                                            _ = v26
                                        } else {
                                            v39 := GenS4{F0: (v3 - v3), F1: (v14 * (v14 % ((v14) | 1)))}
                                            v40 := GenS4{F0: (v3 - v3), F1: (v14 ^ uint32(v3))}
                                            _ = v39
                                            _ = v40
                                        }
                                        _ = v24
                                    }
                                case "foo bar":
                                    if ("foo bar" < ("hello" + "a")) {
                                        v14 -= (v14 / ((154) | 1))
                                        var v41 float64 = 0.1
                                        _ = v41
                                    }
                                default:
                                    var v42 int32 = -498
                                    _ = v42
                                }
                                delete(v4, (v14 | 613))
                            }
                            _ = v14
                        } else {
                            v43 := v4[2147483647]
                            var v44 float32 = 1e10
                            if true {
                                v1 -= v1
                            }
                            _ = v43
                            _ = v44
                        }
                        switch ("foo bar" + "hello") {
                        case "hello":
                            var v45 uint8 = 255
                            _ = v45
                        case "foo bar":
                            delete(v4, 850)
                        case "\n\t":
                            v1 = v1
                        }
                        if len(v2) > 0 {
                            v2[0] = ((((("a" + "") + "a") + "a") + "a") + ("a" + ("foo bar" + "\n\t")))
                        }
                    } else {
                        v1 -= (v1 * (v1 | -2147483648))
                        var v46 uint8 = 0
                        _ = v46
                    }
                    if false {
                        var v47 uint64 = 0
                        for fi48 := 0; fi48 < 10; fi48++ {
                            var v49 uint16 = 373
                            var v50 int64 = -9223372036854775808
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        _ = v47
                    } else {
                        if (1.5 == 3.14159) {
                            fmt.Println(v3, v3)
                        }
                        var v51 int = v1
                        v3 = v3
                        _ = v51
                    }
                }
            default:
                var v52 int16 = 1
                _ = v52
            }
        }
    }
    if true {
        v3 = (v3 & 514)
        fmt.Println(v1, v3)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += v3
    switch "\n\t" {
    case "":
        for fi53 := 0; fi53 < 2; fi53++ {
            if (!(!((int64(fi53) != 0) || (false && (81 > 42))))) {
                for fi54 := 0; fi54 < 3; fi54++ {
                    v1 = fi54
                    var v55 float32 = float32(-2.25)
                    _ = fi54
                    _ = v55
                }
                v56 := GenS1{F0: 572, F1: 1}
                _ = v56
            }
            _ = fi53
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        for fi57 := 0; fi57 < 17; fi57++ {
            v1 = v1
            _ = fi57
        }
        var v58 uint8 = 0
        _ = v58
    } else {
        v2 = append(v2, "foo bar")
        switch -474 {
        case -488:
            for fi59 := 0; fi59 < 1; fi59++ {
                fmt.Println(v1, v1)
                var v60 int16 = -1
                _ = fi59
                _ = v60
            }
        case 128:
            if true {
                fmt.Println(v3)
                fmt.Println(v1, v1)
                v61 := GenS5{F0: (v3 * (^(v3)))}
                _ = v61
            } else {
                v62 := GenS6{F0: (v1 & 42), F1: v3, F2: -10}
                _ = v62
            }
        }
        var v63 uint = v3
        _ = v63
    }
    var v64 int32 = -2147483648
    var v65 uint32 = 2147483647
    fmt.Println(v65, v65, v65)
    v66 := GenS7{F0: 0, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v66
}
