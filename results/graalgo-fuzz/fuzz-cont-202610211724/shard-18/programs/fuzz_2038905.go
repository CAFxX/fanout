package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int32 }
type GenS5 struct { F0 string; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int; F1 int; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint; F2 uint }
type GenS8 struct { F0 int8; F1 uint16; F2 uint8 }
func genFn0(p0 int64, p1 bool) int8 {
    p1 = p1
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 1, F1: -1000, F2: 42}
        if (!true) {
            v8 := GenS2{F0: 380, F1: 0, F2: p0}
            v9 := GenS3{F0: 42, F1: 42}
            _ = v8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if (false || (0.0 >= 0.0)) {
        switch "a" {
        case "hello":
            fmt.Println(p1, p1)
        case "\n\t":
            if ((1 >= 255) && p1) {
                if (915 <= 18446744073709551615) {
                    var v10 uint8 = 0
                    v11 := []int{0, 892, 2147483647, -2147483648}
                    v12 := GenS4{F0: -664, F1: 42, F2: -1}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        p0 += 0
                        _ = fi13
                    }
                    var v14 int32 = -2147483648
                    _ = v14
                }
                fmt.Println(p1, p0)
            } else {
                if (110 == (p0 * (p0 % ((0) | 1)))) {
                    var v15 int32 = -427
                    var v16 uint = 255
                    _ = v15
                    _ = v16
                }
            }
        case "a":
            if p1 {
                if true {
                    var v17 string = ("hello" + (("hello" + ("a" + "hello")) + "hello"))
                    _ = v17
                }
                fmt.Println(p0, p1)
            }
        }
        var v18 float32 = float32(1.5)
        _ = v18
    } else {
        if (true || p1) {
            switch "" {
            case "hello":
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 int = (fi19 % ((fi19) | 1))
                    fmt.Println(v20, v20, v20)
                    _ = fi19
                    _ = v20
                }
            case "a":
                var v21 int32 = 1
                _ = v21
            case "foo bar":
                for fi22 := 0; fi22 < 2; fi22++ {
                    fmt.Println(p1, p1, p0)
                    fmt.Println(p1, fi22)
                    v23 := GenS4{F0: (p0 / ((0) | 1)), F1: 128, F2: 0}
                    _ = fi22
                    _ = v23
                }
            default:
                v24 := GenS2{F0: 1, F1: 0, F2: (p0 % (((p0 << ((4294967295) & 63))) | 1))}
                _ = v24
            }
            if (float64(1e10) > 3.14159) {
                var v25 bool = ("foo bar" == "hello")
                var v26 int8 = 1
                if (!v25) {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v28 := []int{2147483647}
                        _ = fi27
                        _ = v28
                    }
                    v26 -= (v26 / ((v26) | 1))
                } else {
                    switch "\n\t" {
                    case "":
                        var v29 uint64 = 42
                        _ = v29
                    case "a":
                        for fi30 := 0; fi30 < 14; fi30++ {
                            v26 = -1
                            v25 = (1 > 2)
                            _ = fi30
                        }
                    case "foo bar":
                        if true {
                            var v31 float32 = (0.1 + 1.5)
                            for fi32 := 0; fi32 < 12; fi32++ {
                                var v33 uint8 = 128
                                fmt.Println(v25, v25)
                                _ = fi32
                                _ = v33
                            }
                            _ = v31
                        }
                    default:
                        v34 := GenS5{F0: "hello", F1: 1, F2: 475}
                        _ = v34
                    }
                    fmt.Println(p1, p1)
                    p1 = p1
                }
                _ = v25
                _ = v26
            } else {
                if false {
                    v35 := GenS3{F0: 42, F1: -32768}
                    _ = v35
                } else {
                    var v36 uint64 = 673
                    for fi37 := 0; fi37 < 19; fi37++ {
                        p1 = (0 <= 1)
                        _ = fi37
                    }
                    _ = v36
                }
            }
            fmt.Println(p0)
        }
        var v38 uint64 = 905
        v38 += (^(v38))
        _ = v38
    }
    _ = p0
    _ = p1
    return 100
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -810
    var v3 uint32 = 2
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]string)
    for fi39 := 0; fi39 < 1; fi39++ {
        if ((325 > 176) || false) {
            v4 = 810
            if true {
                var v40 uint64 = (v4 >> ((627) & 63))
                _ = v40
            }
            v4 += v4
        } else {
            v3 -= 42
        }
        for fi41 := 0; fi41 < 16; fi41++ {
            fmt.Println(v2, v3, v2)
            v42 := []int{(fi39 % (((v2 % (((fi39 + 9223372036854775807)) | 1))) | 1))}
            v3 = 0
            _ = fi41
            _ = v42
        }
        v4 -= v4
        _ = fi39
    }
    fmt.Println(v2, v3, v4)
    v3 = 65535
    if false {
        fmt.Println(v2)
        v4 -= 4294967295
        if false {
            var v43 int8 = -41
            var v44 bool = true
            v43 = (^(v43))
            _ = v43
            _ = v44
        } else {
            var v45 int64 = 9223372036854775807
            _ = v45
        }
    }
    var v46 int32 = 1
    switch "foo bar" {
    case "":
        switch "hello" {
        case "":
            v47 := v1[int(v46)]
            _ = v47
        case "hello":
            switch -2147483648 {
            case 128:
                v3 = (v3 - 655)
            case 9223372036854775807:
                for fi48 := 0; fi48 < 9; fi48++ {
                    v49 := GenS2{F0: 32767, F1: (v4 << ((uint(v4)) & 63)), F2: 1}
                    _ = fi48
                    _ = v49
                }
            default:
                var v50 uint8 = 1
                _ = v50
            }
        case "\n\t":
            v51 := GenS6{F0: (v2 ^ int(v4)), F1: v2, F2: 42}
            _ = v51
        default:
            for fi52 := 0; fi52 < 8; fi52++ {
                if (1e10 >= ((3.14159 * float32((0.0 + (-2.25 - -2.25)))) / 1.0)) {
                    v53 := GenS5{F0: (("a" + (("" + "hello") + ("\n\t" + "\n\t"))) + "foo bar"), F1: v4, F2: v4}
                    if true {
                        if (!(1 != 1)) {
                            v54 := GenS2{F0: 438, F1: v4, F2: 289}
                            _ = v54
                        }
                        var v55 int8 = -49
                        _ = v55
                    }
                    v4 = (v4 / (((v4 * (v4 % (((v4 >> ((65535) & 63))) | 1)))) | 1))
                    _ = v53
                }
                fmt.Println(v4)
                v56 := GenS4{F0: -221, F1: 42, F2: int32(v2)}
                _ = fi52
                _ = v56
            }
        }
    case "a":
        v57 := []int{(v2 - v2)}
        _ = v57
    case "\n\t":
        for fi58 := 0; fi58 < 1; fi58++ {
            if true {
                for fi59 := 0; fi59 < 8; fi59++ {
                    var v60 uint64 = 18446744073709551615
                    v3 += 576
                    _ = fi59
                    _ = v60
                }
                if (!(0 < 209)) {
                    for fi61 := 0; fi61 < 16; fi61++ {
                        delete(v5, 4294967295)
                        if false {
                            v3 = (v3 / ((v3) | 1))
                            v62 := GenS7{F0: (!false), F1: 65535, F2: 643}
                            v46 = int32(v3)
                            _ = v62
                        } else {
                            var v63 int8 = 89
                            v5[uint32(v4)] = ("" + "a")
                            _ = v63
                        }
                        switch ("hello" + "hello") {
                        case "hello":
                            fmt.Println(v4)
                        }
                        _ = fi61
                    }
                    var v64 uint16 = 987
                    delete(v5, v3)
                    _ = v64
                } else {
                    var v65 string = ("\n\t" + ("hello" + "a"))
                    v66 := GenS1{F0: 42, F1: -1, F2: v46}
                    var v67 uint = 65535
                    _ = v65
                    _ = v66
                    _ = v67
                }
                v4 -= 0
            }
            v68 := GenS8{F0: 1, F1: 0, F2: 1}
            _ = fi58
            _ = v68
        }
    }
    var v69 float32 = (float32((0.0 / -2.25)) - (-2.25 * (-2.25 / -2.25)))
    v69 += v69
    for fi70 := 0; fi70 < 14; fi70++ {
        switch (fi70 | -688) {
        case 127:
            v71 := []int{(v2 >> ((18446744073709551615) & 63)), -9223372036854775808, (v2 | fi70)}
            _ = v71
        }
        v72 := v5[uint32(v4)]
        var v73 float64 = -2.25
        _ = fi70
        _ = v72
        _ = v73
    }
    var v74 uint8 = 127
    for fi75 := 0; fi75 < 1; fi75++ {
        var v76 uint = 214
        for fi77 := 0; fi77 < 6; fi77++ {
            var v78 int64 = -9223372036854775808
            v79 := GenS4{F0: (v78 % ((v78) | 1)), F1: (v74 % ((0) | 1)), F2: (v46 % ((-421) | 1))}
            v80 := GenS2{F0: 1, F1: v4, F2: v78}
            _ = fi77
            _ = v78
            _ = v79
            _ = v80
        }
        _ = fi75
        _ = v76
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v69:", v69)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v69
    _ = v74
}
