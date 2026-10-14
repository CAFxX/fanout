package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 int64; F2 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint; F2 int32 }
type GenS11 struct { F0 int8 }
func genFn0() uint {
    var v7 int = 42
    v7 -= v7
    _ = v7
    return 285
}

func main() {
    var v1 string = ("" + "\n\t")
    v2 := make(map[int]int)
    var v3 uint32 = 2147483647
    v4 := []int{-2147483648, -2147483648, 1, 2147483647}
    v5 := []uint32{v3, 42}
    var v6 int16 = -32768
    if (127 <= 97) {
        v1 = v1
        _ = len(v2)
    } else {
        if (126 <= 42) {
            v4 = append(v4, 9223372036854775807)
            var v8 int64 = -988
            if ((0 != 1) && true) {
                var v9 int8 = -2
                _ = v9
            }
            _ = v8
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                if false {
                    delete(v2, (fi10 ^ 0))
                } else {
                    _ = len(v2)
                }
                v6 += (v6 * v6)
                _ = len(v5)
                _ = fi10
            }
            switch -2147483648 {
            case -9223372036854775808:
                v6 -= v6
            }
            switch v1 {
            case "hello":
                var v11 float32 = ((1.5 + (-2.25 / 1.0)) + (1e10 + 3.14159))
                _ = v11
            }
        }
        var v12 float64 = 0.1
        v13 := GenS1{F0: 42, F1: false, F2: (v1 + "foo bar")}
        _ = v12
        _ = v13
    }
    switch "" {
    case "foo bar":
        if (v6 > -1) {
            fmt.Println(v1)
            v6 += (v6 * (v6 / (((v6 << ((42) & 63))) | 1)))
            switch -9223372036854775808 {
            case 1:
                var v14 float64 = 0.0
                _ = v14
            case 641:
                v15 := v2[32767]
                _ = v15
            case -9223372036854775808:
                v16 := GenS1{F0: 1, F1: false, F2: ("a" + v1)}
                _ = v16
            }
        }
    case "hello":
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 uint8 = 42
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v3, v6, fi19)
        _ = fi19
    }
    switch -961 {
    case 32767:
        v20 := GenS2{F0: 2, F1: 65535, F2: 0}
        _ = v20
    default:
        v21 := GenS3{F0: v3}
        _ = v21
    }
    var v22 uint16 = 637
    for fi23 := 0; fi23 < 2; fi23++ {
        if (v6 == v6) {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 uint = 255
                var v26 uint = 65535
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            var v27 string = ""
            _ = v27
        }
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            var v28 uint8 = 1
            _ = v28
        }
        _ = fi23
    }
    _ = len(v2)
    var v29 uint32 = v3
    _ = len(v5)
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 int = (fi30 ^ (fi30 & fi30))
        var v32 uint64 = 4294967295
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v29, v22, v29)
    fmt.Println(v22)
    v1 = (v1 + "foo bar")
    var v33 int32 = -328
    switch v1 {
    case "":
        fmt.Println(v33)
    case "a":
        switch 1 {
        case -561:
            if len(v5) > 0 {
                v5[0] = (v3 * v3)
            }
        case 61:
            v34 := GenS4{F0: -47, F1: v33}
            _ = v34
        default:
            for fi35 := 0; fi35 < 10; fi35++ {
                v22 = v22
                var v36 int16 = v6
                var v37 int8 = int8(v22)
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    case "\n\t":
        v38 := GenS4{F0: 1, F1: (v33 - (v33 | (v33 ^ (v33 | -1))))}
        _ = v38
    default:
        v39 := GenS5{F0: v6, F1: 0}
        _ = v39
    }
    if false {
        var v40 int = 0
        for fi41 := 0; fi41 < 18; fi41++ {
            switch v40 {
            case 0:
                var v42 int8 = -128
                _ = v42
            case 128:
                if (1e10 >= 1.5) {
                    if (0 != 128) {
                        for fi43 := 0; fi43 < 11; fi43++ {
                            v1 = v1
                            if (-1 != 1) {
                                switch v1 {
                                case "hello":
                                    v5 = append(v5, (v3 * 276))
                                case "\n\t":
                                    fmt.Println(fi43, v1)
                                }
                                if true {
                                    v3 = v3
                                    var v44 float32 = (1.5 / -2.25)
                                    fmt.Println(v3, v29, v29)
                                    _ = v44
                                } else {
                                    v29 = v29
                                    v1 = v1
                                }
                                v45 := GenS1{F0: (v33 / ((v33) | 1)), F1: ((3.14159 - (float32((0.1 / (((((1e10 * 3.14159) + -2.25) + 1e10) / 1.0) * 1.5))) / float32(float64(((1.5 + (3.14159 - 3.14159)) / 3.14159))))) >= -2.25), F2: (v1 + v1)}
                                _ = v45
                            } else {
                                if true {
                                    var v46 string = v1
                                    switch v1 {
                                    case "a":
                                        switch v46 {
                                        case "hello":
                                            switch -2147483648 {
                                            case 32768:
                                                var v47 bool = (true && (!true))
                                                _ = v47
                                            case 127:
                                                v40 = 127
                                            default:
                                                if (!true) {
                                                    switch fi41 {
                                                    case -2147483648:
                                                        var v48 int64 = -364
                                                        _ = v48
                                                    case 312:
                                                        if false {
                                                            fmt.Println(v46, fi43, v3)
                                                        } else {
                                                            var v49 int32 = (v33 << ((18446744073709551615) & 63))
                                                            _ = v49
                                                        }
                                                        fmt.Println(v33, v33)
                                                        v33 -= v33
                                                    case 0:
                                                        v40 = fi43
                                                    }
                                                    var v50 int32 = v33
                                                    fmt.Println(v22, fi43, v22)
                                                    _ = v50
                                                } else {
                                                    v40 += v40
                                                }
                                            }
                                        case "foo bar":
                                            if len(v5) > 0 {
                                                v5[0] = 1
                                            }
                                        case "\n\t":
                                            _ = len(v4)
                                        }
                                    default:
                                        if true {
                                            v51 := GenS4{F0: -128, F1: (v33 * (^(v33)))}
                                            _ = v51
                                        } else {
                                            if (fi43 > (-(v40))) {
                                                v46 = (v46 + v46)
                                                switch (v40 - (fi41 * fi43)) {
                                                case -592:
                                                    switch 42 {
                                                    case -1:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    case 0:
                                                        if (false && ((^(v29)) < (v3 | v29))) {
                                                            var v52 uint64 = 697
                                                            _ = v52
                                                        } else {
                                                            delete(v2, fi41)
                                                        }
                                                        v53 := GenS6{F0: "\n\t"}
                                                        _ = v53
                                                    case 128:
                                                        fmt.Println(v1, v22)
                                                    default:
                                                        switch 0 {
                                                        case 0:
                                                            v40 = v40
                                                        }
                                                    }
                                                case 9223372036854775807:
                                                    if len(v4) > 0 {
                                                        v4[0] = fi43
                                                    }
                                                default:
                                                    v1 = "hello"
                                                }
                                            } else {
                                                if true {
                                                    v29 -= (v29 ^ (v3 | (v29 >> ((42) & 63))))
                                                    var v54 int = fi41
                                                    var v55 int = 376
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    _ = len(v2)
                                                    v33 = (v33 % (((-(v33))) | 1))
                                                    v40 = -583
                                                    v3 = (v29 / ((508) | 1))
                                                }
                                                var v56 uint16 = (^(v22))
                                                _ = len(v2)
                                                _ = v56
                                            }
                                            if true {
                                                v57 := GenS5{F0: 853, F1: 42}
                                                _ = v57
                                            }
                                        }
                                    }
                                    if (!false) {
                                        switch "" {
                                        case "hello":
                                            var v58 int = (fi43 / ((fi41) | 1))
                                            _ = v58
                                        case "\n\t":
                                            v1 = (("foo bar" + (v46 + (v46 + (v1 + "hello")))) + v1)
                                        default:
                                            fmt.Println(v1)
                                        }
                                        var v59 float32 = float32(1.5)
                                        if ((1.5 + 1e10) <= 3.14159) {
                                            switch v1 {
                                            case "foo bar":
                                                v22 = v22
                                            case "a":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "":
                                                v1 = v46
                                            }
                                        } else {
                                            var v60 float64 = float64(v59)
                                            switch fi41 {
                                            case -2147483648:
                                                v46 = v46
                                            case 127:
                                                if (!false) {
                                                    fmt.Println(fi43)
                                                    if false {
                                                        var v61 uint32 = (^(v29))
                                                        v29 = 42
                                                        v2[(fi43 >> ((545) & 63))] = fi41
                                                        _ = v61
                                                    } else {
                                                        v62 := GenS3{F0: (^(v3))}
                                                        switch "hello" {
                                                        case "\n\t":
                                                            fmt.Println(v6)
                                                        case "":
                                                            v1 = (v46 + "a")
                                                        default:
                                                            if true {
                                                                switch fi43 {
                                                                case 128:
                                                                    var v63 string = "\n\t"
                                                                    _ = v63
                                                                case -893:
                                                                    v46 = (v1 + v46)
                                                                case 32768:
                                                                    if (false || true) {
                                                                        v3 -= v29
                                                                    } else {
                                                                        v60 += (-2.25 + v60)
                                                                        v64 := GenS7{F0: v1, F1: 0, F2: (^(v29))}
                                                                        _ = v64
                                                                    }
                                                                }
                                                                switch -1 {
                                                                case -1:
                                                                    var v65 bool = true
                                                                    _ = v65
                                                                case -9223372036854775808:
                                                                    var v66 float64 = (v60 + 3.14159)
                                                                    _ = v66
                                                                case 677:
                                                                    var v67 uint16 = 43
                                                                    _ = v67
                                                                }
                                                                switch v1 {
                                                                case "a":
                                                                    v6 += (v6 - v6)
                                                                case "hello":
                                                                    v1 = "\n\t"
                                                                case "":
                                                                    var v68 uint8 = 128
                                                                    _ = v68
                                                                default:
                                                                    fmt.Println(v1)
                                                                    v69 := GenS8{F0: 2147483647}
                                                                    v5 = append(v5, v3)
                                                                    _ = v69
                                                                }
                                                            } else {
                                                                switch (-(fi43)) {
                                                                case 127:
                                                                    if (false || ("foo bar" >= (((v1 + v46) + "hello") + "a"))) {
                                                                        v70 := GenS9{F0: (v22 | (v22 & (v22 & v22))), F1: uint(v22), F2: 559}
                                                                        var v71 int64 = -208
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                case -974:
                                                                    v59 = (((float32(v60) - v59) * float32(v60)) * (float32(v60) + (((float32(v60) / (((((v59 * v59) / v59) * v59) - v59) / v59)) + v59) / v59)))
                                                                }
                                                                var v72 float32 = (1.5 - float32(float64(((0.1 + v59) - float32(1.5)))))
                                                                _ = v72
                                                            }
                                                        }
                                                        _ = len(v5)
                                                        v73 := GenS10{F0: 315, F1: 2, F2: (v33 << ((4294967295) & 63))}
                                                        if true {
                                                            var v74 uint16 = v22
                                                            v22 -= v22
                                                            if (((42 <= 100) && (true || (-128 != 42))) && (!true)) {
                                                                var v75 uint64 = 42
                                                                if true {
                                                                    if (!false) {
                                                                        v76 := GenS2{F0: 2, F1: (v22 - (v22 ^ v22)), F2: 42}
                                                                        _ = v76
                                                                    } else {
                                                                        v46 = "a"
                                                                        var v77 uint16 = (v74 / (((v74 ^ (v74 >> ((42) & 63)))) | 1))
                                                                        var v78 int8 = -94
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                } else {
                                                                    var v79 float64 = (v60 - 1.5)
                                                                    _ = v79
                                                                }
                                                                var v80 int = fi41
                                                                _ = v75
                                                                _ = v80
                                                            }
                                                            _ = v74
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v81 int64 = 0
                                                        var v82 uint64 = 4294967295
                                                        var v83 int8 = 127
                                                        switch v40 {
                                                        case 32767:
                                                            var v84 int32 = v33
                                                            _ = v84
                                                        }
                                                        _ = v62
                                                        _ = v73
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                }
                                            }
                                            _ = v60
                                        }
                                        _ = v59
                                    }
                                    _ = v46
                                } else {
                                    v85 := GenS11{F0: 0}
                                    _ = v85
                                }
                                var v86 int16 = (-(v6))
                                _ = len(v4)
                                _ = v86
                            }
                            _ = fi43
                        }
                    }
                    v22 -= v22
                } else {
                    v3 = (v3 >> ((42) & 63))
                }
            }
            for fi87 := 0; fi87 < 9; fi87++ {
                fmt.Println(v6)
                if (v33 < v33) {
                    switch fi41 {
                    case 4294967295:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        switch v1 {
                        case "":
                            v33 -= 42
                        case "\n\t":
                            var v88 uint32 = (v3 & (^(v3)))
                            _ = v88
                        default:
                            v6 = v6
                        }
                    }
                } else {
                    fmt.Println(v40)
                    var v89 string = "\n\t"
                    v29 -= (v29 + v29)
                    _ = v89
                }
                _ = fi87
            }
            _ = fi41
        }
        _ = v40
    }
    for fi90 := 0; fi90 < 1; fi90++ {
        var v91 int8 = 42
        v5 = append(v5, 715)
        switch 89 {
        case 2147483647:
            var v92 int16 = 493
            _ = v92
        }
        _ = fi90
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
