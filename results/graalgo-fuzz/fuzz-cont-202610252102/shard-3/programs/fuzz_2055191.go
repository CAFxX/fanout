package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 int16; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 int }
type GenS5 struct { F0 string; F1 int32; F2 int }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 uint8; F1 uint16; F2 uint8 }
func genFn0(p0 string, p1 int8) int8 {
    p1 = p1
    switch 13 {
    case 0:
        var v4 bool = (961 != 65535)
        _ = v4
    case 32767:
        var v5 int32 = 0
        _ = v5
    case 58:
        v6 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808, F2: 638}
        _ = v6
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint8 = 0
    var v2 int16 = -1
    var v3 bool = true
    fmt.Println(v1, v1, v1)
    switch 32767 {
    case 128:
        switch "foo bar" {
        case "":
            v8 := []int{-517}
            _ = v8
        }
    case -619:
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 int = -2147483648
                v12 := []int{(fi10 % (((^(fi9))) | 1)), v11, v11, fi9}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    var v13 uint = 255
    v14 := GenS1{F0: v13, F1: 42, F2: 18446744073709551615}
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS2{F0: "a", F1: fi15, F2: 2}
        if (true || v3) {
            v17 := []int{0, fi15, fi15, fi15}
            if v3 {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS3{F0: 42, F1: -788, F2: (v1 >> (((^(v13))) & 63))}
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = fi18
                    _ = v19
                }
            } else {
                var v20 uint = v13
                fmt.Println(v2, v1)
                _ = v20
            }
            switch ("hello" + "\n\t") {
            case "hello":
                var v21 int64 = -9223372036854775808
                _ = v21
            }
            _ = v17
        } else {
            var v22 uint16 = 1
            fmt.Println(v1, fi15, v13)
            _ = v22
        }
        _ = fi15
        _ = v16
    }
    var v23 int64 = 9223372036854775807
    var v24 float32 = 3.14159
    switch 32767 {
    case -9223372036854775808:
        v25 := GenS4{F0: -675, F1: 631, F2: 598}
        _ = v25
    case 128:
        switch "\n\t" {
        case "a":
            v26 := GenS5{F0: "", F1: 42, F2: 1}
            _ = v26
        default:
            if (!true) {
                switch "\n\t" {
                case "foo bar":
                    if false {
                        var v27 int16 = (v2 + v2)
                        v28 := GenS4{F0: (v27 >> ((v13) & 63)), F1: 4294967295, F2: 1}
                        _ = v27
                        _ = v28
                    } else {
                        switch 1 {
                        case 2147483647:
                            v29 := []int{-812, 1, -832, 32768}
                            _ = v29
                        case -1:
                            var v30 int64 = (v23 * (v23 ^ (v23 - v23)))
                            _ = v30
                        }
                        switch 2147483647 {
                        case 127:
                            for fi31 := 0; fi31 < 1; fi31++ {
                                fmt.Println(v13, v2, v2)
                                _ = fi31
                            }
                        case 330:
                            switch -740 {
                            case 973:
                                var v32 int8 = -1
                                _ = v32
                            case -707:
                                for fi33 := 0; fi33 < 2; fi33++ {
                                    if v3 {
                                        v13 = v13
                                        var v34 uint = v13
                                        v35 := GenS6{F0: v1, F1: (3.14159 != 0.0)}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v24 += (float32(1.5) + (v24 * (v24 + ((v24 - v24) - (v24 + v24)))))
                                    }
                                    v24 = v24
                                    var v36 int8 = 42
                                    _ = fi33
                                    _ = v36
                                }
                            case 479:
                                switch ("a" + "a") {
                                case "\n\t":
                                    for fi37 := 0; fi37 < 8; fi37++ {
                                        v3 = (v2 < (v2 | (v2 & -431)))
                                        v2 = (v2 * v2)
                                        switch "hello" {
                                        case "a":
                                            v38 := GenS3{F0: 1000, F1: v23, F2: v1}
                                            _ = v38
                                        case "\n\t":
                                            v39 := []int{(-(fi37)), (fi37 & fi37)}
                                            _ = v39
                                        }
                                        _ = fi37
                                    }
                                case "a":
                                    switch 979 {
                                    case -74:
                                        v40 := GenS3{F0: (v2 * v2), F1: (v23 * -9223372036854775808), F2: v1}
                                        _ = v40
                                    default:
                                        if ((291 > 590) && true) {
                                            switch ((("\n\t" + (("a" + ("\n\t" + "")) + (("foo bar" + "") + "\n\t"))) + "foo bar") + ("" + ("hello" + "foo bar"))) {
                                            case "\n\t":
                                                if (!v3) {
                                                    var v41 int16 = v2
                                                    _ = v41
                                                }
                                            case "a":
                                                fmt.Println(v3)
                                            case "hello":
                                                v42 := GenS7{F0: (^(v1)), F1: 2}
                                                _ = v42
                                            }
                                            v43 := []int{}
                                            v13 -= (v13 << (((v13 | 42)) & 63))
                                            _ = v43
                                        } else {
                                            if (94 == 32768) {
                                                switch ((("foo bar" + "") + "") + "\n\t") {
                                                case "foo bar":
                                                    switch ("foo bar" + ("" + "")) {
                                                    case "":
                                                        for fi44 := 0; fi44 < 16; fi44++ {
                                                            if ((fi44 ^ -9223372036854775808) > fi44) {
                                                                v13 += (v13 >> ((65535) & 63))
                                                                for fi45 := 0; fi45 < 5; fi45++ {
                                                                    v23 = -206
                                                                    _ = fi45
                                                                }
                                                                if (3.14159 == -2.25) {
                                                                    var v46 uint = (^(v13))
                                                                    _ = v46
                                                                } else {
                                                                    switch (-(fi44)) {
                                                                    case -420:
                                                                        v47 := []int{fi44, (^(fi44)), (fi44 ^ -223)}
                                                                        _ = v47
                                                                    }
                                                                    if (0.0 > 1.5) {
                                                                        if v3 {
                                                                            v48 := GenS6{F0: (^(v1)), F1: false}
                                                                            var v49 int64 = v23
                                                                            _ = v48
                                                                            _ = v49
                                                                        } else {
                                                                            var v50 int64 = -136
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                for fi51 := 0; fi51 < 15; fi51++ {
                                                                                    var v52 int = (fi51 - -87)
                                                                                    _ = fi51
                                                                                    _ = v52
                                                                                }
                                                                            }
                                                                            v53 := []int{(fi44 % ((915) | 1)), (fi44 - (fi44 * (fi44 + (fi44 * fi44)))), (fi44 + fi44)}
                                                                            _ = v50
                                                                            _ = v53
                                                                        }
                                                                        v3 = (v13 >= (v13 << (((v13 & (^(v13)))) & 63)))
                                                                    }
                                                                    fmt.Println(fi44)
                                                                }
                                                            }
                                                            fmt.Println(v1)
                                                            v54 := []int{}
                                                            _ = fi44
                                                            _ = v54
                                                        }
                                                    default:
                                                        if false {
                                                            var v55 uint = v13
                                                            v56 := GenS6{F0: (v1 * v1), F1: v3}
                                                            v57 := []int{42, -985}
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        } else {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v58 := GenS2{F0: "foo bar", F1: -464, F2: 981}
                                                                _ = v58
                                                            }
                                                            fmt.Println(v23, v2)
                                                        }
                                                    }
                                                default:
                                                    v3 = true
                                                }
                                            }
                                        }
                                    }
                                default:
                                    v2 = 1
                                }
                            }
                        }
                    }
                case "\n\t":
                    fmt.Println(v3)
                case "a":
                    var v59 uint64 = 18446744073709551615
                    _ = v59
                }
                v60 := []int{127, 9223372036854775807, -401, 127}
                _ = v60
            }
        }
    case 426:
        if ((v1 ^ 42) >= v1) {
            var v61 bool = v3
            v62 := GenS8{F0: 94}
            _ = v61
            _ = v62
        }
    default:
        var v63 int8 = 0
        _ = v63
    }
    v64 := []int{-362, 283}
    v1 = v1
    for fi65 := 0; fi65 < 10; fi65++ {
        v3 = true
        v66 := GenS9{F0: 2147483647, F1: (v1 << (((v13 + 975)) & 63))}
        if len(v64) >= 2 {
            v64 = v64[1:len(v64)]
        }
        _ = fi65
        _ = v66
    }
    switch 127 {
    case 9223372036854775807:
        var v67 string = ("\n\t" + "hello")
        _ = v67
    default:
        if len(v64) >= 2 {
            v64 = v64[1:len(v64)]
        }
    }
    v23 = (v23 & (v23 + (v23 / ((-632) | 1))))
    v68 := GenS10{F0: (v1 << (((v13 * (v13 >> ((v13) & 63)))) & 63)), F1: 32767, F2: v1}
    v1 = 110
    fmt.Println(v2)
    if (((("a" + ("\n\t" + (("hello" + "foo bar") + ""))) + ("a" + ("a" + "\n\t"))) + ((("a" + "") + "foo bar") + "hello")) != "\n\t") {
        switch "\n\t" {
        case "":
            if (2147483647 >= -2147483648) {
                fmt.Println(v1, v24)
                v64 = append(v64, 32767)
                v69 := GenS1{F0: v13, F1: v23, F2: 42}
                _ = v69
            }
        default:
            v64 = append(v64, 4294967295)
        }
        v13 = v13
    } else {
        v24 += float32(1.5)
    }
    var v70 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v64
    _ = v68
    _ = v70
}
