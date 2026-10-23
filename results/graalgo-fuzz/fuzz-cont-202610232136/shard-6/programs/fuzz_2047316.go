package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int32; F1 int8; F2 uint64 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int; F1 string; F2 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int64; F2 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32; F1 uint }
type GenS11 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint16, p1 uint32, p2 int64) uint64 {
    v7 := []int{4294967295, 9223372036854775807, 2147483647, -9223372036854775808}
    if false {
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                if len(v7) > 0 {
                    v7[0] = fi8
                }
                switch "a" {
                case "":
                    v7 = append(v7, 2147483647)
                case "hello":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    if ((-46 > 127) || (p2 < (p2 - (p2 * (p2 ^ (p2 * (p2 % ((p2) | 1)))))))) {
                        var v10 float32 = (1e10 + 0.1)
                        _ = v10
                    } else {
                        v11 := GenS1{F0: (fi8 - (fi9 - 128)), F1: false, F2: 331}
                        v12 := GenS1{F0: (fi9 ^ (fi8 >> ((589) & 63))), F1: (127 != 127), F2: (p2 * 1)}
                        v13 := GenS1{F0: (fi8 - -929), F1: false, F2: p2}
                        var v14 int16 = 1000
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                v15 := GenS2{F0: 1, F1: -762}
                _ = fi9
                _ = v15
            }
            _ = fi8
        }
        var v16 int8 = -1
        for fi17 := 0; fi17 < 13; fi17++ {
            v16 += (v16 % (((v16 & (v16 / ((int8(fi17)) | 1)))) | 1))
            v18 := GenS3{F0: -128, F1: int32(fi17)}
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        p1 = 607
        fmt.Println(p2)
        if len(v7) > 0 {
            v7[0] = 32768
        }
    }
    v7 = append(v7, -783)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"", "hello"}
    var v3 bool = (1 != 0)
    var v4 uint8 = 255
    var v5 int16 = 797
    var v6 float32 = 3.14159
    var v19 float32 = 0.0
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := GenS4{F0: -2147483648, F1: int8(fi20), F2: 18446744073709551615}
        _ = fi20
        _ = v21
    }
    var v22 string = ""
    switch 4294967295 {
    case -2147483648:
        var v23 int = 568
        _ = v23
    case 2147483647:
        for fi24 := 0; fi24 < 4; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                if (-1 >= (v5 + v5)) {
                    v19 = (0.1 + (v6 / ((((v19 - v6) * ((v19 - v6) * (((float32(-2.25) / 1.0) + ((((-2.25 - (float32(3.14159) - ((v6 - ((0.1 * v19) / (v19 + float32(float64((3.14159 - (v6 / (3.14159 - (1.5 / 3.14159))))))))) * 3.14159))) * 1.5) - v19) + ((0.0 * v6) - float32(3.14159)))) + 0.0))) / 1.5) / v6)))
                    switch (fi24 << ((255) & 63)) {
                    case 935:
                        v22 = v22
                    case 127:
                        if (((("\n\t" + v22) + "hello") + "hello") > v22) {
                            switch v22 {
                            case "hello":
                                _ = len(v1)
                            default:
                                fmt.Println(v5, fi25, v3)
                            }
                        } else {
                            v4 = v4
                        }
                    }
                } else {
                    delete(v1, fi24)
                    v26 := GenS2{F0: (^(v4)), F1: -1}
                    var v27 int32 = -940
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v26
                    _ = v27
                }
                fmt.Println(fi24, v22, v22)
                v19 -= ((v19 - 0.1) - v19)
                _ = fi25
            }
            v28 := GenS2{F0: v4, F1: 1}
            var v29 string = ("foo bar" + v22)
            _ = fi24
            _ = v28
            _ = v29
        }
    default:
        if (-73 >= 1) {
            if ((false || v3) || v3) {
                switch v22 {
                case "hello":
                    fmt.Println(v5, v4, v5)
                case "foo bar":
                    v2 = append(v2, v22)
                default:
                    fmt.Println(v6, v5)
                }
                for fi30 := 0; fi30 < 4; fi30++ {
                    v31 := GenS1{F0: (fi30 / (((^(fi30))) | 1)), F1: ((1e10 >= -2.25) && false), F2: 1}
                    v22 = "\n\t"
                    _ = fi30
                    _ = v31
                }
            } else {
                v3 = ((-128 >= 42) && (float32(-2.25) <= v6))
                v32 := GenS1{F0: 709, F1: true, F2: 1}
                v33 := GenS5{F0: v22, F1: (0 > 42)}
                _ = v32
                _ = v33
            }
        } else {
            if true {
                if true {
                    var v34 int = -2147483648
                    v35 := GenS6{F0: (v34 ^ (v34 | (v34 / ((-9223372036854775808) | 1)))), F1: v22, F2: 540}
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 16; fi36++ {
                    v37 := GenS5{F0: (v22 + ""), F1: (1 == 214)}
                    for fi38 := 0; fi38 < 19; fi38++ {
                        var v39 float64 = 0.1
                        var v40 float32 = ((0.1 + 1.5) - float32(1.5))
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    switch v22 {
                    case "a":
                        var v41 int64 = -9223372036854775808
                        _ = v41
                    case "foo bar":
                        for fi42 := 0; fi42 < 10; fi42++ {
                            v3 = v3
                            _ = fi42
                        }
                    }
                    _ = fi36
                    _ = v37
                }
                v43 := GenS7{F0: v5}
                _ = v43
            } else {
                var v44 uint64 = 4294967295
                _ = v44
            }
        }
    }
    if (858 <= 255) {
        for fi45 := 0; fi45 < 10; fi45++ {
            v46 := v1[fi45]
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi45
            _ = v46
        }
        for fi47 := 0; fi47 < 3; fi47++ {
            var v48 int = fi47
            _ = fi47
            _ = v48
        }
    } else {
        var v49 uint64 = 0
        if (!v3) {
            _ = len(v1)
        }
        if (v3 || v3) {
            v49 -= 88
            if (!false) {
                if (1 > 2147483647) {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    var v50 bool = (2147483647 == 0)
                    _ = v50
                }
                switch v22 {
                case "hello":
                    v51 := GenS8{F0: 988, F1: 1, F2: 127}
                    _ = v51
                case "a":
                    fmt.Println(v19)
                case "":
                    v4 = v4
                }
                v19 = -2.25
            } else {
                fmt.Println(v6, v4)
                switch "a" {
                case "":
                    v4 += 127
                default:
                    v19 = (0.0 / 1.5)
                }
            }
            _ = len(v2)
        }
        _ = v49
    }
    v52 := GenS9{F0: 401}
    if false {
        var v53 int64 = 1
        _ = v53
    } else {
        for fi54 := 0; fi54 < 14; fi54++ {
            v5 = v5
            switch fi54 {
            case 128:
                switch (fi54 & (fi54 | fi54)) {
                case 1:
                    var v55 float64 = -2.25
                    _ = v55
                case -797:
                    for fi56 := 0; fi56 < 6; fi56++ {
                        var v57 int = (fi56 | (fi54 % (((^(fi54))) | 1)))
                        var v58 int64 = 1
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                case -963:
                    var v59 float64 = (-2.25 / (3.14159 - 0.0))
                    _ = v59
                }
            case -733:
                v5 = v5
            case 9223372036854775807:
                var v60 uint64 = 42
                _ = v60
            }
            var v61 int64 = -1
            _ = fi54
            _ = v61
        }
        var v62 int64 = 9223372036854775807
        var v63 float64 = -2.25
        _ = v62
        _ = v63
    }
    v64 := GenS5{F0: "foo bar", F1: (!(v3 && (v5 < v5)))}
    v65 := GenS10{F0: 434, F1: 185}
    if len(v2) > 0 {
        v2[0] = ((v22 + "") + (v22 + v22))
    }
    v66 := v1[2147483647]
    v4 = (v4 & v4)
    v67 := GenS4{F0: -118, F1: 100, F2: 1}
    v68 := GenS7{F0: (v5 - (v5 & v5))}
    var v69 int64 = 719
    v70 := GenS11{F0: 619, F1: 575}
    var v71 uint8 = (v4 + (v4 % (((v4 | (v4 % (((v4 ^ v4)) | 1)))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
