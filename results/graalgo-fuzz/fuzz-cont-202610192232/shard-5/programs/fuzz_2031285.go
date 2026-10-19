package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int32; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64; F1 uint; F2 bool }
type GenS8 struct { F0 int32; F1 string }
func genFn0(p0 uint8) uint {
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 -= p0
        if false {
            for fi7 := 0; fi7 < 17; fi7++ {
                if false {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(fi6, fi6)
                    case "":
                        p0 = p0
                    }
                }
                p0 += p0
                v8 := GenS1{F0: 2}
                _ = fi7
                _ = v8
            }
        } else {
            var v9 int8 = 42
            _ = v9
        }
        if (p0 <= p0) {
            for fi10 := 0; fi10 < 9; fi10++ {
                p0 = p0
                v11 := []int{(fi6 - -2147483648), fi6, 128, (fi6 ^ (fi10 << ((380) & 63)))}
                var v12 float32 = 0.0
                _ = fi10
                _ = v11
                _ = v12
            }
            switch (("foo bar" + "\n\t") + ("foo bar" + ("hello" + "\n\t"))) {
            case "":
                switch 32767 {
                case -481:
                    v13 := []int{341, (fi6 ^ (fi6 ^ (-(fi6)))), fi6}
                    _ = v13
                case 231:
                    if false {
                        var v14 int = (fi6 & (fi6 >> ((603) & 63)))
                        var v15 string = "foo bar"
                        _ = v14
                        _ = v15
                    }
                }
            case "a":
                fmt.Println(p0, fi6)
            case "hello":
                v16 := []int{fi6, fi6, (fi6 * fi6), 1}
                _ = v16
            default:
                if false {
                    p0 -= (^(p0))
                    var v17 uint8 = (p0 >> ((123) & 63))
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        p0 += p0
                        fmt.Println(p0, fi18, p0)
                        if false {
                            fmt.Println(p0, fi18, p0)
                        }
                        _ = fi18
                    }
                }
            }
        }
        _ = fi6
    }
    p0 -= p0
    p0 = p0
    if (0 < 0) {
        var v19 uint64 = 42
        switch -119 {
        case 329:
            for fi20 := 0; fi20 < 16; fi20++ {
                switch fi20 {
                case 528:
                    v21 := GenS2{F0: -1, F1: -326, F2: uint32(v19)}
                    _ = v21
                case -1:
                    v22 := []int{(fi20 % ((-1) | 1)), (^(fi20)), (fi20 >> ((478) & 63)), fi20}
                    _ = v22
                case 32768:
                    var v23 int8 = -4
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 16; fi24++ {
                        if false {
                            var v25 uint8 = p0
                            v26 := GenS3{F0: 0}
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS2{F0: 0, F1: 539, F2: 1}
                        _ = fi24
                        _ = v27
                    }
                }
                _ = fi20
            }
        case -554:
            for fi28 := 0; fi28 < 14; fi28++ {
                v29 := GenS4{F0: -375}
                _ = fi28
                _ = v29
            }
        default:
            if false {
                for fi30 := 0; fi30 < 6; fi30++ {
                    var v31 int64 = 0
                    switch fi30 {
                    case -784:
                        v32 := []int{fi30, (fi30 + 127)}
                        _ = v32
                    }
                    switch "foo bar" {
                    case "foo bar":
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v34 := GenS4{F0: (v31 + (v31 ^ 9223372036854775807))}
                            _ = fi33
                            _ = v34
                        }
                    case "hello":
                        for fi35 := 0; fi35 < 5; fi35++ {
                            fmt.Println(fi30)
                            _ = fi35
                        }
                    default:
                        for fi36 := 0; fi36 < 2; fi36++ {
                            v37 := GenS5{F0: 263}
                            _ = fi36
                            _ = v37
                        }
                    }
                    _ = fi30
                    _ = v31
                }
            }
        }
        fmt.Println(v19, v19)
        _ = v19
    }
    _ = p0
    return 18446744073709551615
}
func genFn1() int8 {
    if (!(42 != 32768)) {
        var v38 int32 = 0
        fmt.Println(v38)
        switch "" {
        case "":
            v38 += (v38 - 688)
        case "a":
            v38 = v38
        case "\n\t":
            var v39 float32 = 3.14159
            _ = v39
        default:
            v40 := []int{128}
            _ = v40
        }
        _ = v38
    }
    for fi41 := 0; fi41 < 11; fi41++ {
        for fi42 := 0; fi42 < 17; fi42++ {
            fmt.Println(fi41, fi41)
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                var v43 int = 128
                _ = v43
            case "foo bar":
                var v44 uint32 = 912
                _ = v44
            }
            var v45 int = 4294967295
            var v46 uint = 18446744073709551615
            _ = fi42
            _ = v45
            _ = v46
        }
        var v47 int = fi41
        _ = fi41
        _ = v47
    }
    var v48 int8 = 1
    var v49 int32 = -2147483648
    _ = v48
    _ = v49
    return v48
}

func main() {
    var v1 uint32 = 740
    var v2 uint16 = 65535
    var v3 int64 = 9223372036854775807
    var v4 int8 = int8(v1)
    var v5 string = "a"
    v50 := GenS5{F0: v1}
    v51 := GenS4{F0: v3}
    fmt.Println(v4, v3, v5)
    v52 := []int{}
    if ((-32768 == 0) && true) {
        for fi53 := 0; fi53 < 10; fi53++ {
            for fi54 := 0; fi54 < 7; fi54++ {
                var v55 int32 = 2147483647
                _ = fi54
                _ = v55
            }
            _ = len(v52)
            _ = fi53
        }
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 int16 = 364
            if (3.14159 != 0.1) {
                v52 = append(v52, (fi56 * fi56))
                var v58 uint16 = 0
                if (true && (((v3 & v3) > (v3 * 404)) && (-1 != -507))) {
                    switch 306 {
                    case 42:
                        for fi59 := 0; fi59 < 6; fi59++ {
                            v2 -= (^(v2))
                            var v60 uint16 = v58
                            fmt.Println(v2)
                            _ = fi59
                            _ = v60
                        }
                    default:
                        if true {
                            _ = len(v52)
                        } else {
                            for fi61 := 0; fi61 < 20; fi61++ {
                                switch "" {
                                case "a":
                                    v62 := GenS6{F0: false}
                                    _ = v62
                                case "hello":
                                    v63 := GenS4{F0: (v3 - -9223372036854775808)}
                                    _ = v63
                                }
                                v58 += v2
                                _ = fi61
                            }
                            var v64 uint = 2
                            if (18446744073709551615 == 0) {
                                v65 := GenS7{F0: 42, F1: (v64 + (^(v64))), F2: false}
                                _ = v65
                            } else {
                                for fi66 := 0; fi66 < 12; fi66++ {
                                    var v67 float32 = 1.5
                                    _ = len(v52)
                                    _ = fi66
                                    _ = v67
                                }
                                v57 -= (v57 + (v57 % ((-32768) | 1)))
                            }
                            _ = v64
                        }
                    }
                    for fi68 := 0; fi68 < 16; fi68++ {
                        if len(v52) >= 2 {
                            v52 = v52[1:len(v52)]
                        }
                        _ = fi68
                    }
                    for fi69 := 0; fi69 < 5; fi69++ {
                        var v70 float64 = 1e10
                        _ = len(v52)
                        if len(v52) >= 2 {
                            v52 = v52[1:len(v52)]
                        }
                        fmt.Println(v2, v58, v70)
                        var v71 int64 = 42
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                } else {
                    fmt.Println(v57, v4)
                }
                _ = v58
            } else {
                v72 := GenS8{F0: 226, F1: (v5 + v5)}
                var v73 int8 = v4
                if len(v52) >= 2 {
                    v52 = v52[1:len(v52)]
                }
                _ = v72
                _ = v73
            }
            _ = fi56
            _ = v57
        }
    } else {
        v2 = 1
        var v74 uint64 = 42
        for fi75 := 0; fi75 < 3; fi75++ {
            fmt.Println(v3, v2, v2)
            for fi76 := 0; fi76 < 16; fi76++ {
                var v77 uint16 = (v2 << ((0) & 63))
                if len(v52) > 0 {
                    v52[0] = fi76
                }
                _ = fi76
                _ = v77
            }
            _ = fi75
        }
        _ = v74
    }
    if false {
        _ = len(v52)
    }
    for fi78 := 0; fi78 < 10; fi78++ {
        v5 = (("hello" + "a") + v5)
        _ = fi78
    }
    for fi79 := 0; fi79 < 3; fi79++ {
        var v80 int8 = 1
        _ = len(v52)
        _ = fi79
        _ = v80
    }
    v1 -= v1
    v52 = append(v52, -205)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
