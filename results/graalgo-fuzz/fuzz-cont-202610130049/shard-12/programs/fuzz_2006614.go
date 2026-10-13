package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 string; F1 uint32; F2 int32 }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int = fi7
            v10 := []int{(fi8 / (((fi8 + (-(v9)))) | 1)), v9, (-(v9))}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 uint = 255
    if ((true && (1.5 >= 0.1)) || false) {
        fmt.Println(v11)
        for fi12 := 0; fi12 < 14; fi12++ {
            if false {
                switch "" {
                case "foo bar":
                    if false {
                        switch "foo bar" {
                        case "a":
                            var v13 int8 = -128
                            _ = v13
                        case "":
                            fmt.Println(v11, v11)
                        }
                        fmt.Println(v11, v11)
                        var v14 float64 = 1.5
                        _ = v14
                    } else {
                        v11 = 255
                        for fi15 := 0; fi15 < 15; fi15++ {
                            fmt.Println(fi12)
                            if false {
                                fmt.Println(v11, fi15, fi15)
                                v11 -= v11
                                v11 += 18
                            } else {
                                var v16 int64 = 42
                                v17 := GenS1{F0: int32(v16), F1: (fi15 != (fi12 / ((-2147483648) | 1))), F2: 1}
                                v11 += 415
                                _ = v16
                                _ = v17
                            }
                            fmt.Println(fi12)
                            _ = fi15
                        }
                        fmt.Println(v11)
                    }
                case "a":
                    v18 := []int{128, fi12}
                    _ = v18
                case "\n\t":
                    v19 := GenS1{F0: 82, F1: true, F2: 0}
                    _ = v19
                }
                for fi20 := 0; fi20 < 8; fi20++ {
                    v21 := []int{547, (fi12 & (fi20 - fi20)), (^(fi20)), (fi12 % ((-2147483648) | 1))}
                    if false {
                        var v22 int32 = -1
                        switch 2147483647 {
                        case 1:
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                        }
                        var v23 int64 = 9223372036854775807
                        v24 := GenS1{F0: v22, F1: false, F2: 429}
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        if len(v21) > 0 {
                            v21[0] = (fi12 >> ((0) & 63))
                        }
                        var v25 int64 = 1
                        _ = v25
                    }
                    _ = fi20
                    _ = v21
                }
                var v26 int32 = -483
                _ = v26
            }
            _ = fi12
        }
        var v27 int64 = -9223372036854775808
        _ = v27
    } else {
        var v28 uint64 = 2
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 int = (fi29 & (fi29 % (((fi29 >> (((^(v11))) & 63))) | 1)))
            for fi31 := 0; fi31 < 19; fi31++ {
                fmt.Println(v11, fi31)
                _ = fi31
            }
            for fi32 := 0; fi32 < 14; fi32++ {
                var v33 float64 = 3.14159
                v34 := GenS2{F0: 42, F1: 65535}
                var v35 float32 = 3.14159
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            _ = fi29
            _ = v30
        }
        v28 = v28
        _ = v28
    }
    v36 := []int{-130}
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v37 int = 18
        _ = v37
    case "foo bar":
        if false {
            if len(v36) > 0 {
                v36[0] = 4294967295
            }
        }
    case "":
        var v38 int = 127
        _ = v38
    default:
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
    }
    _ = v11
    _ = v36
    return 2
}

func main() {
    var v1 int32 = -1
    var v2 int8 = 91
    v3 := make(map[int]string)
    var v4 int = -936
    var v5 int8 = (v2 - (v2 | v2))
    var v6 string = "foo bar"
    var v39 uint16 = genFn0()
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 int8 = (v2 + -116)
        _ = fi40
        _ = v41
    }
    switch ("\n\t" + v6) {
    case "hello":
        switch (^(v4)) {
        case 4294967295:
            v42 := GenS2{F0: 2147483647, F1: v39}
            _ = v42
        case 32768:
            var v43 float32 = (3.14159 + (1.5 / 1.0))
            _ = v43
        case 1:
            if (!false) {
                if ((v5 - (v2 / ((85) | 1))) != (v5 % (((v2 >> ((565) & 63))) | 1))) {
                    v44 := GenS3{F0: "foo bar", F1: 846, F2: (v1 >> ((350) & 63))}
                    _ = v44
                }
                for fi45 := 0; fi45 < 1; fi45++ {
                    for fi46 := 0; fi46 < 9; fi46++ {
                        switch (v6 + "foo bar") {
                        case "":
                            fmt.Println(v2, v1, v1)
                        case "hello":
                            var v47 bool = ((^(v39)) != genFn0())
                            _ = v47
                        }
                        delete(v3, (fi46 / ((fi45) | 1)))
                        var v48 uint = 74
                        if ((float32(-2.25) + 0.0) >= (1e10 + 0.0)) {
                            if (uint8(v48) == 24) {
                                switch (^(v4)) {
                                case 2147483647:
                                    v49 := []int{v4, (v4 % ((int(v39)) | 1))}
                                    _ = v49
                                }
                                v50 := GenS4{F0: true, F1: -1}
                                _ = v50
                            } else {
                                delete(v3, v4)
                                var v51 float64 = ((float64(0.0) + ((1.5 + (0.0 + 0.1)) + float64(-2.25))) - float64((0.0 * (3.14159 / ((0.1 - 1.5) + float32((1.5 - 1.5)))))))
                                _ = v51
                            }
                            v52 := GenS3{F0: v6, F1: 497, F2: 0}
                            switch (v6 + "\n\t") {
                            case "foo bar":
                                var v53 uint32 = 65535
                                v3[(v4 % (((fi46 + (fi46 << (((v48 / ((v48) | 1))) & 63)))) | 1))] = v6
                                _ = v53
                            }
                            fmt.Println(v1)
                            _ = v52
                        } else {
                            v6 = v6
                            v54 := GenS4{F0: true, F1: -539}
                            _ = v54
                        }
                        _ = fi46
                        _ = v48
                    }
                    v55 := GenS3{F0: (v6 + v6), F1: 0, F2: 1}
                    _ = fi45
                    _ = v55
                }
                for fi56 := 0; fi56 < 12; fi56++ {
                    if false {
                        v57 := []int{v4, fi56, -381}
                        _ = v57
                    } else {
                        for fi58 := 0; fi58 < 4; fi58++ {
                            fmt.Println(v39)
                            var v59 uint8 = 1
                            _ = fi58
                            _ = v59
                        }
                        v60 := GenS1{F0: (v1 >> ((2) & 63)), F1: false, F2: -32768}
                        _ = v60
                    }
                    v6 = v6
                    _ = fi56
                }
            }
        }
    }
    for fi61 := 0; fi61 < 19; fi61++ {
        var v62 bool = ((v39 - genFn0()) < (v39 >> ((799) & 63)))
        v6 = v6
        _ = fi61
        _ = v62
    }
    switch v4 {
    case 0:
        for fi63 := 0; fi63 < 17; fi63++ {
            v3[4294967295] = (v6 + (v6 + v6))
            fmt.Println(v5, v6)
            _ = fi63
        }
    case 32767:
        var v64 int32 = (v1 & (v1 / (((v1 | (v1 / ((v1) | 1)))) | 1)))
        _ = v64
    case 9223372036854775807:
        fmt.Println(v6, v5, v5)
    default:
        switch (v4 * (v4 >> ((1) & 63))) {
        case 127:
            if true {
                var v65 int64 = -1
                v5 += 1
                _ = v65
            }
        }
    }
    v1 += (v1 % ((1) | 1))
    if (!false) {
        var v66 int16 = 919
        v67 := GenS3{F0: "\n\t", F1: 0, F2: (v1 * v1)}
        v3[(v4 + (v4 | 32767))] = v6
        _ = v66
        _ = v67
    }
    switch (v4 & (v4 ^ (v4 * (v4 >> ((uint(v39)) & 63))))) {
    case 32767:
        switch "\n\t" {
        case "hello":
            var v68 uint64 = 18446744073709551615
            _ = v68
        case "":
            v2 += v5
        default:
            v5 = (v2 | (v2 << ((uint(v39)) & 63)))
        }
    }
    var v69 uint64 = 598
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
