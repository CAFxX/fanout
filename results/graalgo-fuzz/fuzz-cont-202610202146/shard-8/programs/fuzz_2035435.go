package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int64; F2 int16 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int32; F1 uint16; F2 int64 }
type GenS8 struct { F0 string; F1 int16; F2 int32 }
func genFn0() string {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 18446744073709551615}
        var v8 int32 = 1
        switch fi6 {
        case -264:
            fmt.Println(v8, fi6)
        case 127:
            v8 -= v8
        case 18:
            v9 := GenS2{F0: -6}
            _ = v9
        default:
            switch "\n\t" {
            case "":
                v8 = (v8 - (^(v8)))
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if (float64(1e10) >= ((3.14159 * 1e10) + (float64(float32(1e10)) + 0.0))) {
        if (true || ((683 <= 479) && false)) {
            var v10 int = 127
            _ = v10
        }
        var v11 uint16 = 820
        _ = v11
    }
    if false {
        if (false && (true && true)) {
            if false {
                for fi12 := 0; fi12 < 15; fi12++ {
                    if true {
                        var v13 uint32 = 176
                        v13 -= (v13 ^ (^(v13)))
                        _ = v13
                    } else {
                        var v14 uint8 = 2
                        v15 := []int{-9223372036854775808, -785, 9223372036854775807, fi12}
                        fmt.Println(v14, fi12)
                        _ = v14
                        _ = v15
                    }
                    v16 := []int{-9223372036854775808, 32767, fi12}
                    _ = fi12
                    _ = v16
                }
            }
            switch "hello" {
            case "hello":
                v17 := GenS2{F0: 0}
                _ = v17
            default:
                if (!true) {
                    var v18 uint16 = 0
                    fmt.Println(v18)
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        fmt.Println(fi19)
                        _ = fi19
                    }
                }
            }
            var v20 string = "foo bar"
            _ = v20
        }
        var v21 int = -451
        _ = v21
    } else {
        var v22 float32 = -2.25
        var v23 int8 = 100
        if false {
            v24 := GenS3{F0: "\n\t"}
            v23 += v23
            var v25 int32 = -1
            _ = v24
            _ = v25
        }
        _ = v22
        _ = v23
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int8 = -1
        v27 = 100
        _ = fi26
        _ = v27
    }
    var v28 uint8 = 42
    _ = v28
    return ("a" + "")
}

func main() {
    v1 := []string{("foo bar" + "\n\t")}
    var v2 int = -2147483648
    v3 := []int{(v2 / (((v2 >> ((175) & 63))) | 1)), (v2 * 32768)}
    v4 := make(map[string]bool)
    var v5 uint64 = 235
    var v29 uint16 = 1
    v5 += v5
    for fi30 := 0; fi30 < 13; fi30++ {
        v29 += (v29 % ((0) | 1))
        _ = fi30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS3{F0: ("" + "hello")}
        var v33 float32 = (0.1 - 3.14159)
        _ = len(v4)
        _ = fi31
        _ = v32
        _ = v33
    }
    v5 = (v5 + 380)
    v34 := GenS4{F0: (v29 / ((823) | 1))}
    v5 = v5
    for fi35 := 0; fi35 < 7; fi35++ {
        v2 = -9223372036854775808
        if ((true && false) && (true || false)) {
            if len(v3) > 0 {
                v3[0] = (^(v2))
            }
        } else {
            var v36 float64 = 1e10
            switch "foo bar" {
            case "a":
                if true {
                    if false {
                        if (!true) {
                            var v37 float64 = v36
                            if len(v3) > 0 {
                                v3[0] = (fi35 << ((542) & 63))
                            }
                            v37 = v36
                            _ = v37
                        }
                        if false {
                            v29 += v29
                            for fi38 := 0; fi38 < 10; fi38++ {
                                if (-1 <= 0) {
                                    _ = len(v3)
                                    var v39 float64 = (1e10 + (v36 * v36))
                                    var v40 uint16 = v29
                                    v39 = float64(0.0)
                                    v41 := GenS5{F0: "a", F1: 1, F2: 42}
                                    v42 := GenS6{F0: (((("foo bar" + (("a" + "hello") + "\n\t")) + "a") + ((("foo bar" + "foo bar") + ((("a" + "a") + "a") + (("" + "foo bar") + "foo bar"))) + ("foo bar" + "\n\t"))) + "a"), F1: 466}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                                _ = fi38
                            }
                            v3 = append(v3, int(v5))
                        }
                    } else {
                        v43 := GenS2{F0: -128}
                        _ = v43
                    }
                    v36 -= float64(1.5)
                } else {
                    var v44 uint64 = v5
                    _ = v44
                }
            case "hello":
                fmt.Println(v36)
            case "foo bar":
                v5 = 42
            }
            switch "a" {
            case "foo bar":
                for fi45 := 0; fi45 < 14; fi45++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        v36 = 1e10
                        var v46 int16 = -1
                        v47 := v4[("hello" + "hello")]
                        _ = v46
                        _ = v47
                    } else {
                        if (!false) {
                            var v48 int16 = 32767
                            v49 := GenS3{F0: (("hello" + "hello") + "foo bar")}
                            fmt.Println(fi35, v36)
                            _ = v48
                            _ = v49
                        } else {
                            fmt.Println(v2)
                        }
                    }
                    if true {
                        v2 = (fi35 / (((v2 * 785)) | 1))
                    } else {
                        var v50 float64 = v36
                        var v51 int8 = 127
                        var v52 uint32 = 2
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    if (65535 == 65535) {
                        v53 := GenS5{F0: "", F1: -1, F2: 764}
                        var v54 float64 = float64(3.14159)
                        switch -1 {
                        case 127:
                            var v55 float32 = -2.25
                            _ = v55
                        }
                        fmt.Println(fi45)
                        switch 0 {
                        case -820:
                            if (!(127 > -1)) {
                                var v56 int16 = -1
                                _ = v56
                            } else {
                                var v57 uint64 = v5
                                v57 = v57
                                _ = v57
                            }
                        case -712:
                            switch fi35 {
                            case 2147483647:
                                v54 = v54
                            case -2147483648:
                                var v58 int32 = 42
                                _ = v58
                            case 9223372036854775807:
                                var v59 float32 = (float32((0.0 / v54)) / ((1e10 * 1.5) * float32((1.5 / v54))))
                                _ = v59
                            }
                        case 65:
                            switch "" {
                            case "":
                                v5 = 42
                                v60 := v4["a"]
                                v60 = ((v5 * 0) >= (v5 % ((v5) | 1)))
                                _ = v60
                            case "a":
                                v54 += 3.14159
                            case "foo bar":
                                switch (fi45 * (fi35 & 1)) {
                                case 42:
                                    var v61 int = (fi35 | (fi35 & fi35))
                                    _ = v61
                                case -1:
                                    v62 := GenS7{F0: -2147483648, F1: (v29 ^ v29), F2: 185}
                                    _ = v62
                                default:
                                    fmt.Println(fi45, v36)
                                }
                            default:
                                v36 = 1e10
                            }
                        }
                        _ = v53
                        _ = v54
                    }
                    _ = fi45
                }
            }
            _ = v36
        }
        v63 := GenS7{F0: -2147483648, F1: (v29 % (((v29 & v29)) | 1)), F2: 0}
        _ = fi35
        _ = v63
    }
    if false {
        var v64 string = "a"
        for fi65 := 0; fi65 < 14; fi65++ {
            v66 := GenS8{F0: ("\n\t" + v64), F1: int16(v29), F2: 0}
            _ = fi65
            _ = v66
        }
        _ = v64
    }
    var v67 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v67
}
