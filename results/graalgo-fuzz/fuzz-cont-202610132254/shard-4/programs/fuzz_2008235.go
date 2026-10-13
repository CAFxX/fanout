package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint; F1 int64; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 string }
func genFn0() int32 {
    var v6 uint8 = 255
    var v7 int16 = 0
    v8 := GenS1{F0: 553}
    v9 := GenS2{F0: 840, F1: 0}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() int64 {
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := []int{}
        v12 := GenS1{F0: 42}
        var v13 float32 = 3.14159
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 float64 = 0.1
    for fi15 := 0; fi15 < 11; fi15++ {
        v14 -= 0.1
        fmt.Println(fi15)
        _ = fi15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v14, fi16)
        var v17 float32 = 0.1
        if false {
            v18 := GenS3{F0: 51, F1: 42, F2: -1}
            _ = v18
        } else {
            v19 := []int{fi16, (fi16 >> ((1) & 63)), (-(fi16)), 32767}
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    v20 := []int{829, -2147483648, 100}
    _ = v14
    _ = v20
    return -346
}

func main() {
    var v1 int32 = -1
    var v2 uint8 = 2
    v3 := []uint32{985, 0}
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    if (!(32767 > -1)) {
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint16 = 218
            if (65535 != 65535) {
                v23 := v5[(fi21 - (-(fi21)))]
                var v24 float32 = 3.14159
                delete(v5, fi21)
                _ = v23
                _ = v24
            } else {
                fmt.Println(v2)
                v1 -= v1
                fmt.Println(v2)
            }
            for fi25 := 0; fi25 < 8; fi25++ {
                v2 = (v2 ^ 128)
                if (51 > -128) {
                    v1 = (-(v1))
                    if false {
                        if ((((-2.25 * ((-2.25 - 1e10) + (1.5 / 0.1))) * 0.0) == (0.1 - 1.5)) && (2 == 467)) {
                            fmt.Println(v22, v22, v1)
                        } else {
                            if (9223372036854775807 <= -1) {
                                v1 += v1
                            }
                            if (false && true) {
                                switch (fi21 << ((1) & 63)) {
                                case 127:
                                    _ = len(v4)
                                case -1:
                                    _ = len(v5)
                                case 4294967295:
                                    var v26 string = "\n\t"
                                    _ = v26
                                }
                                if ("foo bar" < "hello") {
                                    if (float64(0.1) != 1.5) {
                                        v1 -= (v1 * -1)
                                        var v27 int64 = 42
                                        v28 := GenS4{F0: (^(v22)), F1: 4294967295, F2: -128}
                                        _ = v27
                                        _ = v28
                                    } else {
                                        if ((1e10 - 3.14159) != 1e10) {
                                            v29 := GenS4{F0: v22, F1: 4294967295, F2: 1}
                                            var v30 uint32 = 42
                                            v2 = (v2 << ((4294967295) & 63))
                                            _ = v29
                                            _ = v30
                                        }
                                        fmt.Println(fi21)
                                    }
                                    v31 := v4[32767]
                                    var v32 int = (fi21 % ((-1) | 1))
                                    _ = v31
                                    _ = v32
                                }
                                switch fi25 {
                                case 473:
                                    if true {
                                        v1 -= 42
                                    }
                                    v2 -= (v2 + (v2 % (((^(v2))) | 1)))
                                    v2 -= 1
                                    v5[128] = false
                                    fmt.Println(v2, v22, v1)
                                    _ = len(v3)
                                default:
                                    v22 += (v22 & 693)
                                }
                            } else {
                                var v33 int = fi25
                                if true {
                                    if (!(!((v22 ^ (v22 >> ((0) & 63))) < v22))) {
                                        v33 += (v33 >> ((4294967295) & 63))
                                        var v34 int16 = 1000
                                        var v35 int16 = -32768
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v33 = (fi25 * (fi21 | fi25))
                                    }
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                    _ = len(v4)
                                }
                                _ = v33
                            }
                        }
                        v36 := GenS3{F0: 65535, F1: -894, F2: -793}
                        v2 -= 2
                        _ = v36
                    }
                    var v37 int = 32767
                    _ = v37
                } else {
                    switch ("a" + ("\n\t" + ("a" + "foo bar"))) {
                    case "\n\t":
                        v22 += (^(v22))
                    }
                    v2 = (v2 - v2)
                }
                var v38 string = ""
                _ = fi25
                _ = v38
            }
            _ = fi21
            _ = v22
        }
        if false {
            v39 := GenS5{F0: -968}
            _ = v39
        } else {
            v1 = v1
            var v40 bool = false
            _ = v40
        }
        v41 := v5[-126]
        _ = v41
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        v43 := GenS2{F0: 0, F1: v2}
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        fmt.Println(fi44, fi44)
        var v45 int16 = -1
        for fi46 := 0; fi46 < 6; fi46++ {
            if (407 >= 0) {
                v1 -= v1
            } else {
                v47 := GenS1{F0: (-(v1))}
                _ = v47
            }
            _ = fi46
        }
        _ = fi44
        _ = v45
    }
    switch "a" {
    case "a":
        switch 1 {
        case 4294967295:
            v2 -= v2
        }
    }
    if (!true) {
        v48 := GenS3{F0: 0, F1: -1, F2: 1}
        v49 := GenS6{F0: -697, F1: 2}
        v3 = append(v3, 65535)
        _ = v48
        _ = v49
    } else {
        var v50 int = 9223372036854775807
        _ = v50
    }
    v1 -= (v1 * (^(v1)))
    if (212 < 1) {
        for fi51 := 0; fi51 < 5; fi51++ {
            for fi52 := 0; fi52 < 3; fi52++ {
                v53 := v4[fi51]
                if ((((3.14159 - -2.25) + (1e10 / 1.0)) == 3.14159) || (!(32767 == 1))) {
                    if (false && (65535 == 493)) {
                        v53 = true
                        var v54 uint16 = 255
                        _ = v54
                    }
                    var v55 int64 = 9223372036854775807
                    _ = v55
                }
                if ((1.5 >= 0.1) && (!v53)) {
                    if true {
                        var v56 uint64 = 2
                        var v57 uint8 = (v2 ^ v2)
                        _ = v56
                        _ = v57
                    } else {
                        fmt.Println(fi51, fi51)
                        var v58 string = "foo bar"
                        if (!false) {
                            if (((v58 + "foo bar") + "a") != ((v58 + v58) + "a")) {
                                var v59 float32 = (((0.0 * (0.1 / 1.0)) * (0.1 + (((1.5 * ((-2.25 + ((float32(-2.25) * (1.5 - 1e10)) / (0.0 + 1.5))) - 1e10)) - (-2.25 - -2.25)) / 3.14159))) - 1e10)
                                _ = v59
                            }
                        } else {
                            var v60 bool = (42 >= 2147483647)
                            _ = v60
                        }
                        switch v58 {
                        case "a":
                            v61 := GenS7{F0: v58}
                            _ = v61
                        case "hello":
                            var v62 uint16 = 1
                            _ = v62
                        case "foo bar":
                            v1 = -1
                        }
                        _ = v58
                    }
                    if true {
                        if false {
                            if (!false) {
                                fmt.Println(v1)
                                var v63 string = "a"
                                _ = v63
                            }
                            v1 = (v1 % (((v1 << ((296) & 63))) | 1))
                            if false {
                                var v64 int = (fi52 * fi51)
                                delete(v4, fi51)
                                _ = v64
                            } else {
                                var v65 uint64 = 0
                                _ = v65
                            }
                            v66 := GenS6{F0: 42, F1: 0}
                            _ = v66
                        } else {
                            fmt.Println(v1)
                            var v67 uint8 = v2
                            var v68 uint = 1
                            _ = v67
                            _ = v68
                        }
                        v69 := v5[fi52]
                        var v70 string = "a"
                        _ = v69
                        _ = v70
                    }
                    v2 += 241
                }
                _ = fi52
                _ = v53
            }
            _ = fi51
        }
    }
    var v71 int16 = -324
    if ((3.14159 > 1.5) && true) {
        fmt.Println(v71, v1, v71)
        if (873 >= 2) {
            var v72 uint64 = 253
            var v73 uint16 = 32767
            var v74 int16 = (v71 % ((int16(v73)) | 1))
            _ = v72
            _ = v73
            _ = v74
        }
    } else {
        _ = len(v4)
        v2 -= v2
        var v75 int8 = -7
        _ = v75
    }
    v71 += v71
    for fi76 := 0; fi76 < 13; fi76++ {
        var v77 uint16 = 84
        v2 = v2
        if (750 >= 0) {
            v2 = (v2 / (((v2 * (v2 & (v2 % (((v2 | v2)) | 1))))) | 1))
        }
        _ = fi76
        _ = v77
    }
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
