package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 int16 }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS7 struct { F0 int8; F1 uint8 }
func genFn0() uint8 {
    v6 := GenS1{F0: 42}
    var v7 int16 = -680
    var v8 uint16 = 65535
    v8 += v8
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = 42
    var v2 float32 = float32(0.1)
    v3 := make(map[int]string)
    var v4 uint = 913
    var v5 uint32 = 218
    var v9 string = ("foo bar" + ("\n\t" + ("hello" + ("a" + ("foo bar" + "\n\t")))))
    if (112 <= 42) {
        if ((1.5 - v2) == (v2 / 1.5)) {
            v10 := []int{-1}
            if false {
                v1 = v1
                v1 = (v1 + v1)
                fmt.Println(v2, v9, v9)
            } else {
                v1 -= v1
                v11 := GenS2{F0: -1, F1: 934, F2: 42}
                for fi12 := 0; fi12 < 6; fi12++ {
                    v5 = (v5 ^ v5)
                    _ = fi12
                }
                _ = v11
            }
            v13 := GenS1{F0: 32767}
            _ = v10
            _ = v13
        } else {
            var v14 float64 = ((-2.25 * (0.1 / 3.14159)) / ((3.14159 + 3.14159) / float64(((1e10 * v2) - v2))))
            fmt.Println(v14, v1, v5)
            _ = v14
        }
        if (true && true) {
            v5 += v5
            var v15 uint64 = 1
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v4, fi16, v5)
        for fi17 := 0; fi17 < 14; fi17++ {
            v3[fi16] = ("a" + "a")
            var v18 int16 = 1
            var v19 string = v9
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 int16 = 0
        _ = fi16
        _ = v20
    }
    var v21 uint64 = 1
    v22 := []int{-415, 9223372036854775807, 32768}
    if false {
        var v23 bool = true
        _ = v23
    } else {
        var v24 int16 = -1
        if true {
            for fi25 := 0; fi25 < 15; fi25++ {
                switch ((v9 + (v9 + v9)) + v9) {
                case "":
                    switch v9 {
                    case "foo bar":
                        var v26 uint = (v4 % (((^(v4))) | 1))
                        _ = v26
                    case "hello":
                        if (true && (919 != 2)) {
                            var v27 float32 = v2
                            if false {
                                var v28 int = int(v24)
                                fmt.Println(v21)
                                fmt.Println(v9, v9)
                                _ = v28
                            }
                            _ = v27
                        }
                    case "\n\t":
                        fmt.Println(v9)
                    }
                default:
                    if (!false) {
                        switch "foo bar" {
                        case "\n\t":
                            v2 = v2
                        }
                        for fi29 := 0; fi29 < 5; fi29++ {
                            v30 := GenS2{F0: 1, F1: -1, F2: v24}
                            _ = fi29
                            _ = v30
                        }
                        v31 := GenS3{F0: true, F1: 42, F2: 1}
                        _ = v31
                    } else {
                        v32 := GenS4{F0: 402, F1: 869, F2: 65535}
                        for fi33 := 0; fi33 < 10; fi33++ {
                            switch (fi25 / ((332) | 1)) {
                            case 957:
                                v4 = (v4 % (((^(v4))) | 1))
                            case 225:
                                v21 -= (v21 - v21)
                            default:
                                v4 = (v4 >> ((867) & 63))
                            }
                            v21 += (v21 / (((v21 | (v21 % ((v21) | 1)))) | 1))
                            if ((!(4294967295 >= (v21 & 4294967295))) && false) {
                                var v34 uint32 = (v5 | v5)
                                var v35 int64 = -1
                                _ = len(v3)
                                _ = v34
                                _ = v35
                            }
                            _ = fi33
                        }
                        var v36 uint16 = 32767
                        _ = v32
                        _ = v36
                    }
                }
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 int32 = -1
                    _ = fi37
                    _ = v38
                }
                _ = fi25
            }
        } else {
            v39 := GenS4{F0: 65535, F1: 42, F2: (v4 - uint(v5))}
            var v40 int64 = 273
            for fi41 := 0; fi41 < 7; fi41++ {
                v40 += v40
                if true {
                    var v42 float64 = float64((v2 / v2))
                    _ = v42
                }
                _ = fi41
            }
            _ = v39
            _ = v40
        }
        _ = v24
    }
    if true {
        for fi43 := 0; fi43 < 20; fi43++ {
            fmt.Println(v1, v4, v9)
            v3[fi43] = v9
            v44 := GenS3{F0: (false && true), F1: -1, F2: 118}
            _ = fi43
            _ = v44
        }
    } else {
        fmt.Println(v4, v4)
        if true {
            for fi45 := 0; fi45 < 5; fi45++ {
                var v46 uint8 = 21
                if true {
                    if false {
                        var v47 int64 = 42
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 15; fi48++ {
                        fmt.Println(v21)
                        _ = fi48
                    }
                    v3[int(v46)] = v9
                } else {
                    v22 = append(v22, fi45)
                    v49 := v3[(fi45 - fi45)]
                    _ = v49
                }
                _ = fi45
                _ = v46
            }
            v50 := GenS5{F0: v5, F1: 194}
            _ = v50
        } else {
            v51 := GenS1{F0: 32767}
            _ = v51
        }
        for fi52 := 0; fi52 < 18; fi52++ {
            fmt.Println(v5)
            _ = fi52
        }
    }
    if ((v21 & v21) == 755) {
        var v53 int16 = 32767
        _ = v53
    }
    if (255 == 255) {
        for fi54 := 0; fi54 < 14; fi54++ {
            v5 += v5
            v55 := GenS2{F0: 1, F1: -2147483648, F2: 959}
            _ = fi54
            _ = v55
        }
        v5 -= 2147483648
        var v56 float64 = (float64(float32(-2.25)) - (3.14159 - ((((0.0 * -2.25) - float64(v2)) * float64(1e10)) + -2.25)))
        _ = v56
    } else {
        var v57 int8 = -106
        if false {
            for fi58 := 0; fi58 < 16; fi58++ {
                for fi59 := 0; fi59 < 18; fi59++ {
                    switch (fi59 - fi58) {
                    case 32768:
                        if (9223372036854775807 <= -9223372036854775808) {
                            v9 = v9
                        } else {
                            if (271 != -452) {
                                v60 := GenS4{F0: 235, F1: uint16(v21), F2: (v4 * v4)}
                                var v61 string = v9
                                v62 := GenS6{F0: v5, F1: 2, F2: v21}
                                v3[fi58] = v9
                                _ = v60
                                _ = v61
                                _ = v62
                            } else {
                                var v63 uint = (v4 + (v4 - v4))
                                fmt.Println(fi59, v57, v63)
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                var v64 uint = (v4 | v63)
                                _ = v63
                                _ = v64
                            }
                            v5 += 58
                            v3[(fi58 / ((-2147483648) | 1))] = ""
                        }
                    case 831:
                        v65 := v3[(fi59 & (fi58 | fi58))]
                        _ = v65
                    case 307:
                        var v66 string = v9
                        _ = v66
                    }
                    v57 = v57
                    _ = fi59
                }
                var v67 float64 = 0.0
                switch fi58 {
                case 2147483647:
                    v68 := GenS5{F0: 2147483647, F1: 143}
                    _ = v68
                case 0:
                    for fi69 := 0; fi69 < 11; fi69++ {
                        var v70 uint16 = 65535
                        _ = len(v3)
                        switch (fi58 ^ fi58) {
                        case 4294967295:
                            var v71 int = (fi69 % ((fi69) | 1))
                            _ = v71
                        case 127:
                            v22 = append(v22, (^(fi69)))
                        case 32768:
                            v67 = ((0.0 + -2.25) * (float64(1e10) + (float64(0.1) + v67)))
                        default:
                            var v72 int64 = -239
                            _ = v72
                        }
                        _ = fi69
                        _ = v70
                    }
                case 128:
                    v57 -= (-(v57))
                }
                _ = fi58
                _ = v67
            }
        }
        for fi73 := 0; fi73 < 5; fi73++ {
            if (v4 >= v4) {
                var v74 int8 = (v57 * v57)
                var v75 int8 = v74
                fmt.Println(v9)
                _ = v74
                _ = v75
            } else {
                var v76 int16 = -32768
                var v77 uint = v4
                var v78 bool = (0 >= 1)
                _ = v76
                _ = v77
                _ = v78
            }
            _ = fi73
        }
        _ = v57
    }
    v79 := GenS2{F0: 9223372036854775807, F1: v1, F2: -1}
    if true {
        for fi80 := 0; fi80 < 11; fi80++ {
            v22 = append(v22, (fi80 | (fi80 / ((fi80) | 1))))
            var v81 uint32 = 65535
            v82 := GenS7{F0: 42, F1: 0}
            _ = fi80
            _ = v81
            _ = v82
        }
    }
    if (true && false) {
        fmt.Println(v5)
        v5 += 0
    }
    fmt.Println(v2, v4)
    if false {
        for fi83 := 0; fi83 < 13; fi83++ {
            fmt.Println(v5, v1)
            _ = len(v22)
            _ = fi83
        }
        if true {
            _ = len(v3)
            delete(v3, 32767)
            var v84 uint8 = 2
            _ = v84
        } else {
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        }
        _ = len(v22)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v9
}
