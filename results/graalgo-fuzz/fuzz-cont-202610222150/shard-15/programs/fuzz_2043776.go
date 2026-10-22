package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int }
func genFn0(p0 int16) int16 {
    if (2147483647 > 435) {
        p0 -= -1
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 1, F1: "a"}
            v8 := GenS1{F0: 42, F1: ("\n\t" + ("\n\t" + (("hello" + "") + "hello")))}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := []int{}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(p0)
    fmt.Println(p0)
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(fi11)
        var v12 uint = 987
        _ = fi11
        _ = v12
    }
    _ = p0
    return p0
}
func genFn1() uint8 {
    if (1 <= 4294967295) {
        v13 := GenS2{F0: 2, F1: 218}
        _ = v13
    } else {
        var v14 uint8 = 235
        var v15 uint64 = 4294967295
        _ = v14
        _ = v15
    }
    var v16 bool = false
    _ = v16
    return 84
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int16 = -32768
    var v3 int32 = -2147483648
    var v4 uint64 = 1
    var v5 uint = 255
    v5 -= v5
    var v17 string = ("foo bar" + "")
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 int16 = v2
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS2{F0: 174, F1: 2}
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 1, F1: (("a" + (v17 + v17)) + (("\n\t" + "a") + v17))}
    for fi23 := 0; fi23 < 13; fi23++ {
        if true {
            if false {
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 int16 = genFn0(v2)
                    _ = fi24
                    _ = v25
                }
                var v26 float32 = 0.0
                _ = v26
            }
            v5 += (v5 << ((v5) & 63))
        } else {
            v27 := []int{(fi23 ^ fi23), (fi23 | 1), (^(fi23)), (fi23 >> ((v5) & 63))}
            switch ("a" + v17) {
            case "hello":
                v28 := GenS3{F0: 255, F1: 4294967295, F2: 660}
                _ = v28
            case "\n\t":
                v1 = v1
            case "a":
                for fi29 := 0; fi29 < 9; fi29++ {
                    v1 = fi23
                    switch "foo bar" {
                    case "\n\t":
                        v2 = (v2 << ((65535) & 63))
                    case "":
                        var v30 uint32 = 65535
                        _ = v30
                    case "hello":
                        var v31 uint = v5
                        v31 += v31
                        v17 = v17
                        _ = v31
                    default:
                        fmt.Println(v5, v17)
                    }
                    fmt.Println(v1, v17)
                    v3 = (v3 ^ (^(v3)))
                    v32 := GenS4{F0: 32767}
                    _ = fi29
                    _ = v32
                }
            }
            if (0.1 >= float64((((1.5 - 1.5) - -2.25) * 3.14159))) {
                var v33 uint16 = 1
                switch v1 {
                case 9223372036854775807:
                    for fi34 := 0; fi34 < 7; fi34++ {
                        fmt.Println(v33, v1, v5)
                        switch "" {
                        case "hello":
                            v17 = "\n\t"
                        case "foo bar":
                            var v35 int16 = v2
                            _ = v35
                        case "":
                            var v36 int8 = -26
                            _ = v36
                        }
                        if (true || true) {
                            var v37 uint8 = 234
                            var v38 int16 = 1
                            _ = v37
                            _ = v38
                        } else {
                            if true {
                                fmt.Println(fi23)
                                v3 = (v3 * (^(v3)))
                            } else {
                                if true {
                                    v33 -= v33
                                }
                                var v39 int = -232
                                if (false && (18446744073709551615 > (v4 >> (((v5 | (v5 + (^(v5))))) & 63)))) {
                                    v40 := GenS1{F0: 34, F1: ("foo bar" + v17)}
                                    v3 += int32(v2)
                                    switch v17 {
                                    case "\n\t":
                                        if true {
                                            var v41 int32 = v3
                                            _ = v41
                                        } else {
                                            v1 -= (v39 >> ((v5) & 63))
                                        }
                                    case "a":
                                        v4 -= uint64(v33)
                                    case "foo bar":
                                        v17 = v17
                                    default:
                                        v4 -= (^(v4))
                                    }
                                    _ = v40
                                }
                                _ = v39
                            }
                        }
                        v3 -= v3
                        fmt.Println(v4, v2)
                        _ = fi34
                    }
                case -9223372036854775808:
                    v33 += v33
                case 32768:
                    for fi42 := 0; fi42 < 16; fi42++ {
                        if false {
                            var v43 uint32 = 2147483647
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            _ = v43
                        } else {
                            v44 := GenS1{F0: 0, F1: ("a" + (v17 + ("\n\t" + "")))}
                            if false {
                                var v45 uint8 = 0
                                if len(v27) > 0 {
                                    v27[0] = fi23
                                }
                                if (!(false || (98 != 127))) {
                                    if (v5 == v5) {
                                        v46 := GenS1{F0: -108, F1: "a"}
                                        var v47 uint32 = 2
                                        _ = v46
                                        _ = v47
                                    } else {
                                        _ = len(v27)
                                        v27 = append(v27, -56)
                                        fmt.Println(v3)
                                    }
                                }
                                _ = v45
                            }
                            v48 := GenS4{F0: 608}
                            fmt.Println(v3)
                            _ = v44
                            _ = v48
                        }
                        _ = fi42
                    }
                default:
                    var v49 float64 = ((0.0 - float64(1.5)) + 1.5)
                    _ = v49
                }
                _ = v33
            }
            _ = v27
        }
        var v50 uint = (v5 ^ 65535)
        v50 -= v5
        _ = fi23
        _ = v50
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        if (3.14159 <= 1e10) {
            v2 = 1000
            v1 = (v1 & v1)
            v5 -= (v5 >> ((v5) & 63))
        } else {
            v2 -= -1
        }
        fmt.Println(v5, v5, v2)
        _ = fi51
    }
    if (1 != 255) {
        if true {
            var v52 int16 = (-(v2))
            var v53 float32 = 0.1
            v54 := GenS1{F0: -2, F1: "foo bar"}
            _ = v52
            _ = v53
            _ = v54
        } else {
            v55 := []int{-2147483648}
            var v56 float32 = 3.14159
            _ = v55
            _ = v56
        }
    } else {
        for fi57 := 0; fi57 < 15; fi57++ {
            v58 := GenS1{F0: 100, F1: v17}
            _ = fi57
            _ = v58
        }
        v59 := []int{741, 1, (v1 << (((^(v5))) & 63))}
        var v60 uint16 = 915
        _ = v59
        _ = v60
    }
    if (-67 == 127) {
        v61 := []int{(v1 * v1)}
        _ = v61
    }
    v62 := []int{v1}
    switch (v17 + v17) {
    case "":
        v3 = (v3 + (v3 - v3))
    case "a":
        v62 = append(v62, v1)
    case "foo bar":
        v2 += 198
    }
    fmt.Println(v17, v3)
    v63 := GenS3{F0: 1, F1: v4, F2: 4294967295}
    fmt.Println(v3, v2, v2)
    for fi64 := 0; fi64 < 6; fi64++ {
        switch ("hello" + ((v17 + v17) + ("a" + v17))) {
        case "foo bar":
            v4 = (v4 << (((v5 % ((138) | 1))) & 63))
        case "hello":
            var v65 uint16 = 42
            _ = v65
        default:
            if len(v62) > 0 {
                v62[0] = (fi64 + 9223372036854775807)
            }
        }
        var v66 int = v1
        var v67 uint16 = 0
        _ = fi64
        _ = v66
        _ = v67
    }
    v68 := GenS5{F0: (^(v4)), F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v68
}
