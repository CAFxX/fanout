package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int32; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 int8; F2 int16 }
type GenS6 struct { F0 int; F1 uint32; F2 int64 }
type GenS7 struct { F0 uint64; F1 uint16 }
func genFn0(p0 bool, p1 uint8) int16 {
    if true {
        var v6 int32 = 0
        _ = v6
    } else {
        var v7 uint = 42
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint16 = 42
        v10 := []int{fi8, fi8, fi8, fi8}
        for fi11 := 0; fi11 < 9; fi11++ {
            if ((float64(3.14159) * ((0.1 + -2.25) * (0.0 * (0.0 * 1.5)))) != 0.1) {
                if len(v10) > 0 {
                    v10[0] = (fi8 << ((18446744073709551615) & 63))
                }
            }
            fmt.Println(p1, fi8)
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v12 := GenS1{F0: -887, F1: ("hello" + "")}
    p1 = 2
    for fi13 := 0; fi13 < 17; fi13++ {
        switch "" {
        case "":
            p1 = p1
        default:
            for fi14 := 0; fi14 < 18; fi14++ {
                p1 -= (p1 ^ p1)
                v15 := []int{-1, 32767, fi14}
                var v16 uint8 = p1
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        p1 = (p1 ^ (p1 + p1))
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    return -1000
}
func genFn1(p0 uint32) uint64 {
    v17 := GenS1{F0: 977, F1: ("a" + "")}
    var v18 uint32 = p0
    var v19 int64 = -1
    for fi20 := 0; fi20 < 9; fi20++ {
        switch fi20 {
        case -141:
            var v21 float32 = (0.1 + (((1e10 + -2.25) * 0.0) / 1.0))
            _ = v21
        default:
            var v22 float64 = (-2.25 - (float64((((0.0 / -2.25) / 1e10) - 1e10)) - (0.0 + 1e10)))
            _ = v22
        }
        fmt.Println(v19)
        _ = fi20
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 532
}

func main() {
    var v1 int64 = 216
    var v2 int8 = 127
    var v3 uint8 = 128
    v4 := make(map[int]string)
    v5 := []string{"", "\n\t"}
    if ((990 != 772) || ((-2.25 - -2.25) == 1.5)) {
        var v23 string = ""
        if (!false) {
            for fi24 := 0; fi24 < 14; fi24++ {
                fmt.Println(fi24, v1, v1)
                var v25 int8 = v2
                var v26 uint = uint(v3)
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            var v27 uint = 75
            if false {
                if true {
                    var v28 uint32 = 2147483648
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 8; fi29++ {
                        fmt.Println(v1, v23)
                        var v30 float64 = 3.14159
                        fmt.Println(v3, v3)
                        _ = fi29
                        _ = v30
                    }
                    var v31 int8 = -76
                    v5 = append(v5, "\n\t")
                    _ = v31
                }
            } else {
                if (-800 < 1) {
                    switch "hello" {
                    case "hello":
                        v32 := GenS1{F0: -9223372036854775808, F1: v23}
                        _ = v32
                    case "foo bar":
                        v27 = v27
                    default:
                        var v33 float32 = (0.0 - 1.5)
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 9; fi34++ {
                        v35 := GenS2{F0: 658, F1: 0, F2: (v1 * (v1 ^ (v1 << ((v27) & 63))))}
                        delete(v4, (fi34 | (fi34 % ((fi34) | 1))))
                        if (0.0 >= 0.0) {
                            switch (v23 + ("foo bar" + ("foo bar" + v23))) {
                            case "":
                                v27 = 1
                            case "foo bar":
                                for fi36 := 0; fi36 < 9; fi36++ {
                                    if false {
                                        v3 += v3
                                        var v37 bool = (!(!(((v3 ^ 255) >= (v3 % (((v3 - v3)) | 1))) && true)))
                                        _ = v37
                                    } else {
                                        var v38 int32 = -2147483648
                                        fmt.Println(v1, fi36, v23)
                                        _ = v38
                                    }
                                    _ = fi36
                                }
                            case "\n\t":
                                fmt.Println(v23, fi34, v23)
                            }
                            v39 := GenS2{F0: 1, F1: 0, F2: (-(v1))}
                            for fi40 := 0; fi40 < 8; fi40++ {
                                if (!false) {
                                    fmt.Println(v1)
                                    if (v27 > v27) {
                                        v23 = ""
                                        if (2 != 2) {
                                            fmt.Println(fi34)
                                        } else {
                                            v41 := GenS2{F0: 2147483647, F1: 42, F2: -1}
                                            var v42 int = (fi34 | 42)
                                            _ = v41
                                            _ = v42
                                        }
                                    } else {
                                        v43 := GenS3{F0: (v3 >> ((152) & 63))}
                                        _ = v43
                                    }
                                } else {
                                    delete(v4, fi34)
                                }
                                var v44 int8 = (v2 + v2)
                                fmt.Println(v27, v3)
                                v2 -= (^(v44))
                                _ = fi40
                                _ = v44
                            }
                            _ = v39
                        }
                        _ = fi34
                        _ = v35
                    }
                } else {
                    delete(v4, -1)
                    var v45 int32 = -2147483648
                    v2 = -30
                    _ = v45
                }
                var v46 float64 = 1.5
                _ = v46
            }
            _ = v27
        }
        _ = v23
    }
    fmt.Println(v1)
    _ = len(v5)
    switch -128 {
    case 2147483647:
        v3 = (v3 & v3)
    case -224:
        for fi47 := 0; fi47 < 18; fi47++ {
            var v48 bool = (-2.25 != 3.14159)
            _ = fi47
            _ = v48
        }
    default:
        fmt.Println(v3, v2, v1)
    }
    var v49 string = "foo bar"
    switch 0 {
    case 127:
        v2 -= (v2 + v2)
    }
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    v50 := GenS4{F0: (v2 * (v2 | (v2 >> ((725) & 63))))}
    for fi51 := 0; fi51 < 17; fi51++ {
        v52 := GenS5{F0: 42, F1: (v2 / ((v2) | 1)), F2: 42}
        _ = fi51
        _ = v52
    }
    fmt.Println(v2, v1, v2)
    v53 := GenS6{F0: 9223372036854775807, F1: 2147483648, F2: v1}
    v2 += (v2 / ((v2) | 1))
    switch v49 {
    case "\n\t":
        v54 := GenS1{F0: -594, F1: (v49 + v49)}
        _ = v54
    }
    _ = len(v4)
    if (true || false) {
        _ = len(v5)
        if false {
            v55 := GenS7{F0: 18446744073709551615, F1: 666}
            if false {
                for fi56 := 0; fi56 < 20; fi56++ {
                    if (65535 >= 30) {
                        v1 = (^(v1))
                        var v57 float64 = (1e10 / 1e10)
                        var v58 uint = 18446744073709551615
                        _ = v57
                        _ = v58
                    } else {
                        v49 = ("foo bar" + v49)
                        v59 := GenS1{F0: fi56, F1: ((v49 + v49) + v49)}
                        for fi60 := 0; fi60 < 18; fi60++ {
                            _ = len(v4)
                            _ = fi60
                        }
                        _ = v59
                    }
                    _ = fi56
                }
            } else {
                v3 += 255
                if false {
                    var v61 float64 = 0.0
                    v49 = "a"
                    v61 = (v61 - v61)
                    _ = v61
                }
                fmt.Println(v2, v1, v3)
            }
            var v62 uint = 885
            _ = v55
            _ = v62
        }
    } else {
        for fi63 := 0; fi63 < 19; fi63++ {
            var v64 uint = 18446744073709551615
            var v65 uint16 = 141
            _ = fi63
            _ = v64
            _ = v65
        }
        v66 := GenS6{F0: 32768, F1: 2147483648, F2: v1}
        _ = v66
    }
    var v67 float64 = -2.25
    v49 = v49
    fmt.Println(v49, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v67
}
