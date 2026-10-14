package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int32 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 string; F1 uint8; F2 string }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16 }
func genFn0() string {
    var v4 int64 = 42
    v5 := []int{1, 788}
    if (18446744073709551615 <= 42) {
        var v6 uint8 = 0
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint32 = 2147483648
                var v10 int64 = v4
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        _ = v6
    } else {
        var v11 bool = (127 >= 127)
        if true {
            if (325 >= 65535) {
                if (458 != 4294967295) {
                    if v11 {
                        if (0 >= 4294967295) {
                            switch 0 {
                            case -1:
                                v12 := GenS1{F0: (float32(-2.25) == (float32(1.5) - 1.5)), F1: 1}
                                _ = v12
                            case 4294967295:
                                fmt.Println(v4, v4)
                            case -2147483648:
                                for fi13 := 0; fi13 < 12; fi13++ {
                                    var v14 uint16 = 819
                                    var v15 float64 = -2.25
                                    v11 = (0 > -1)
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            default:
                                v4 = v4
                            }
                        }
                        v16 := GenS2{F0: 255}
                        _ = v16
                    } else {
                        fmt.Println(v4)
                        var v17 uint64 = 0
                        fmt.Println(v4)
                        _ = v17
                    }
                }
                v18 := GenS3{F0: 0, F1: 368, F2: -2147483648}
                v19 := GenS4{F0: 504, F1: 0}
                _ = v18
                _ = v19
            }
        } else {
            var v20 string = "hello"
            v21 := GenS4{F0: 0, F1: 255}
            _ = v20
            _ = v21
        }
        fmt.Println(v11, v4)
        _ = v11
    }
    if len(v5) > 0 {
        v5[0] = 127
    }
    var v22 int32 = 42
    _ = v22
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1(p0 int8) uint64 {
    var v23 float64 = ((0.1 + -2.25) - (1.5 / (-2.25 * (3.14159 / 3.14159))))
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 uint8 = 222
        _ = fi24
        _ = v25
    }
    v26 := GenS5{F0: 32768, F1: 32768}
    if true {
        for fi27 := 0; fi27 < 1; fi27++ {
            switch ("" + ((("" + "") + "hello") + "a")) {
            case "hello":
                var v28 int8 = (p0 | p0)
                _ = v28
            }
            _ = fi27
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            if true {
                for fi30 := 0; fi30 < 14; fi30++ {
                    v31 := []int{867}
                    v23 = (v23 * (v23 * v23))
                    p0 = (-(p0))
                    var v32 int8 = -1
                    v33 := GenS2{F0: 151}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                p0 += 113
            }
            _ = fi29
        }
        p0 += (p0 % ((127) | 1))
    }
    _ = p0
    _ = v23
    _ = v26
    return 0
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    var v3 string = "hello"
    switch ("hello" + v3) {
    case "a":
        if true {
            var v34 uint32 = 65535
            _ = v34
        } else {
            for fi35 := 0; fi35 < 4; fi35++ {
                switch "a" {
                case "foo bar":
                    var v36 float32 = 0.1
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 10; fi37++ {
                        v1 += (^(v1))
                        _ = fi37
                    }
                }
                for fi38 := 0; fi38 < 6; fi38++ {
                    switch "a" {
                    case "a":
                        var v39 uint8 = 1
                        _ = v39
                    case "foo bar":
                        v40 := v2[(-(fi38))]
                        _ = v40
                    case "hello":
                        var v41 int64 = -124
                        _ = v41
                    }
                    switch "" {
                    case "hello":
                        v1 += (v1 ^ (v1 ^ genFn1(100)))
                    }
                    switch 919 {
                    case 1:
                        v3 = "foo bar"
                    case -9223372036854775808:
                        v3 = v3
                    }
                    var v42 int64 = -1
                    switch v3 {
                    case "a":
                        delete(v2, fi35)
                    case "foo bar":
                        v43 := GenS6{F0: v3, F1: 128, F2: (("\n\t" + v3) + (((v3 + v3) + v3) + v3))}
                        _ = v43
                    case "":
                        v44 := GenS1{F0: (-1 != -128), F1: 0}
                        _ = v44
                    default:
                        if false {
                            delete(v2, fi35)
                            v45 := []int{}
                            v46 := GenS4{F0: 42, F1: 255}
                            _ = v45
                            _ = v46
                        } else {
                            var v47 float64 = 1.5
                            _ = len(v2)
                            _ = v47
                        }
                        var v48 uint8 = 128
                        _ = v48
                    }
                    if ((v42 << ((2) & 63)) >= (v42 & v42)) {
                        v42 += (v42 & (v42 ^ (v42 * -46)))
                    }
                    _ = fi38
                    _ = v42
                }
                _ = fi35
            }
            for fi49 := 0; fi49 < 16; fi49++ {
                var v50 string = ("" + "")
                v3 = (v3 + v50)
                _ = fi49
                _ = v50
            }
            var v51 uint64 = v1
            _ = v51
        }
    case "":
        switch "hello" {
        case "foo bar":
            var v52 int32 = int32(v1)
            _ = v52
        }
    case "\n\t":
        fmt.Println(v3, v1)
    }
    v53 := GenS7{F0: "", F1: -642}
    var v54 int64 = 9223372036854775807
    var v55 int64 = 42
    fmt.Println(v3, v3)
    var v56 uint16 = 87
    var v57 bool = (v3 <= v3)
    if ((v55 ^ (-(v55))) <= v54) {
        var v58 int = int(v56)
        if (v57 && (0 == (^(v1)))) {
            v59 := GenS1{F0: (v57 && v57), F1: int8(v58)}
            for fi60 := 0; fi60 < 9; fi60++ {
                v61 := []int{128}
                v2[(v58 & 4294967295)] = (v58 | (v58 + v58))
                _ = fi60
                _ = v61
            }
            v54 = v55
            _ = v59
        }
        _ = v58
    }
    var v62 int32 = -1
    v63 := []int{6, 127, 185}
    var v64 int32 = v62
    v65 := GenS8{F0: "\n\t"}
    switch (v3 + v3) {
    case "a":
        if false {
            for fi66 := 0; fi66 < 5; fi66++ {
                v54 -= (v55 - v55)
                v62 += v62
                switch v3 {
                case "":
                    if v57 {
                        v62 = v62
                        v67 := GenS9{F0: 0}
                        if true {
                            switch ((v3 + v3) + (("" + v3) + v3)) {
                            case "\n\t":
                                fmt.Println(v64, v57, v54)
                            case "foo bar":
                                v3 = "foo bar"
                            case "a":
                                if v57 {
                                    v2[fi66] = fi66
                                    var v68 bool = v57
                                    _ = v68
                                }
                            }
                            if v57 {
                                var v69 int8 = 42
                                v1 += 18446744073709551615
                                _ = v69
                            } else {
                                v57 = v57
                                v62 += int32(v56)
                                var v70 uint = 760
                                _ = v70
                            }
                        } else {
                            fmt.Println(v55)
                        }
                        _ = v67
                    } else {
                        if len(v63) >= 2 {
                            v63 = v63[1:len(v63)]
                        }
                        if (true && true) {
                            v71 := GenS5{F0: (fi66 | -1), F1: fi66}
                            v56 -= 255
                            v62 = v62
                            _ = v71
                        }
                    }
                default:
                    v54 = (v55 & (v54 * (v55 + (^(v55)))))
                }
                _ = fi66
            }
            v57 = (v56 <= (v56 + v56))
        }
    case "foo bar":
        if (1 < 0) {
            var v72 int16 = 32767
            _ = v72
        }
    case "hello":
        switch "hello" {
        case "a":
            v73 := GenS5{F0: -9223372036854775808, F1: 32768}
            _ = v73
        case "":
            switch v3 {
            case "a":
                v1 -= 685
            case "\n\t":
                if len(v63) >= 2 {
                    v63 = v63[1:len(v63)]
                }
            case "foo bar":
                v64 -= (^(v62))
            default:
                var v74 uint64 = v1
                _ = v74
            }
        case "hello":
            v57 = (1e10 != 0.0)
        default:
            var v75 float64 = 1.5
            _ = v75
        }
    default:
        var v76 int8 = 121
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
