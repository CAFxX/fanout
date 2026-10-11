package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int64; F1 uint; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 uint32 }
func genFn0(p0 int64, p1 int8, p2 uint32) int8 {
    v7 := []int{2147483647}
    p1 = -23
    if false {
        var v8 int64 = -377
        _ = v8
    }
    v7 = append(v7, -2147483648)
    v9 := GenS1{F0: 128, F1: 28, F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return (p1 & p1)
}
func genFn1(p0 int64) int8 {
    v10 := []int{-9223372036854775808, 100, 42, -76}
    p0 -= 42
    _ = p0
    _ = v10
    return -1
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 255
    v3 := []int{127, -629, -1, 128}
    var v4 int8 = 0
    var v5 uint = uint(v2)
    v6 := []int{32768, -355, -992}
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int16 = int16(v1)
        v13 := GenS1{F0: (^(v2)), F1: v1, F2: (false && ((v12 % (((v12 + (v12 - (-(v12))))) | 1)) != (v12 ^ v12)))}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: v2, F1: (v1 >> ((v5) & 63)), F2: (!true)}
    var v15 int = -9223372036854775808
    v15 = (v15 / ((-9223372036854775808) | 1))
    var v16 int64 = 63
    v5 += (v5 << (((^(v5))) & 63))
    for fi17 := 0; fi17 < 11; fi17++ {
        v2 = v2
        switch "foo bar" {
        case "hello":
            if true {
                v18 := GenS1{F0: v2, F1: 42, F2: true}
                if true {
                    fmt.Println(v15, v2, v5)
                    var v19 uint64 = 0
                    v20 := GenS1{F0: v2, F1: v1, F2: true}
                    _ = v19
                    _ = v20
                }
                v5 -= (v5 + (v5 + 4294967295))
                _ = v18
            } else {
                var v21 int16 = -1
                var v22 uint32 = 100
                var v23 uint64 = 1
                _ = v21
                _ = v22
                _ = v23
            }
        default:
            v1 -= 1
        }
        switch "a" {
        case "\n\t":
            var v24 float64 = ((0.1 + 1.5) + (1.5 / 1.0))
            _ = v24
        }
        _ = fi17
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 int8 = (v4 % ((v4) | 1))
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    fmt.Println(v1, v4, v2)
    if (false || true) {
        if (2 == 562) {
            var v28 int64 = v16
            _ = v28
        } else {
            _ = len(v3)
        }
    } else {
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 bool = (255 >= v5)
            fmt.Println(v4, v16, fi29)
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            v3 = append(v3, (fi31 ^ (v15 * (v15 << ((255) & 63)))))
            if len(v6) > 0 {
                v6[0] = 128
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi31
        }
    }
    v32 := GenS1{F0: (v2 & v2), F1: (v1 | (v1 & (v1 * (v1 + v1)))), F2: (("" + "a") < ("" + "hello"))}
    fmt.Println(v5, v15)
    if len(v6) > 0 {
        v6[0] = int(v2)
    }
    var v33 string = ""
    switch (v15 * v15) {
    case -824:
        var v34 float32 = 3.14159
        _ = v34
    case 42:
        v35 := GenS1{F0: (v2 | 26), F1: 0, F2: false}
        _ = v35
    case -9223372036854775808:
        for fi36 := 0; fi36 < 12; fi36++ {
            v15 -= fi36
            v6 = append(v6, fi36)
            v37 := GenS2{F0: -9223372036854775808, F1: 1000}
            _ = fi36
            _ = v37
        }
    default:
        if true {
            v6 = append(v6, -171)
            for fi38 := 0; fi38 < 1; fi38++ {
                switch (fi38 + -1) {
                case 127:
                    v1 += (v1 % (((v1 | v1)) | 1))
                case 128:
                    var v39 bool = (genFn0((^(v16)), (v4 ^ v4), 65535) == -56)
                    _ = v39
                case 666:
                    switch -1 {
                    case -814:
                        if ("\n\t" < v33) {
                            if false {
                                if (1e10 < 0.0) {
                                    v33 = (v33 + v33)
                                } else {
                                    switch "hello" {
                                    case "":
                                        if false {
                                            switch "foo bar" {
                                            case "":
                                                if false {
                                                    v40 := GenS3{F0: (v16 % ((v16) | 1)), F1: 4294967295, F2: 4294967295}
                                                    switch (v33 + "") {
                                                    case "\n\t":
                                                        var v41 int64 = (v16 / ((v16) | 1))
                                                        _ = v41
                                                    case "foo bar":
                                                        var v42 int8 = int8(v1)
                                                        _ = v42
                                                    default:
                                                        var v43 bool = (((-2147483648 != -971) && false) && (1e10 >= 0.0))
                                                        _ = v43
                                                    }
                                                    _ = v40
                                                }
                                            case "a":
                                                for fi44 := 0; fi44 < 19; fi44++ {
                                                    v15 = (v15 - 698)
                                                    _ = fi44
                                                }
                                            case "hello":
                                                fmt.Println(v5)
                                            }
                                            v1 -= (v1 / (((v1 / (((^(v1))) | 1))) | 1))
                                        } else {
                                            for fi45 := 0; fi45 < 4; fi45++ {
                                                switch fi45 {
                                                case 4294967295:
                                                    v4 = (v4 - v4)
                                                case -9223372036854775808:
                                                    switch (fi38 ^ (-(v15))) {
                                                    case -361:
                                                        var v46 int16 = -32768
                                                        _ = v46
                                                    }
                                                    var v47 float32 = 0.1
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    _ = v47
                                                case -190:
                                                    switch v33 {
                                                    case "foo bar":
                                                        if ((((1e10 - float64(float32(0.0))) + 0.0) + float64(0.0)) < (-2.25 + -2.25)) {
                                                            var v48 string = "a"
                                                            _ = v48
                                                        } else {
                                                            v49 := GenS2{F0: (v16 % (((-(v16))) | 1)), F1: 42}
                                                            if false {
                                                                v50 := GenS3{F0: (v16 >> ((v5) & 63)), F1: v5, F2: 888}
                                                                _ = v50
                                                            } else {
                                                                v2 = (v2 + (v2 * 1))
                                                                var v51 uint8 = (v2 >> (((^(v5))) & 63))
                                                                v16 += 952
                                                                v15 = 4294967295
                                                                _ = v51
                                                            }
                                                            var v52 int8 = genFn0(int64(v2), v4, 0)
                                                            _ = v49
                                                            _ = v52
                                                        }
                                                    case "\n\t":
                                                        var v53 int64 = (v16 >> ((v5) & 63))
                                                        _ = v53
                                                    case "":
                                                        v3 = append(v3, (v15 >> ((v5) & 63)))
                                                    default:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v4 += (v4 + v4)
                                                default:
                                                    switch v15 {
                                                    case -1:
                                                        v5 += uint(v2)
                                                    case 2147483647:
                                                        switch v33 {
                                                        case "":
                                                            fmt.Println(v5)
                                                        case "\n\t":
                                                            var v54 uint8 = (v2 % (((v2 >> ((uint(v2)) & 63))) | 1))
                                                            _ = v54
                                                        case "hello":
                                                            fmt.Println(v4, v5, v33)
                                                        }
                                                    case -9223372036854775808:
                                                        if (-1 != 32767) {
                                                            v55 := GenS3{F0: v16, F1: v5, F2: 1}
                                                            if (767 == 1) {
                                                                var v56 int8 = (v4 / ((v4) | 1))
                                                                _ = v56
                                                            } else {
                                                                fmt.Println(v2, v15)
                                                                fmt.Println(fi38, v4, v5)
                                                            }
                                                            _ = v55
                                                        }
                                                    }
                                                }
                                                v57 := GenS2{F0: -927, F1: 848}
                                                var v58 bool = ((v15 ^ 753) <= fi45)
                                                _ = fi45
                                                _ = v57
                                                _ = v58
                                            }
                                            if ((v1 - v1) != 230) {
                                                fmt.Println(v2, v5)
                                                for fi59 := 0; fi59 < 4; fi59++ {
                                                    if (18446744073709551615 == v5) {
                                                        v15 = (^(v15))
                                                        var v60 uint = (^(v5))
                                                        var v61 bool = false
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        v62 := GenS3{F0: 1, F1: (v5 + v5), F2: 42}
                                                        _ = v62
                                                    }
                                                    _ = fi59
                                                }
                                                switch (v15 >> ((v5) & 63)) {
                                                case 1:
                                                    var v63 uint32 = 1
                                                    _ = v63
                                                case 737:
                                                    var v64 int = 9223372036854775807
                                                    _ = v64
                                                default:
                                                    var v65 int16 = 370
                                                    _ = v65
                                                }
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                    case "hello":
                                        _ = len(v3)
                                    case "foo bar":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                                var v66 string = "a"
                                for fi67 := 0; fi67 < 20; fi67++ {
                                    v33 = ("foo bar" + (v33 + "foo bar"))
                                    _ = fi67
                                }
                                _ = v66
                            } else {
                                fmt.Println(v2)
                                v6 = append(v6, fi38)
                            }
                        }
                    case 32768:
                        var v68 uint8 = 177
                        _ = v68
                    }
                default:
                    v69 := GenS4{F0: -947, F1: v1, F2: 612}
                    _ = v69
                }
                _ = fi38
            }
        }
    }
    v16 = (v16 ^ 0)
    v33 = v33
    v3 = append(v3, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
