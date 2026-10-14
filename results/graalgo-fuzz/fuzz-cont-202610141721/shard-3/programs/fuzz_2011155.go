package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int16; F2 int32 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string; F1 int64; F2 int16 }
func genFn0() uint8 {
    v6 := GenS1{F0: 806, F1: 18446744073709551615, F2: 45}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint16 = 402
        if (128 >= 127) {
            v8 -= (^(v8))
        } else {
            v8 -= (v8 | (v8 - (v8 + 42)))
        }
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return 128
}
func genFn1(p0 bool) string {
    var v10 int8 = -1
    var v11 bool = p0
    v10 = (^(v10))
    if (!(32767 == 151)) {
        v12 := []int{42, -9223372036854775808}
        _ = v12
    } else {
        v10 = (v10 - (v10 ^ (v10 | v10)))
    }
    switch "" {
    case "\n\t":
        if false {
            if v11 {
                v11 = true
                v13 := GenS1{F0: 1, F1: 18446744073709551615, F2: 1}
                _ = v13
            } else {
                p0 = v11
                v14 := []int{128, 32768, 358}
                _ = v14
            }
            v11 = (1 != 0)
        }
    case "hello":
        switch "a" {
        case "\n\t":
            if v11 {
                var v15 bool = (547 <= 854)
                var v16 int64 = 0
                if (445 == 4294967295) {
                    v10 = v10
                    if v15 {
                        var v17 float32 = 1.5
                        if false {
                            var v18 uint32 = 2147483647
                            _ = v18
                        } else {
                            var v19 int8 = (v10 >> ((262) & 63))
                            for fi20 := 0; fi20 < 2; fi20++ {
                                fmt.Println(fi20, fi20, fi20)
                                _ = fi20
                            }
                            _ = v19
                        }
                        switch ("" + "foo bar") {
                        case "a":
                            if (1 < -32768) {
                                if true {
                                    fmt.Println(p0, v10, p0)
                                    fmt.Println(v15, v10, v16)
                                }
                                fmt.Println(p0)
                            }
                        case "hello":
                            v21 := GenS1{F0: v16, F1: 692, F2: 2}
                            _ = v21
                        case "\n\t":
                            if (!("\n\t" < "a")) {
                                v16 = v16
                            } else {
                                for fi22 := 0; fi22 < 19; fi22++ {
                                    if (255 >= 1) {
                                        v16 = int64(v10)
                                    } else {
                                        v23 := []int{}
                                        _ = v23
                                    }
                                    _ = fi22
                                }
                            }
                        }
                        _ = v17
                    }
                } else {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 int8 = (v10 / ((-128) | 1))
                            p0 = (p0 && true)
                            _ = fi25
                            _ = v26
                        }
                        var v27 int32 = 2147483647
                        _ = fi24
                        _ = v27
                    }
                    fmt.Println(v16, v11)
                }
                _ = v15
                _ = v16
            }
        case "foo bar":
            p0 = v11
        case "a":
            v28 := []int{-702}
            _ = v28
        default:
            var v29 uint64 = 0
            _ = v29
        }
    case "foo bar":
        if (!p0) {
            switch 127 {
            case 1:
                for fi30 := 0; fi30 < 2; fi30++ {
                    for fi31 := 0; fi31 < 19; fi31++ {
                        var v32 float64 = (0.0 - 1.5)
                        var v33 int = -329
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    if v11 {
                        p0 = true
                        for fi34 := 0; fi34 < 16; fi34++ {
                            p0 = ((-(v10)) != v10)
                            _ = fi34
                        }
                        var v35 int32 = -1
                        _ = v35
                    }
                    _ = fi30
                }
            }
            switch 1 {
            case 349:
                v36 := GenS2{F0: 0}
                _ = v36
            default:
                switch 0 {
                case 128:
                    v37 := GenS3{F0: (100 == (v10 * (v10 / (((v10 / ((v10) | 1))) | 1))))}
                    _ = v37
                default:
                    fmt.Println(p0, p0, v10)
                }
            }
            p0 = v11
        }
    default:
        p0 = p0
    }
    _ = p0
    _ = v10
    _ = v11
    return (("hello" + "\n\t") + (("foo bar" + "hello") + "hello"))
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]bool)
    var v3 uint64 = 0
    v4 := []int{}
    var v5 int8 = 0
    _ = len(v4)
    if (((3.14159 - 0.1) * 0.0) < 1e10) {
        switch "hello" {
        case "a":
            var v38 bool = (("hello" + "\n\t") > ("\n\t" + ("hello" + ("hello" + ("a" + "a")))))
            _ = v38
        case "foo bar":
            fmt.Println(v5)
        case "\n\t":
            fmt.Println(v3, v1)
        default:
            v39 := GenS2{F0: uint8(v3)}
            _ = v39
        }
    } else {
        v40 := GenS3{F0: (false || true)}
        _ = v40
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        fmt.Println(fi41, v1)
        for fi42 := 0; fi42 < 8; fi42++ {
            v3 -= v3
            v43 := GenS3{F0: true}
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    if true {
        var v44 uint64 = 42
        v45 := GenS2{F0: 127}
        _ = v44
        _ = v45
    } else {
        _ = len(v2)
        v3 -= 42
    }
    v5 = (v5 ^ v5)
    for fi46 := 0; fi46 < 9; fi46++ {
        v5 += -128
        for fi47 := 0; fi47 < 2; fi47++ {
            v3 += uint64(v1)
            v48 := GenS4{F0: 134, F1: 194, F2: 1}
            _ = fi47
            _ = v48
        }
        v5 = v5
        _ = fi46
    }
    var v49 float32 = 0.0
    var v50 int = -54
    v51 := GenS1{F0: 0, F1: v3, F2: 1}
    var v52 uint8 = 255
    var v53 int32 = 42
    if false {
        switch v50 {
        case 4294967295:
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 uint64 = (v3 * (v3 << ((56) & 63)))
                if true {
                    v56 := GenS2{F0: genFn0()}
                    v5 -= v5
                    v57 := GenS5{F0: (v52 | v52), F1: (v5 ^ 0)}
                    _ = v56
                    _ = v57
                } else {
                    fmt.Println(v52, v49, v55)
                    var v58 int16 = -573
                    _ = v58
                }
                _ = fi54
                _ = v55
            }
        default:
            v59 := GenS3{F0: false}
            _ = v59
        }
        var v60 float32 = (float32(1e10) / 0.1)
        _ = v60
    }
    for fi61 := 0; fi61 < 1; fi61++ {
        v1 += 32767
        _ = fi61
    }
    if false {
        if true {
            var v62 string = ("" + "a")
            _ = v62
        }
        v1 += (v1 * v1)
    } else {
        v3 = v3
        v1 -= 42
    }
    v5 = (v5 << ((42) & 63))
    if true {
        if (((v5 << ((2) & 63)) < (v5 / (((v5 - (v5 >> ((18446744073709551615) & 63)))) | 1))) || true) {
            fmt.Println(v52, v1, v3)
            var v63 bool = true
            _ = v63
        } else {
            var v64 float32 = ((v49 * (1.5 * 1.5)) * (v49 / ((v49 - 0.0) / v49)))
            switch "foo bar" {
            case "a":
                v4 = append(v4, 4294967295)
            case "\n\t":
                fmt.Println(v1)
            case "":
                v52 += (v52 / ((v52) | 1))
            }
            var v65 string = "a"
            _ = v64
            _ = v65
        }
        for fi66 := 0; fi66 < 3; fi66++ {
            if len(v4) > 0 {
                v4[0] = v50
            }
            v67 := GenS4{F0: (^(v50)), F1: 42, F2: (v53 >> ((864) & 63))}
            v5 -= (v5 / (((v5 >> ((4294967295) & 63))) | 1))
            _ = fi66
            _ = v67
        }
        switch (((((("foo bar" + (("a" + "\n\t") + "a")) + "a") + ("\n\t" + "hello")) + "") + "foo bar") + "foo bar") {
        case "a":
            if ((v5 >> ((446) & 63)) > v5) {
                if false {
                    var v68 uint = 42
                    _ = v68
                } else {
                    if false {
                        switch 257 {
                        case -14:
                            v49 = (0.0 + 3.14159)
                        case 909:
                            var v69 string = "a"
                            _ = v69
                        case 42:
                            switch "" {
                            case "\n\t":
                                if (v52 <= (v52 | (v52 % ((genFn0()) | 1)))) {
                                    fmt.Println(v1, v53, v50)
                                    for fi70 := 0; fi70 < 3; fi70++ {
                                        v71 := GenS6{F0: "", F1: 919, F2: 1000}
                                        fmt.Println(v53)
                                        _ = fi70
                                        _ = v71
                                    }
                                    fmt.Println(v3)
                                } else {
                                    v72 := GenS3{F0: false}
                                    switch ("hello" + (("foo bar" + "\n\t") + "")) {
                                    case "":
                                        for fi73 := 0; fi73 < 5; fi73++ {
                                            v5 = v5
                                            _ = fi73
                                        }
                                    default:
                                        var v74 int = (v50 ^ (v50 >> ((111) & 63)))
                                        _ = v74
                                    }
                                    v1 -= v1
                                    _ = v72
                                }
                            case "a":
                                v50 += -374
                            default:
                                fmt.Println(v53, v3)
                            }
                        default:
                            if true {
                                v2[152] = (false || true)
                                if false {
                                    v75 := GenS3{F0: true}
                                    switch ("hello" + "foo bar") {
                                    case "hello":
                                        switch v50 {
                                        case -1:
                                            if true {
                                                for fi76 := 0; fi76 < 8; fi76++ {
                                                    v1 = v1
                                                    var v77 int32 = v53
                                                    _ = fi76
                                                    _ = v77
                                                }
                                                _ = len(v2)
                                            }
                                        case 9223372036854775807:
                                            var v78 int8 = v5
                                            _ = v78
                                        }
                                    }
                                    _ = v75
                                }
                                for fi79 := 0; fi79 < 4; fi79++ {
                                    var v80 float64 = 1e10
                                    for fi81 := 0; fi81 < 4; fi81++ {
                                        switch "" {
                                        case "a":
                                            var v82 int8 = (-(v5))
                                            _ = v82
                                        case "hello":
                                            var v83 float32 = v49
                                            _ = v83
                                        }
                                        _ = len(v4)
                                        var v84 uint32 = 221
                                        _ = fi81
                                        _ = v84
                                    }
                                    switch ("hello" + "hello") {
                                    case "\n\t":
                                        for fi85 := 0; fi85 < 9; fi85++ {
                                            fmt.Println(v50)
                                            _ = fi85
                                        }
                                    case "hello":
                                        switch 0 {
                                        case 127:
                                            for fi86 := 0; fi86 < 13; fi86++ {
                                                v1 += v1
                                                _ = len(v4)
                                                _ = fi86
                                            }
                                        case 42:
                                            fmt.Println(v5)
                                        }
                                    default:
                                        for fi87 := 0; fi87 < 16; fi87++ {
                                            v80 = v80
                                            v88 := GenS6{F0: "", F1: 1, F2: 1000}
                                            _ = fi87
                                            _ = v88
                                        }
                                    }
                                    _ = fi79
                                    _ = v80
                                }
                            }
                        }
                    } else {
                        _ = len(v2)
                        for fi89 := 0; fi89 < 2; fi89++ {
                            var v90 uint = 65535
                            var v91 uint16 = v1
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                        v2[v50] = ((false || (v1 >= (^(v1)))) || true)
                    }
                }
                v92 := v2[(v50 - (v50 * (v50 << ((1) & 63))))]
                v1 = 42
                _ = v92
            }
        case "hello":
            fmt.Println(v49, v5, v52)
        }
    } else {
        v3 -= v3
        v5 += v5
        v3 = v3
    }
    var v93 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v93
}
