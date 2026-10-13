package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 bool) uint64 {
    v6 := []int{-211, 1, -9223372036854775808}
    var v7 int8 = -1
    _ = p0
    _ = v6
    _ = v7
    return 724
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = -15
    var v3 string = ""
    var v4 uint = 1
    v5 := []int{-2147483648, 32767}
    if (!(-1 <= -9223372036854775808)) {
        fmt.Println(v3, v4)
        fmt.Println(v2)
    } else {
        if (255 < 32767) {
            var v8 bool = ((!true) || false)
            v9 := GenS1{F0: 42, F1: v2, F2: (^(v1))}
            var v10 string = (((v3 + "foo bar") + v3) + ("foo bar" + v3))
            _ = v10
            _ = v8
            _ = v9
        }
        if len(v5) > 0 {
            v5[0] = 0
        }
        if true {
            v11 := GenS1{F0: 1, F1: (v2 / (((^(v2))) | 1)), F2: genFn0((!false))}
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 uint8 = 128
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            var v14 int32 = v2
            for fi15 := 0; fi15 < 17; fi15++ {
                switch v3 {
                case "foo bar":
                    if false {
                        var v16 int64 = -94
                        _ = v16
                    } else {
                        v4 -= (v4 << (((v4 >> ((v4) & 63))) & 63))
                    }
                case "\n\t":
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if ((!true) || (!(true || (float32(1.5) != (-2.25 - 1e10))))) {
                            var v18 uint64 = 42
                            if len(v5) > 0 {
                                v5[0] = fi17
                            }
                            v14 = (v14 % ((v14) | 1))
                            _ = v18
                        } else {
                            fmt.Println(v2)
                        }
                        _ = fi17
                    }
                case "a":
                    var v19 string = ("" + "hello")
                    _ = v19
                default:
                    if len(v5) > 0 {
                        v5[0] = fi15
                    }
                }
                if (v1 != 18446744073709551615) {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        switch (v3 + (v3 + "")) {
                        case "foo bar":
                            var v21 float32 = -2.25
                            _ = v21
                        case "":
                            v22 := GenS2{F0: (fi20 | fi15), F1: v1}
                            _ = v22
                        default:
                            if len(v5) > 0 {
                                v5[0] = fi20
                            }
                        }
                        _ = fi20
                    }
                    if false {
                        if true {
                            v1 -= v1
                        } else {
                            var v23 uint = 0
                            for fi24 := 0; fi24 < 14; fi24++ {
                                _ = len(v5)
                                _ = fi24
                            }
                            _ = v23
                        }
                        var v25 bool = false
                        _ = v25
                    }
                    var v26 bool = (!false)
                    _ = v26
                } else {
                    var v27 float64 = 1.5
                    _ = v27
                }
                var v28 int64 = 1
                _ = fi15
                _ = v28
            }
            if (v3 < v3) {
                if false {
                    var v29 float64 = 1e10
                    _ = v29
                }
                var v30 string = ("foo bar" + v3)
                var v31 uint64 = genFn0(("foo bar" >= v30))
                _ = v30
                _ = v31
            } else {
                v5 = append(v5, 2147483647)
                v14 = (v2 / ((v14) | 1))
            }
            _ = v14
        }
    }
    var v32 int = 32768
    v3 = v3
    fmt.Println(v1, v4, v1)
    v33 := GenS2{F0: (v32 + v32), F1: (v1 << ((2) & 63))}
    if (((1e10 / 1.0) / 3.14159) < (0.0 * float32(float64(0.1)))) {
        var v34 float32 = 0.1
        v3 = ""
        v35 := GenS3{F0: 506}
        _ = v34
        _ = v35
    }
    var v36 int16 = 1
    if (!(49 != -1)) {
        fmt.Println(v32, v3, v4)
    }
    var v37 string = ("" + v3)
    var v38 uint32 = 0
    v5 = append(v5, (v32 ^ 9223372036854775807))
    if true {
        var v39 uint8 = 128
        _ = v39
    }
    var v40 bool = (128 > 5)
    switch (v32 >> (((v4 << ((598) & 63))) & 63)) {
    case -524:
        var v41 int32 = 42
        _ = v41
    }
    if ((v1 ^ (^(v1))) >= (v1 << (((^(v4))) & 63))) {
        v42 := GenS4{F0: (^(v38)), F1: genFn0((!false)), F2: 9223372036854775807}
        for fi43 := 0; fi43 < 16; fi43++ {
            v3 = "foo bar"
            _ = fi43
        }
        _ = v42
    }
    switch int(v4) {
    case -9223372036854775808:
        _ = len(v5)
    case 431:
        v44 := GenS5{F0: (v2 * (v2 * (v2 + v2))), F1: (v4 * v4)}
        _ = v44
    case -973:
        fmt.Println(v2, v37)
    }
    if v40 {
        v5 = append(v5, v32)
        for fi45 := 0; fi45 < 8; fi45++ {
            switch 4294967295 {
            case 4294967295:
                v3 = v3
            case 2147483647:
                for fi46 := 0; fi46 < 18; fi46++ {
                    if ((v4 & 255) <= v4) {
                        if v40 {
                            if v40 {
                                if (670 == 482) {
                                    v3 = (v37 + (v3 + "hello"))
                                } else {
                                    v47 := GenS6{F0: (^(fi45)), F1: v4}
                                    v36 -= v36
                                    v32 += -2147483648
                                    _ = v47
                                }
                                v48 := GenS2{F0: (fi46 + (fi45 ^ fi45)), F1: genFn0((0 == 52))}
                                switch fi45 {
                                case -296:
                                    fmt.Println(v2, v36)
                                case 32767:
                                    if (v1 < uint64(v4)) {
                                        fmt.Println(fi45, v37)
                                    } else {
                                        if ((0.0 == float32(((3.14159 + 1e10) + 3.14159))) && v40) {
                                            _ = len(v5)
                                            var v49 uint32 = 669
                                            _ = v49
                                        }
                                        var v50 float64 = (1e10 - (float64(1.5) - 0.0))
                                        _ = v50
                                    }
                                case 127:
                                    fmt.Println(v36)
                                default:
                                    var v51 int = fi45
                                    _ = v51
                                }
                                v52 := GenS2{F0: (v32 - 9223372036854775807), F1: 321}
                                _ = v48
                                _ = v52
                            }
                        } else {
                            if true {
                                fmt.Println(v40)
                                if v40 {
                                    var v53 uint16 = 255
                                    var v54 uint8 = 0
                                    _ = v53
                                    _ = v54
                                }
                                var v55 uint64 = 2
                                _ = v55
                            }
                            if false {
                                v56 := GenS1{F0: -149, F1: -1, F2: 2}
                                _ = v56
                            } else {
                                var v57 bool = v40
                                _ = v57
                            }
                        }
                        var v58 int32 = v2
                        _ = v58
                    } else {
                        var v59 int8 = 1
                        v60 := GenS3{F0: (v1 ^ v1)}
                        var v61 float32 = 1.5
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    v62 := GenS7{F0: v2, F1: v38, F2: (v38 + v38)}
                    _ = fi46
                    _ = v62
                }
            }
            _ = fi45
        }
        var v63 float32 = (0.1 + (1e10 / ((-2.25 + 1.5) + float32(3.14159))))
        _ = v63
    } else {
        fmt.Println(v4, v36, v38)
        if (float64(float32(1e10)) == (0.1 + -2.25)) {
            v2 = v2
            v36 = -32768
            var v64 float64 = ((3.14159 + 0.1) * (1e10 * -2.25))
            _ = v64
        }
        var v65 float64 = float64(0.1)
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
