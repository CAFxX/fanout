package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0(p0 int64, p1 uint, p2 string) uint16 {
    v7 := []int{1, 0, -9223372036854775808, 848}
    p0 -= -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 565
}
func genFn1() uint16 {
    var v8 uint64 = 663
    for fi9 := 0; fi9 < 17; fi9++ {
        v8 -= (v8 * 1)
        _ = fi9
    }
    if (100 == -128) {
        var v10 uint8 = 127
        _ = v10
    } else {
        v8 -= (v8 % ((724) | 1))
    }
    _ = v8
    return 255
}

func main() {
    var v1 float32 = 1.5
    var v2 float64 = (-2.25 + ((-2.25 - float64(-2.25)) - 1e10))
    var v3 uint64 = 0
    var v4 bool = (-629 <= 1)
    var v5 int32 = -2147483648
    v6 := make(map[uint32]uint32)
    var v11 int16 = -176
    v12 := GenS1{F0: 9223372036854775807, F1: -2147483648, F2: v11}
    var v13 uint16 = 32767
    var v14 int32 = (v5 % (((v5 / (((v5 & 0)) | 1))) | 1))
    switch 1 {
    case 42:
        if v4 {
            if (1 == 42) {
                v1 = 3.14159
            } else {
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        if v4 {
                            var v17 int32 = (v14 * (-(v14)))
                            _ = v17
                        } else {
                            if false {
                                if (v13 != 268) {
                                    if false {
                                        if v4 {
                                            if ((0.1 * v1) > float32(1e10)) {
                                                var v18 float32 = (v1 - v1)
                                                var v19 int16 = v11
                                                if v4 {
                                                    v20 := v6[2147483648]
                                                    var v21 uint32 = v20
                                                    v22 := GenS2{F0: 591, F1: 0}
                                                    delete(v6, uint32(v3))
                                                    _ = v20
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    v23 := []int{(-(fi16)), (fi15 >> ((255) & 63))}
                                                    var v24 string = ("foo bar" + "foo bar")
                                                    _ = v23
                                                    _ = v24
                                                }
                                                v25 := []int{(fi16 << ((65535) & 63)), fi16, fi16}
                                                if v4 {
                                                    switch "foo bar" {
                                                    case "a":
                                                        _ = len(v6)
                                                    }
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                } else {
                                                    var v26 int8 = 45
                                                    v27 := GenS1{F0: 42, F1: fi15, F2: v19}
                                                    _ = v26
                                                    _ = v27
                                                }
                                                switch (fi15 * (^(fi16))) {
                                                case 32767:
                                                    var v28 uint64 = 2
                                                    _ = v28
                                                default:
                                                    fmt.Println(fi16)
                                                }
                                                var v29 int32 = -2147483648
                                                _ = v18
                                                _ = v19
                                                _ = v25
                                                _ = v29
                                            }
                                        }
                                        var v30 int16 = (-(v11))
                                        _ = v30
                                    }
                                    v6[65535] = 4294967295
                                    var v31 uint = 42
                                    _ = v31
                                }
                            }
                        }
                        fmt.Println(v14, fi16)
                        v1 = (float32((1e10 - v2)) / 1.5)
                        _ = fi16
                    }
                    v13 -= (^(v13))
                    var v32 int64 = 0
                    _ = fi15
                    _ = v32
                }
            }
        }
    default:
        fmt.Println(v4)
    }
    var v33 uint64 = (v3 ^ (v3 / (((v3 << ((1) & 63))) | 1)))
    for fi34 := 0; fi34 < 15; fi34++ {
        v6[971] = 148
        v35 := GenS2{F0: (v5 & (v5 >> ((65535) & 63))), F1: 1}
        if false {
            v36 := []int{fi34, 9223372036854775807, fi34, fi34}
            if true {
                if (100 != -1) {
                    v37 := GenS1{F0: -867, F1: int(v3), F2: v11}
                    if v4 {
                        switch ("foo bar" + ("" + "a")) {
                        case "":
                            v38 := GenS2{F0: (v5 >> ((1) & 63)), F1: -1}
                            _ = v38
                        default:
                            var v39 int8 = -32
                            _ = v39
                        }
                    }
                    var v40 int32 = v14
                    _ = v37
                    _ = v40
                }
                switch fi34 {
                case 1:
                    v41 := GenS2{F0: 42, F1: 42}
                    _ = v41
                case 127:
                    var v42 uint16 = v13
                    _ = v42
                }
            } else {
                v36 = append(v36, fi34)
                var v43 uint8 = 127
                for fi44 := 0; fi44 < 18; fi44++ {
                    if ((^(v13)) <= (v13 >> ((18446744073709551615) & 63))) {
                        if true {
                            v14 -= v5
                        }
                        v43 -= (v43 << ((18446744073709551615) & 63))
                        var v45 uint32 = 4294967295
                        _ = v45
                    } else {
                        fmt.Println(v33, v3, v33)
                    }
                    var v46 int8 = -1
                    _ = fi44
                    _ = v46
                }
                _ = v43
            }
            if (v4 || v4) {
                switch "foo bar" {
                case "foo bar":
                    v13 -= genFn1()
                case "hello":
                    v47 := GenS2{F0: v14, F1: 13}
                    _ = v47
                default:
                    fmt.Println(v13, v2, v4)
                }
            }
            _ = v36
        } else {
            if (v4 && v4) {
                fmt.Println(v13, fi34)
                var v48 bool = (-1 != int64(v11))
                var v49 uint32 = 4294967295
                _ = v48
                _ = v49
            }
            for fi50 := 0; fi50 < 12; fi50++ {
                v14 -= v14
                fmt.Println(v1)
                if v4 {
                    fmt.Println(v3, v14, v13)
                    v13 += (v13 % ((v13) | 1))
                }
                _ = fi50
            }
            if ((v4 || false) || (!(v4 || (4294967295 == 0)))) {
                var v51 float64 = ((float64(0.1) / (v2 + 1.5)) - v2)
                for fi52 := 0; fi52 < 5; fi52++ {
                    if false {
                        var v53 uint8 = uint8(v13)
                        var v54 int = (fi34 / ((fi34) | 1))
                        v13 += v13
                        _ = v53
                        _ = v54
                    }
                    v14 = (v14 & (^(v14)))
                    v13 = (v13 >> ((4294967295) & 63))
                    _ = fi52
                }
                v5 = (-(v14))
                _ = v51
            } else {
                v11 = v11
                for fi55 := 0; fi55 < 3; fi55++ {
                    v11 -= v11
                    v56 := v6[4294967295]
                    v57 := []int{(fi34 / ((32768) | 1))}
                    _ = fi55
                    _ = v56
                    _ = v57
                }
            }
        }
        _ = fi34
        _ = v35
    }
    fmt.Println(v14, v1)
    var v58 bool = (112 >= 0)
    v14 -= (v14 / (((v5 | (v14 ^ 2147483647))) | 1))
    v6[631] = 226
    for fi59 := 0; fi59 < 7; fi59++ {
        var v60 int64 = 525
        _ = fi59
        _ = v60
    }
    switch 1 {
    case -293:
        var v61 uint16 = v13
        _ = v61
    }
    v62 := v6[42]
    for fi63 := 0; fi63 < 8; fi63++ {
        v13 += (v13 * genFn1())
        if (!false) {
            v64 := []int{fi63}
            var v65 float64 = (v2 / (1e10 - v2))
            _ = v64
            _ = v65
        } else {
            for fi66 := 0; fi66 < 10; fi66++ {
                v67 := []int{(fi66 >> ((880) & 63))}
                if ("\n\t" != "foo bar") {
                    fmt.Println(v2, v13, v62)
                    v68 := GenS2{F0: (^(v14)), F1: -9223372036854775808}
                    var v69 int = (^(fi63))
                    _ = v68
                    _ = v69
                } else {
                    var v70 bool = (!v58)
                    if len(v67) > 0 {
                        v67[0] = (fi66 + (fi66 / (((fi66 + (fi63 % ((-283) | 1)))) | 1)))
                    }
                    v71 := GenS1{F0: -9223372036854775808, F1: 32768, F2: 0}
                    _ = v70
                    _ = v71
                }
                _ = len(v67)
                _ = fi66
                _ = v67
            }
        }
        _ = fi63
    }
    v72 := GenS2{F0: v5, F1: 9223372036854775807}
    v11 -= 1000
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v62
    _ = v72
}
