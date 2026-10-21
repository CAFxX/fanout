package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16; F1 bool; F2 uint8 }
type GenS6 struct { F0 int32 }
func genFn0() int16 {
    switch ("foo bar" + ("" + "")) {
    case "\n\t":
        var v6 uint8 = 146
        _ = v6
    default:
        var v7 uint32 = 668
        _ = v7
    }
    if false {
        v8 := GenS1{F0: -32768, F1: "foo bar"}
        v9 := []int{127, 9223372036854775807}
        if true {
            switch "hello" {
            case "\n\t":
                var v10 int8 = -128
                _ = v10
            }
            switch ("a" + "hello") {
            case "a":
                if (!true) {
                    var v11 int16 = 32767
                    _ = v11
                } else {
                    var v12 int32 = -156
                    if len(v9) > 0 {
                        v9[0] = 128
                    }
                    switch "hello" {
                    case "\n\t":
                        if false {
                            var v13 int32 = (v12 >> ((0) & 63))
                            var v14 float64 = 1e10
                            v15 := GenS1{F0: 42, F1: "foo bar"}
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            if false {
                                var v16 uint16 = 707
                                _ = v16
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = -648
                                }
                                if true {
                                    var v17 int64 = 1
                                    v12 -= 0
                                    v17 = (v17 ^ -475)
                                    _ = v17
                                } else {
                                    if false {
                                        v12 -= 0
                                        v18 := GenS2{F0: 42, F1: 65535, F2: 106}
                                        v19 := GenS1{F0: 32767, F1: ""}
                                        _ = v18
                                        _ = v19
                                    }
                                    v12 += (v12 - -1)
                                }
                                for fi20 := 0; fi20 < 19; fi20++ {
                                    var v21 bool = (true || false)
                                    var v22 int16 = 0
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                            for fi23 := 0; fi23 < 2; fi23++ {
                                var v24 int16 = 42
                                _ = fi23
                                _ = v24
                            }
                            switch "\n\t" {
                            case "\n\t":
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    _ = fi25
                                }
                            case "hello":
                                var v26 int16 = 289
                                _ = v26
                            case "foo bar":
                                v12 += v12
                            }
                        }
                    case "hello":
                        if true {
                            v12 = v12
                            var v27 int64 = 0
                            var v28 float64 = 1.5
                            _ = v27
                            _ = v28
                        }
                    case "foo bar":
                        for fi29 := 0; fi29 < 9; fi29++ {
                            if len(v9) > 0 {
                                v9[0] = 32768
                            }
                            _ = fi29
                        }
                    default:
                        var v30 int = 32768
                        _ = v30
                    }
                    _ = v12
                }
            case "\n\t":
                var v31 uint16 = 35
                _ = v31
            case "":
                var v32 int64 = -1
                _ = v32
            }
        } else {
            var v33 float64 = (3.14159 - 3.14159)
            _ = v33
        }
        _ = v8
        _ = v9
    }
    return -1
}
func genFn1(p0 int32) int16 {
    fmt.Println(p0)
    if true {
        var v34 int64 = 174
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 int = fi35
            var v37 int8 = -128
            _ = fi35
            _ = v36
            _ = v37
        }
        v34 = v34
        _ = v34
    } else {
        p0 = p0
    }
    var v38 int8 = 100
    fmt.Println(p0, v38)
    _ = p0
    _ = v38
    return 32767
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 4294967295
    var v3 string = "a"
    var v4 int32 = 1
    var v5 float32 = 1.5
    v1[("foo bar" + ("hello" + "\n\t"))] = v3
    for fi39 := 0; fi39 < 2; fi39++ {
        if false {
            for fi40 := 0; fi40 < 20; fi40++ {
                fmt.Println(v5)
                _ = fi40
            }
        }
        if false {
            v41 := GenS2{F0: 128, F1: 65535, F2: 105}
            _ = v41
        } else {
            v4 = v4
        }
        v3 = "\n\t"
        _ = fi39
    }
    v2 -= (v2 ^ v2)
    for fi42 := 0; fi42 < 8; fi42++ {
        if false {
            var v43 int32 = v4
            v44 := v1["a"]
            for fi45 := 0; fi45 < 8; fi45++ {
                v2 = v2
                _ = fi45
            }
            _ = v43
            _ = v44
        }
        for fi46 := 0; fi46 < 9; fi46++ {
            v3 = (("foo bar" + v3) + "foo bar")
            v4 -= 333
            _ = fi46
        }
        _ = fi42
    }
    v47 := GenS2{F0: 35, F1: 210, F2: 1}
    if false {
        v48 := GenS3{F0: v3, F1: 4294967295}
        v49 := []int{-229}
        _ = v48
        _ = v49
    }
    v50 := GenS3{F0: ((v3 + (("\n\t" + ("a" + ("hello" + (v3 + ("" + ""))))) + v3)) + "foo bar"), F1: 18446744073709551615}
    v51 := GenS1{F0: genFn1(v4), F1: v3}
    if true {
        v52 := GenS4{F0: (969 == 662), F1: v4}
        _ = v52
    } else {
        if (2147483647 < 2147483647) {
            if false {
                if (1 >= 42) {
                    var v53 float32 = float32(0.0)
                    switch v3 {
                    case "foo bar":
                        if ("foo bar" != (((("foo bar" + "foo bar") + "\n\t") + (v3 + ("\n\t" + v3))) + ((v3 + (("\n\t" + v3) + v3)) + ""))) {
                            switch 32768 {
                            case -9223372036854775808:
                                fmt.Println(v5, v5, v3)
                            case -998:
                                var v54 uint16 = 231
                                _ = v54
                            }
                            v3 = ("" + ("\n\t" + ((v3 + "\n\t") + "")))
                        }
                    case "":
                        v4 = (v4 | v4)
                    case "\n\t":
                        v4 -= (^(v4))
                    default:
                        fmt.Println(v4)
                    }
                    _ = v53
                } else {
                    for fi55 := 0; fi55 < 18; fi55++ {
                        var v56 float64 = ((3.14159 + float64(0.1)) * 0.1)
                        _ = fi55
                        _ = v56
                    }
                    v3 = "foo bar"
                    for fi57 := 0; fi57 < 7; fi57++ {
                        for fi58 := 0; fi58 < 20; fi58++ {
                            if true {
                                v3 = v3
                                var v59 int = (-(fi58))
                                v5 += 0.1
                                _ = v59
                            } else {
                                v60 := GenS2{F0: 1, F1: 42, F2: -95}
                                fmt.Println(fi57)
                                _ = v60
                            }
                            v61 := GenS5{F0: 32767, F1: (!(!((fi57 + fi57) == -1))), F2: 1}
                            _ = fi58
                            _ = v61
                        }
                        var v62 int = fi57
                        if false {
                            var v63 int = v62
                            switch (("foo bar" + "") + v3) {
                            case "a":
                                var v64 int = (^(v62))
                                _ = v64
                            case "\n\t":
                                var v65 uint16 = 255
                                _ = v65
                            }
                            _ = v63
                        }
                        _ = fi57
                        _ = v62
                    }
                }
            } else {
                for fi66 := 0; fi66 < 9; fi66++ {
                    var v67 uint16 = 1
                    _ = fi66
                    _ = v67
                }
                if (((-116 != -128) || (((v3 + (v3 + ((("\n\t" + ((v3 + "foo bar") + ("hello" + (("a" + v3) + ("" + v3))))) + "\n\t") + ("\n\t" + "\n\t")))) + "a") < v3)) && false) {
                    var v68 int8 = -78
                    for fi69 := 0; fi69 < 18; fi69++ {
                        var v70 int16 = genFn1((v4 & (v4 / ((v4) | 1))))
                        delete(v1, ((v3 + ("\n\t" + (v3 + (("" + v3) + "foo bar")))) + ("a" + (v3 + v3))))
                        fmt.Println(v2, v4, v2)
                        _ = fi69
                        _ = v70
                    }
                    for fi71 := 0; fi71 < 7; fi71++ {
                        fmt.Println(v2, v2)
                        for fi72 := 0; fi72 < 9; fi72++ {
                            var v73 uint32 = 65535
                            fmt.Println(v73, fi71)
                            _ = fi72
                            _ = v73
                        }
                        _ = fi71
                    }
                    _ = v68
                }
                v3 = v3
            }
        } else {
            for fi74 := 0; fi74 < 8; fi74++ {
                var v75 int32 = (v4 % ((154) | 1))
                v3 = (v3 + "hello")
                _ = fi74
                _ = v75
            }
            if true {
                for fi76 := 0; fi76 < 8; fi76++ {
                    if (!(false || true)) {
                        v77 := GenS1{F0: 1000, F1: "a"}
                        var v78 string = ""
                        _ = v77
                        _ = v78
                    }
                    v5 -= 0.0
                    _ = fi76
                }
                for fi79 := 0; fi79 < 9; fi79++ {
                    v80 := GenS3{F0: ((v3 + v3) + v3), F1: 617}
                    v81 := GenS2{F0: 127, F1: 32767, F2: 93}
                    _ = fi79
                    _ = v80
                    _ = v81
                }
            }
            v3 = "\n\t"
        }
    }
    var v82 int64 = 0
    if (false || false) {
        v1[v3] = v3
        v83 := v1[v3]
        for fi84 := 0; fi84 < 2; fi84++ {
            delete(v1, "foo bar")
            _ = fi84
        }
        _ = v83
    }
    if true {
        v85 := GenS2{F0: 226, F1: 65535, F2: 127}
        v5 = v5
        fmt.Println(v4)
        _ = v85
    }
    fmt.Println(v3, v5)
    if (!(127 <= 2147483647)) {
        if (v2 > (v2 & v2)) {
            var v86 float32 = float32((1e10 * 0.0))
            v87 := GenS4{F0: (!true), F1: v4}
            _ = v86
            _ = v87
        } else {
            for fi88 := 0; fi88 < 1; fi88++ {
                var v89 int = int(v82)
                _ = fi88
                _ = v89
            }
            var v90 string = v3
            var v91 int32 = (v4 << ((42) & 63))
            _ = v90
            _ = v91
        }
    } else {
        v4 = v4
    }
    for fi92 := 0; fi92 < 18; fi92++ {
        if (346 > 2) {
            fmt.Println(v4)
            var v93 bool = (0 <= -1)
            _ = v93
        } else {
            for fi94 := 0; fi94 < 14; fi94++ {
                fmt.Println(v4)
                v95 := v1[(((v3 + ("\n\t" + ("hello" + v3))) + ("a" + "\n\t")) + "hello")]
                v96 := GenS6{F0: 0}
                _ = fi94
                _ = v95
                _ = v96
            }
            fmt.Println(v4, v4)
        }
        v82 = 1
        var v97 float64 = 1.5
        _ = fi92
        _ = v97
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v50
    _ = v51
    _ = v82
}
