package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 int64; F2 int }
type GenS6 struct { F0 uint; F1 string }
func genFn0(p0 uint8, p1 uint16, p2 int) bool {
    var v7 int8 = int8(p2)
    var v8 float64 = -2.25
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int) uint32 {
    p0 = 42
    p0 += 32768
    _ = p0
    return 2
}

func main() {
    var v1 int64 = 0
    var v2 string = ""
    var v3 bool = (159 <= 2147483647)
    v4 := make(map[uint32]bool)
    var v5 uint32 = 1
    v6 := make(map[uint32]int)
    var v9 int32 = -92
    v5 = 0
    switch ("hello" + ((("hello" + ((v2 + v2) + "a")) + (v2 + v2)) + "a")) {
    case "a":
        var v10 uint16 = 255
        _ = v10
    case "":
        fmt.Println(v2, v3)
    default:
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 42, F1: v5}
            if v3 {
                switch v2 {
                case "foo bar":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v14 := GenS1{F0: 1000, F1: (v5 / (((v5 / (((v5 & v5)) | 1))) | 1))}
                        var v15 uint64 = 212
                        v16 := []int{fi11, fi11}
                        var v17 int64 = -9223372036854775808
                        v18 := GenS1{F0: 0, F1: v5}
                        v15 -= (v15 * v15)
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                case "hello":
                    if ((v2 >= v2) && (fi11 >= fi11)) {
                        var v19 uint = 1
                        _ = v19
                    }
                case "a":
                    v20 := v6[(v5 * genFn1(fi11))]
                    _ = v20
                default:
                    v3 = ((v1 < v1) || v3)
                }
                for fi21 := 0; fi21 < 12; fi21++ {
                    fmt.Println(fi21)
                    v22 := []int{}
                    _ = fi21
                    _ = v22
                }
            } else {
                switch "a" {
                case "\n\t":
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 float64 = (0.0 - (-2.25 - (((1.5 - float64((-2.25 / 1e10))) * 3.14159) - 1e10)))
                        var v25 float32 = 1e10
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case "":
                    if v3 {
                        var v26 uint32 = v5
                        for fi27 := 0; fi27 < 1; fi27++ {
                            fmt.Println(fi11, v9)
                            fmt.Println(fi27, v5)
                            _ = fi27
                        }
                        v1 += 1
                        _ = v26
                    }
                case "foo bar":
                    switch (fi11 / (((^(fi11))) | 1)) {
                    case 777:
                        v28 := GenS1{F0: 42, F1: genFn1(fi11)}
                        _ = v28
                    case 359:
                        if (v3 && v3) {
                            for fi29 := 0; fi29 < 14; fi29++ {
                                v4[0] = (!false)
                                var v30 uint32 = v5
                                switch ("a" + "\n\t") {
                                case "a":
                                    if (0 <= -32768) {
                                        v31 := GenS1{F0: 32767, F1: 1}
                                        _ = len(v4)
                                        _ = v31
                                    }
                                case "\n\t":
                                    fmt.Println(fi11, fi11, v5)
                                case "":
                                    if (((1.5 - 0.1) - float64(0.1)) <= 3.14159) {
                                        v32 := GenS2{F0: 127, F1: 1, F2: (v9 * v9)}
                                        v33 := []int{fi29, fi11}
                                        v3 = (0 <= 1)
                                        _ = v32
                                        _ = v33
                                    } else {
                                        var v34 int8 = -97
                                        _ = v34
                                    }
                                default:
                                    v35 := GenS2{F0: 128, F1: 691, F2: (v9 ^ v9)}
                                    _ = v35
                                }
                                _ = fi29
                                _ = v30
                            }
                        } else {
                            v5 = 330
                            v2 = ""
                            var v36 uint8 = 148
                            _ = v36
                        }
                    default:
                        if (!((v3 || (1 <= v9)) || v3)) {
                            if v3 {
                                var v37 uint64 = 1
                                delete(v4, (v5 * v5))
                                v1 = v1
                                _ = v37
                            }
                            var v38 uint16 = 610
                            switch fi11 {
                            case 42:
                                var v39 int8 = 76
                                _ = v39
                            case 688:
                                for fi40 := 0; fi40 < 13; fi40++ {
                                    if (!true) {
                                        if false {
                                            switch v2 {
                                            case "":
                                                switch (v2 + "hello") {
                                                case "hello":
                                                    v1 += v1
                                                case "\n\t":
                                                    v2 = v2
                                                }
                                            case "\n\t":
                                                var v41 uint64 = 4294967295
                                                _ = v41
                                            default:
                                                v42 := GenS3{F0: 42, F1: v5}
                                                _ = v42
                                            }
                                            fmt.Println(v2, fi40)
                                        } else {
                                            fmt.Println(v2)
                                        }
                                    } else {
                                        var v43 int8 = 100
                                        _ = v43
                                    }
                                    _ = fi40
                                }
                            case -257:
                                v9 = v9
                            }
                            _ = v38
                        } else {
                            v1 += 611
                        }
                    }
                default:
                    v6[genFn1(fi11)] = fi11
                }
                v2 = v2
                v5 -= (v5 - 2)
            }
            _ = fi11
            _ = v12
        }
    }
    var v44 int32 = v9
    v45 := GenS1{F0: -1, F1: genFn1(709)}
    for fi46 := 0; fi46 < 18; fi46++ {
        switch "a" {
        case "\n\t":
            v5 = 1
        case "":
            v3 = true
        default:
            v2 = v2
        }
        _ = fi46
    }
    if v3 {
        switch ("hello" + v2) {
        case "hello":
            v47 := GenS4{F0: 995}
            _ = v47
        default:
            v5 -= (v5 >> ((18446744073709551615) & 63))
        }
        var v48 int32 = v9
        var v49 uint64 = 1
        _ = v48
        _ = v49
    } else {
        v3 = false
        switch 127 {
        case -2147483648:
            v2 = v2
        default:
            var v50 uint64 = 0
            _ = v50
        }
    }
    if false {
        if v3 {
            var v51 int8 = -45
            _ = v51
        }
    }
    if (128 <= 0) {
        switch (v2 + v2) {
        case "a":
            var v52 string = v2
            _ = v52
        case "hello":
            v53 := []int{-988, 552, 128, -1}
            _ = v53
        }
    } else {
        if v3 {
            var v54 uint32 = (v5 ^ (v5 >> ((1) & 63)))
            _ = v54
        }
        if false {
            var v55 uint16 = 950
            if (128 >= 121) {
                for fi56 := 0; fi56 < 2; fi56++ {
                    v44 -= v44
                    if (0.0 >= ((((float64(0.0) * ((((3.14159 / ((-2.25 - (1e10 + 0.1)) - 1e10)) + float64(((1e10 / 3.14159) / 1.5))) - 1e10) + 1e10)) - ((3.14159 * -2.25) / 1.0)) + (float64((3.14159 / float32(3.14159))) + 0.1)) + 3.14159)) {
                        var v57 uint8 = 127
                        _ = v57
                    }
                    _ = fi56
                }
                v5 += (v5 % (((^(v5))) | 1))
            }
            _ = v55
        }
    }
    var v58 float64 = (1e10 * 3.14159)
    for fi59 := 0; fi59 < 19; fi59++ {
        for fi60 := 0; fi60 < 17; fi60++ {
            v2 = (v2 + v2)
            v9 += (v9 / (((v44 & (^(v44)))) | 1))
            _ = fi60
        }
        _ = fi59
    }
    for fi61 := 0; fi61 < 4; fi61++ {
        v62 := []int{(fi61 * fi61), (^(fi61))}
        if len(v62) > 0 {
            v62[0] = fi61
        }
        v63 := GenS3{F0: 4294967295, F1: 329}
        _ = fi61
        _ = v62
        _ = v63
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        v5 = (v5 | v5)
        fmt.Println(fi64, v9)
        fmt.Println(v3, v58)
        _ = fi64
    }
    if (((710 != -541) && ((v1 >> ((4294967295) & 63)) == (-(v1)))) && v3) {
        switch v2 {
        case "foo bar":
            var v65 int32 = v44
            _ = v65
        case "\n\t":
            var v66 int64 = int64(v5)
            _ = v66
        case "a":
            v2 = ((("foo bar" + v2) + (v2 + v2)) + (v2 + v2))
        }
        v58 -= (v58 * float64(-2.25))
    } else {
        if true {
            v5 += (v5 % ((2) | 1))
            _ = len(v6)
        } else {
            v67 := GenS5{F0: (v9 * (v9 << ((18446744073709551615) & 63))), F1: (-(v1)), F2: 1}
            v3 = (float64(0.1) <= 3.14159)
            var v68 int8 = 34
            _ = v67
            _ = v68
        }
    }
    v1 = 1
    v1 = v1
    for fi69 := 0; fi69 < 9; fi69++ {
        for fi70 := 0; fi70 < 13; fi70++ {
            v44 = (-(v9))
            v71 := GenS6{F0: 42, F1: ""}
            v72 := v6[0]
            var v73 int8 = 42
            var v74 uint8 = 107
            fmt.Println(v5, v44, fi70)
            fmt.Println(v3, v58, v5)
            fmt.Println(v44, fi69, v58)
            _ = fi70
            _ = v71
            _ = v72
            _ = v73
            _ = v74
        }
        _ = fi69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v58
    _ = v6
    _ = v9
}
