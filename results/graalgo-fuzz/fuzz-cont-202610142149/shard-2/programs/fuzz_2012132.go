package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }
type GenS5 struct { F0 int; F1 int32; F2 uint16 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint32; F1 int64; F2 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 string = "a"
    var v4 uint64 = 991
    var v5 int32 = 42
    var v6 string = v3
    _ = len(v1)
    v7 := GenS1{F0: -1, F1: 0, F2: 1}
    v5 = (v5 + 2147483647)
    fmt.Println(v5, v4)
    for fi8 := 0; fi8 < 12; fi8++ {
        if (v5 > -652) {
            fmt.Println(v5, fi8, v3)
            v6 = v6
            var v9 int = fi8
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 uint32 = 323
                    if false {
                        var v12 uint16 = 2
                        var v13 int16 = -1
                        var v14 uint16 = (^(v12))
                        switch (fi8 & fi10) {
                        case -287:
                            if true {
                                v6 = v3
                            } else {
                                if ((v11 >> ((42) & 63)) > (v11 << ((uint(v14)) & 63))) {
                                    fmt.Println(v5)
                                    v15 := GenS2{F0: uint8(v11), F1: 109, F2: 2}
                                    _ = v15
                                } else {
                                    var v16 uint8 = 6
                                    var v17 int8 = -1
                                    var v18 uint = 4294967295
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                            }
                        case 501:
                            switch 32767 {
                            case 0:
                                v19 := GenS1{F0: 42, F1: 31, F2: 0}
                                _ = v19
                            default:
                                delete(v1, v11)
                            }
                        case -1:
                            v20 := GenS3{F0: 1, F1: 255, F2: (v12 | (v12 + 255))}
                            _ = v20
                        }
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                    _ = v11
                }
            } else {
                v21 := GenS4{F0: (v4 % (((v4 ^ v4)) | 1)), F1: 127, F2: 1}
                v22 := []int{}
                var v23 int = fi8
                _ = v21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                if ((int64(v5) < 713) && true) {
                    if (!(32768 < fi8)) {
                        if false {
                            var v25 uint = 235
                            var v26 uint = 42
                            switch -9223372036854775808 {
                            case 42:
                                v27 := GenS2{F0: 2, F1: 127, F2: (v25 + v25)}
                                _ = v27
                            default:
                                if (!(int16(v5) <= -561)) {
                                    v3 = v6
                                } else {
                                    fmt.Println(v4, v26, fi24)
                                }
                                v28 := GenS5{F0: fi8, F1: v5, F2: 2}
                                if true {
                                    v29 := GenS6{F0: 42, F1: fi24}
                                    _ = v29
                                }
                                v5 = (v5 >> ((v26) & 63))
                                _ = v28
                            }
                            _ = v25
                            _ = v26
                        }
                    } else {
                        v5 = v5
                    }
                } else {
                    switch (fi24 ^ (fi8 & (fi24 * (fi24 % ((fi24) | 1))))) {
                    case -9223372036854775808:
                        v4 = (v4 - v4)
                    case -463:
                        v30 := []int{}
                        _ = v30
                    case -928:
                        fmt.Println(v6, v3, v6)
                        if false {
                            switch fi8 {
                            case 9223372036854775807:
                                if (!(!(float64(1.5) >= 1e10))) {
                                    var v31 int8 = -83
                                    var v32 bool = (!true)
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := []int{127}
                                    if len(v33) > 0 {
                                        v33[0] = fi8
                                    }
                                    v6 = ((v6 + v6) + v6)
                                    _ = v33
                                }
                            case 1:
                                v34 := GenS1{F0: 42, F1: -815, F2: -567}
                                _ = v34
                            default:
                                v35 := []int{fi8}
                                _ = v35
                            }
                            v36 := []int{(-(fi8))}
                            if (!(-412 != (-(fi8)))) {
                                if false {
                                    fmt.Println(v3, fi24)
                                    v6 = v3
                                    var v37 bool = true
                                    _ = v37
                                } else {
                                    if len(v36) >= 2 {
                                        v36 = v36[1:len(v36)]
                                    }
                                }
                                v6 = v6
                            }
                            var v38 float64 = (1e10 + ((3.14159 - ((1.5 + (1.5 * 0.1)) / (-2.25 * (0.1 / -2.25)))) - 0.1))
                            v6 = "\n\t"
                            _ = v36
                            _ = v38
                        }
                        var v39 int8 = 28
                        v39 = v39
                        var v40 int16 = -1
                        _ = v39
                        _ = v40
                    default:
                        var v41 uint16 = 65535
                        _ = v41
                    }
                    v5 = v5
                    delete(v2, 0)
                }
                _ = fi24
            }
        }
        v5 = (^(v5))
        var v42 float32 = 0.0
        _ = fi8
        _ = v42
    }
    var v43 float32 = (float32(((0.0 - (float64((((float32(-2.25) * ((1.5 - float32((1e10 / (-2.25 - 0.1)))) / 1.0)) + (1.5 * float32(0.1))) + 1.5)) + 3.14159)) / float64(-2.25))) / float32(3.14159))
    var v44 uint16 = 65535
    v43 += (3.14159 + v43)
    if ((v44 >> ((4294967295) & 63)) == 42) {
        v45 := GenS7{F0: 1, F1: 9223372036854775807, F2: v6}
        v46 := GenS6{F0: 9223372036854775807, F1: 127}
        var v47 uint16 = v44
        _ = v45
        _ = v46
        _ = v47
    } else {
        for fi48 := 0; fi48 < 2; fi48++ {
            v6 = v3
            fmt.Println(v44)
            _ = fi48
        }
    }
    v1[2] = 2147483648
    switch int(v5) {
    case 9223372036854775807:
        for fi49 := 0; fi49 < 14; fi49++ {
            var v50 uint32 = 42
            v51 := GenS7{F0: 557, F1: -1, F2: ""}
            for fi52 := 0; fi52 < 3; fi52++ {
                if (((v43 - v43) / ((float32(0.1) - ((3.14159 + (v43 - (v43 * v43))) / v43)) - (v43 / v43))) <= (1e10 * (v43 * -2.25))) {
                    var v53 int16 = 42
                    fmt.Println(v3)
                    v43 = (float32(-2.25) - 1e10)
                    _ = v53
                }
                var v54 uint16 = (v44 * 255)
                if false {
                    if ((!(554 == uint(v4))) || (!true)) {
                        v43 -= (v43 * v43)
                        var v55 float64 = ((((((float64((((v43 / ((v43 / 1e10) - (v43 * float32(-2.25)))) + float32(0.1)) * v43)) - 1.5) / (0.0 + (1.5 + 3.14159))) + ((-2.25 / 3.14159) * 3.14159)) - 1e10) - -2.25) + ((1.5 + (0.1 - (0.0 * float64(((v43 + v43) / ((0.1 - ((((v43 * v43) + v43) * float32(3.14159)) / v43)) * v43)))))) / 1.0))
                        fmt.Println(v50, v54, v54)
                        var v56 uint = 1
                        v5 += (v5 - -534)
                        _ = v55
                        _ = v56
                    } else {
                        v50 += (v50 % ((v50) | 1))
                        var v57 int8 = 10
                        v2[(fi52 + 32767)] = (fi52 & (-(fi49)))
                        _ = v57
                    }
                }
                _ = fi52
                _ = v54
            }
            _ = fi49
            _ = v50
            _ = v51
        }
    default:
        fmt.Println(v43)
    }
    v58 := v2[159]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v58
    _ = v6
    _ = v7
}
