package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int32; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }

func main() {
    var v1 string = "a"
    var v2 int64 = 1
    var v3 uint64 = 1
    var v4 int64 = v2
    v5 := make(map[int]string)
    v6 := make(map[int]string)
    v3 -= v3
    for fi7 := 0; fi7 < 18; fi7++ {
        v2 = v4
        if ((false || true) && false) {
            if (!false) {
                var v8 uint = 1
                v9 := GenS1{F0: (v8 + (^(v8))), F1: 704, F2: v8}
                if (182 >= 17) {
                    v10 := GenS1{F0: (v8 / (((v8 * v8)) | 1)), F1: 0, F2: uint(v3)}
                    _ = v10
                } else {
                    v11 := GenS1{F0: v8, F1: 1, F2: v8}
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := v6[fi12]
                        _ = fi12
                        _ = v13
                    }
                    switch v1 {
                    case "foo bar":
                        for fi14 := 0; fi14 < 19; fi14++ {
                            var v15 int8 = -1
                            var v16 int64 = 42
                            fmt.Println(fi14, v15)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    case "hello":
                        v17 := []int{(fi7 << (((v8 + 255)) & 63)), (fi7 / ((0) | 1)), fi7, (fi7 % (((fi7 - fi7)) | 1))}
                        _ = v17
                    case "":
                        delete(v6, fi7)
                    default:
                        v1 = v1
                    }
                    _ = v11
                }
                _ = v8
                _ = v9
            } else {
                if (0 != -128) {
                    v1 = v1
                }
                _ = len(v6)
                var v18 float64 = ((1e10 * 0.1) / float64(1e10))
                _ = v18
            }
            v4 = v4
        }
        switch v1 {
        case "foo bar":
            var v19 float64 = 0.0
            _ = v19
        default:
            for fi20 := 0; fi20 < 8; fi20++ {
                if ((1 < 65535) || (v1 <= "foo bar")) {
                    if ((-2.25 + ((float32((0.1 / 1.5)) + (float32(float64(-2.25)) + (3.14159 / -2.25))) + 1.5)) < -2.25) {
                        v21 := GenS1{F0: 930, F1: 42, F2: 959}
                        if false {
                            v2 = (v4 * v2)
                            var v22 uint64 = v3
                            v2 = (v4 % (((v2 + v4)) | 1))
                            _ = v22
                        } else {
                            var v23 uint64 = v3
                            _ = v23
                        }
                        _ = v21
                    } else {
                        var v24 string = v1
                        if false {
                            if (false || ((42 != 42) && false)) {
                                if true {
                                    v5[(fi20 * fi7)] = v24
                                    v25 := []int{320, 128}
                                    _ = v25
                                }
                            } else {
                                var v26 uint = 726
                                var v27 uint16 = 32767
                                if false {
                                    if (!((v27 - v27) > v27)) {
                                        v28 := GenS1{F0: 0, F1: 42, F2: v26}
                                        var v29 int = 1
                                        _ = v28
                                        _ = v29
                                    } else {
                                        v30 := []int{fi7, 4294967295}
                                        var v31 uint16 = (^(v27))
                                        _ = v30
                                        _ = v31
                                    }
                                    var v32 uint32 = 65535
                                    _ = v32
                                }
                                _ = v26
                                _ = v27
                            }
                        }
                        fmt.Println(fi7)
                        switch v1 {
                        case "":
                            var v33 float32 = (((0.1 - 0.1) * ((0.1 * 1e10) - (3.14159 / 1.5))) * 1e10)
                            _ = v33
                        case "\n\t":
                            v4 += (v4 / (((v2 << ((4294967295) & 63))) | 1))
                            fmt.Println(v3)
                        }
                        _ = v24
                    }
                    v2 = (-(v4))
                    _ = len(v5)
                }
                var v34 float64 = -2.25
                var v35 float32 = 1.5
                _ = fi20
                _ = v34
                _ = v35
            }
        }
        _ = fi7
    }
    v36 := GenS1{F0: 18446744073709551615, F1: 0, F2: 4294967295}
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 bool = (-250 > (v2 >> ((0) & 63)))
        v39 := GenS1{F0: 419, F1: 42, F2: 255}
        _ = fi37
        _ = v38
        _ = v39
    }
    if (677 != -1) {
        for fi40 := 0; fi40 < 17; fi40++ {
            var v41 int32 = 2147483647
            _ = fi40
            _ = v41
        }
    }
    v42 := GenS1{F0: 643, F1: 2147483647, F2: uint(v3)}
    v43 := GenS2{F0: (1.5 > float32(3.14159))}
    v44 := GenS3{F0: (v1 + ""), F1: 217, F2: 255}
    fmt.Println(v4, v4)
    v45 := GenS4{F0: 0, F1: -2147483648, F2: -2147483648}
    for fi46 := 0; fi46 < 2; fi46++ {
        for fi47 := 0; fi47 < 12; fi47++ {
            fmt.Println(v4, v1)
            if (false || (!true)) {
                if (325 > 42) {
                    v48 := GenS5{F0: 18446744073709551615}
                    v49 := GenS1{F0: 1, F1: 594, F2: 1}
                    _ = v48
                    _ = v49
                }
                var v50 uint32 = 188
                fmt.Println(v1, v4)
                var v51 float64 = (0.1 / 1.5)
                _ = v50
                _ = v51
            } else {
                v4 = v4
                if true {
                    v52 := GenS5{F0: 522}
                    v4 = (v2 ^ (v4 / (((v4 ^ 1)) | 1)))
                    v53 := GenS3{F0: "\n\t", F1: 2, F2: 0}
                    _ = v52
                    _ = v53
                } else {
                    switch "hello" {
                    case "a":
                        fmt.Println(fi46, v1, v4)
                    case "foo bar":
                        if false {
                            fmt.Println(v4)
                            switch (^(fi47)) {
                            case -1:
                                var v54 string = v1
                                _ = v54
                            case 42:
                                var v55 bool = false
                                _ = v55
                            case 696:
                                switch 32768 {
                                case 73:
                                    switch (v1 + v1) {
                                    case "":
                                        v3 -= 253
                                    case "\n\t":
                                        var v56 int8 = 100
                                        _ = v56
                                    case "a":
                                        v2 -= 9223372036854775807
                                    default:
                                        var v57 bool = (false && (-45 != -18))
                                        _ = v57
                                    }
                                case -768:
                                    if true {
                                        var v58 uint8 = 97
                                        switch -294 {
                                        case 252:
                                            v58 = (v58 & v58)
                                        case 561:
                                            v59 := GenS6{F0: false}
                                            _ = v59
                                        }
                                        var v60 uint32 = 1
                                        _ = v58
                                        _ = v60
                                    }
                                default:
                                    if true {
                                        v1 = v1
                                    }
                                }
                            }
                            v61 := GenS4{F0: uint8(v3), F1: 685, F2: int32(v3)}
                            v62 := GenS2{F0: (2 < 255)}
                            _ = v61
                            _ = v62
                        }
                    }
                    v63 := []int{(fi46 / (((fi46 * 127)) | 1))}
                    var v64 int16 = -810
                    v4 = 0
                    v1 = ("foo bar" + "hello")
                    _ = v63
                    _ = v64
                }
            }
            _ = fi47
        }
        _ = fi46
    }
    v4 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
