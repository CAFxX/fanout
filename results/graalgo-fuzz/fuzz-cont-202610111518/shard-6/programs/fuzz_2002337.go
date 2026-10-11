package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint64; F2 string }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0(p0 bool, p1 uint16) int8 {
    if false {
        p0 = false
    }
    v6 := []int{-9223372036854775808, 42, 560, -174}
    p1 = (^(p1))
    if len(v6) > 0 {
        v6[0] = int(p1)
    }
    p0 = (p0 || ((int16(p1) < 1) || p0))
    _ = p0
    _ = p1
    _ = v6
    return 100
}
func genFn1(p0 int64, p1 int16) uint32 {
    fmt.Println(p0, p0)
    v7 := []int{32767, -296}
    var v8 bool = false
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: ("a" + ("hello" + "")), F1: (p1 & -32768), F2: 42}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        switch 128 {
        case 9223372036854775807:
            v12 := GenS1{F0: ("foo bar" + "a"), F1: 1000, F2: 100}
            _ = v12
        case -1:
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 uint16 = 655
                switch fi11 {
                case -880:
                    var v15 int = fi11
                    fmt.Println(fi11, v15)
                    _ = v15
                case 32767:
                    var v16 int64 = -279
                    _ = v16
                default:
                    var v17 uint8 = 197
                    _ = v17
                }
                var v18 int32 = 632
                switch (fi13 % (((^(fi11))) | 1)) {
                case 42:
                    var v19 int16 = (p1 * (p1 + (p1 * 46)))
                    _ = v19
                case 9223372036854775807:
                    switch fi11 {
                    case 127:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    case -853:
                        _ = len(v7)
                    case 785:
                        var v20 int32 = (v18 << ((65535) & 63))
                        _ = v20
                    }
                }
                _ = len(v7)
                _ = fi13
                _ = v14
                _ = v18
            }
        case 2147483647:
            var v21 uint = 4294967295
            _ = v21
        default:
            p0 += 1
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 int16 = 0
    var v2 string = (("\n\t" + "hello") + ("a" + ("" + ("foo bar" + ("hello" + (("\n\t" + (("foo bar" + ("foo bar" + "hello")) + ((("\n\t" + "a") + "") + "hello"))) + ""))))))
    var v3 float64 = 0.1
    var v4 int32 = -2147483648
    var v5 int32 = (v4 & v4)
    v22 := []int{}
    v1 -= (^(v1))
    var v23 float32 = (1e10 + 0.1)
    var v24 bool = (!(!(v2 == "foo bar")))
    if v24 {
        var v25 uint = 1
        if true {
            v24 = v24
            switch -1 {
            case 0:
                v26 := GenS1{F0: v2, F1: -32768, F2: -1}
                _ = v26
            case 127:
                v1 = 42
            }
            v27 := GenS1{F0: "\n\t", F1: (^(v1)), F2: int8(v25)}
            _ = v27
        } else {
            var v28 uint = (v25 % ((v25) | 1))
            switch ((v2 + "hello") + v2) {
            case "hello":
                switch v2 {
                case "a":
                    for fi29 := 0; fi29 < 7; fi29++ {
                        if ((2 <= 496) || v24) {
                            v1 = 1000
                            fmt.Println(v2)
                            v22 = append(v22, 42)
                        } else {
                            switch 4294967295 {
                            case 1:
                                switch (fi29 * (fi29 - (fi29 + fi29))) {
                                case 120:
                                    for fi30 := 0; fi30 < 19; fi30++ {
                                        v22 = append(v22, (-(fi30)))
                                        _ = fi30
                                    }
                                case 4294967295:
                                    v31 := GenS1{F0: "a", F1: 42, F2: 0}
                                    _ = v31
                                case -204:
                                    v22 = append(v22, fi29)
                                default:
                                    v2 = "a"
                                }
                            case 395:
                                v32 := GenS1{F0: ("a" + v2), F1: (v1 ^ v1), F2: -128}
                                _ = v32
                            default:
                                switch "foo bar" {
                                case "a":
                                    for fi33 := 0; fi33 < 19; fi33++ {
                                        v24 = v24
                                        fmt.Println(v1, v5)
                                        var v34 uint8 = 0
                                        _ = fi33
                                        _ = v34
                                    }
                                case "foo bar":
                                    v25 -= v28
                                case "\n\t":
                                    _ = len(v22)
                                default:
                                    v25 = v28
                                }
                            }
                            var v35 uint64 = 2
                            var v36 int32 = (v5 | (v4 / (((v4 / (((v5 | (v4 ^ v4))) | 1))) | 1)))
                            _ = v35
                            _ = v36
                        }
                        _ = fi29
                    }
                case "hello":
                    var v37 int64 = -223
                    _ = v37
                case "\n\t":
                    v24 = (((("\n\t" + v2) + "foo bar") + ((("hello" + v2) + "hello") + "a")) >= v2)
                }
            case "":
                for fi38 := 0; fi38 < 3; fi38++ {
                    var v39 int16 = int16(v28)
                    v4 = (v5 / ((0) | 1))
                    _ = fi38
                    _ = v39
                }
            default:
                v28 = (v28 / ((v25) | 1))
            }
            fmt.Println(v3)
            _ = v28
        }
        v40 := GenS2{F0: (v1 - (^(v1))), F1: v25}
        _ = v25
        _ = v40
    }
    if v24 {
        for fi41 := 0; fi41 < 6; fi41++ {
            var v42 string = v2
            var v43 float64 = float64(1.5)
            if len(v22) > 0 {
                v22[0] = fi41
            }
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    var v44 int = -666
    for fi45 := 0; fi45 < 9; fi45++ {
        if (9223372036854775807 != -209) {
            var v46 uint64 = 613
            var v47 int8 = 6
            v2 = v2
            _ = v46
            _ = v47
        } else {
            for fi48 := 0; fi48 < 14; fi48++ {
                fmt.Println(v3)
                _ = fi48
            }
        }
        if (72 > 116) {
            switch v2 {
            case "foo bar":
                if true {
                    if false {
                        fmt.Println(v3, v1, v24)
                        if (genFn1(0, -1) < 4294967295) {
                            if v24 {
                                v23 = ((v23 * v23) + v23)
                            }
                            var v49 int = (v44 >> ((58) & 63))
                            v49 -= v44
                            _ = v49
                        } else {
                            var v50 bool = false
                            v51 := GenS3{F0: 4294967295, F1: fi45, F2: (v5 % ((v5) | 1))}
                            _ = v50
                            _ = v51
                        }
                        v1 = (v1 >> ((363) & 63))
                    } else {
                        v23 = v23
                        for fi52 := 0; fi52 < 2; fi52++ {
                            var v53 int8 = int8(v5)
                            _ = fi52
                            _ = v53
                        }
                    }
                    switch v2 {
                    case "hello":
                        var v54 int16 = (v1 >> ((18446744073709551615) & 63))
                        _ = v54
                    }
                    fmt.Println(v24, v4, v2)
                }
            case "":
                var v55 float32 = (float32(-2.25) * ((-2.25 / v23) / v23))
                _ = v55
            case "a":
                v56 := GenS4{F0: 1}
                _ = v56
            }
            v44 += (v44 | -2147483648)
        }
        _ = fi45
    }
    for fi57 := 0; fi57 < 9; fi57++ {
        if true {
            fmt.Println(v4, fi57, v1)
            v44 -= v44
        } else {
            v4 -= 1
        }
        for fi58 := 0; fi58 < 11; fi58++ {
            var v59 uint = 285
            _ = fi58
            _ = v59
        }
        _ = fi57
    }
    if ((!true) || (!(104 <= 2))) {
        v24 = (942 <= 4294967295)
        for fi60 := 0; fi60 < 1; fi60++ {
            switch v44 {
            case 42:
                if v24 {
                    v61 := GenS5{F0: v1, F1: 0, F2: "foo bar"}
                    var v62 uint32 = 108
                    var v63 int8 = 71
                    _ = v61
                    _ = v62
                    _ = v63
                }
            case -532:
                var v64 float64 = ((((3.14159 / (v3 / (((v3 - 0.0) + (((((v3 + v3) * (((1.5 - v3) + (float64(0.1) - v3)) / (v3 * v3))) + v3) / v3) + float64((((v23 / (((-2.25 * (v23 + v23)) / 1.0) * v23)) + v23) * v23)))) * (v3 * (-2.25 + v3))))) * v3) / v3) + 0.1)
                _ = v64
            }
            fmt.Println(v3, v23)
            _ = fi60
        }
        switch (v2 + "foo bar") {
        case "\n\t":
            for fi65 := 0; fi65 < 11; fi65++ {
                fmt.Println(v2, fi65, fi65)
                v5 = v4
                v5 -= v4
                _ = fi65
            }
        }
    }
    fmt.Println(v2, v3, v23)
    v66 := GenS6{F0: 934, F1: 9223372036854775807}
    v24 = (1 >= 42)
    v67 := GenS2{F0: (^(v1)), F1: 42}
    fmt.Println(v1, v5, v1)
    v2 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v66
    _ = v67
}
