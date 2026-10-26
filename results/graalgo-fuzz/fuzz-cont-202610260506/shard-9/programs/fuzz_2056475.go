package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 bool; F2 int64 }
type GenS6 struct { F0 int32; F1 int; F2 int16 }
type GenS7 struct { F0 uint64; F1 int16; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint8; F2 uint32 }

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = 2
    var v4 uint = 65535
    var v5 int = int(v2)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(v5, v3)
        _ = fi6
    }
    if false {
        v7 := []int{}
        v8 := GenS1{F0: 55, F1: (v5 / ((9223372036854775807) | 1)), F2: -72}
        _ = v7
        _ = v8
    }
    v3 += v3
    v3 = v3
    v3 = v3
    v9 := GenS2{F0: -1}
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int8 = 118
            fmt.Println(fi11)
            v13 := GenS1{F0: v12, F1: (-(fi11)), F2: -79}
            var v14 int16 = 1
            switch (v5 | 32768) {
            case 32768:
                var v15 uint16 = (^(v3))
                _ = v15
            case -927:
                if false {
                    switch ("a" + "hello") {
                    case "a":
                        v16 := GenS3{F0: (v1 % ((uint8(v4)) | 1)), F1: 310}
                        _ = v16
                    }
                    var v17 int64 = 9223372036854775807
                    v18 := GenS1{F0: -77, F1: (fi10 / (((fi10 >> ((v4) & 63))) | 1)), F2: -4}
                    fmt.Println(v2, v12)
                    _ = v17
                    _ = v18
                } else {
                    if true {
                        v14 -= v14
                    }
                    v19 := []int{v5, v5, v5}
                    _ = v19
                }
            case -657:
                v20 := []int{(v5 + 32768), (v5 / (((fi10 ^ fi10)) | 1)), fi11}
                _ = v20
            }
            var v21 int32 = 1
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
            _ = v21
        }
        v22 := GenS3{F0: (v1 << ((65535) & 63)), F1: 729}
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := []int{}
            if (941 <= (v5 ^ (v5 | 32768))) {
                if false {
                    var v25 uint16 = 54
                    _ = v25
                } else {
                    switch (v5 & (fi10 & 1)) {
                    case 2147483647:
                        var v26 int32 = 690
                        _ = v26
                    default:
                        v27 := GenS4{F0: "hello"}
                        _ = v27
                    }
                    v4 += uint(v1)
                    fmt.Println(v2, v4, fi23)
                    v4 = 0
                    var v28 uint8 = (v1 + 1)
                    v29 := GenS3{F0: v28, F1: 4294967295}
                    _ = v28
                    _ = v29
                }
                switch fi23 {
                case 128:
                    switch (("" + "\n\t") + "foo bar") {
                    case "":
                        if len(v24) > 0 {
                            v24[0] = -1
                        }
                    case "hello":
                        fmt.Println(v1)
                    }
                case 42:
                    v30 := GenS4{F0: "hello"}
                    _ = v30
                case -2147483648:
                    var v31 float64 = 1.5
                    _ = v31
                }
                v24 = append(v24, v5)
                _ = len(v24)
                _ = len(v24)
            } else {
                switch (fi10 / ((127) | 1)) {
                case 330:
                    fmt.Println(v3)
                case 0:
                    v24 = append(v24, fi23)
                case 128:
                    _ = len(v24)
                }
            }
            v32 := GenS2{F0: 2147483647}
            var v33 int = fi10
            switch (-(fi10)) {
            case 1:
                _ = len(v24)
            default:
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                var v34 bool = true
                _ = v34
            }
            v4 = v4
            v35 := GenS5{F0: 183, F1: (false && false), F2: 0}
            v36 := GenS5{F0: 2147483647, F1: false, F2: 42}
            _ = fi23
            _ = v24
            _ = v32
            _ = v33
            _ = v35
            _ = v36
        }
        _ = fi10
        _ = v22
    }
    if ((((1.5 + float64(0.1)) + 1e10) < 1.5) && true) {
        v5 -= (^(v5))
        v2 = v2
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 int = (fi37 | (^(fi37)))
        if false {
            var v39 uint16 = (v3 % ((v3) | 1))
            v38 = (v5 >> (((v4 / (((v4 >> (((v4 + (v4 & (v4 + (^(v4)))))) & 63))) | 1))) & 63))
            if true {
                fmt.Println(v38)
                switch "foo bar" {
                case "":
                    var v40 float64 = 1e10
                    _ = v40
                case "foo bar":
                    if true {
                        v1 -= (v1 & (v1 | 219))
                        if (false && true) {
                            switch "" {
                            case "a":
                                var v41 uint = v4
                                _ = v41
                            case "":
                                v39 = 1
                            default:
                                var v42 uint64 = v2
                                _ = v42
                            }
                            v38 += 271
                        } else {
                            for fi43 := 0; fi43 < 18; fi43++ {
                                v3 -= (v39 | uint16(v4))
                                if (!true) {
                                    if false {
                                        fmt.Println(v3, v39, v1)
                                        v44 := []int{(v38 >> (((v4 - 0)) & 63)), fi37}
                                        _ = v44
                                    } else {
                                        fmt.Println(v5, v5)
                                        fmt.Println(v3, fi37)
                                        v45 := GenS6{F0: 0, F1: v5, F2: 52}
                                        _ = v45
                                    }
                                    var v46 float32 = 0.0
                                    fmt.Println(v1)
                                    _ = v46
                                }
                                switch v38 {
                                case 128:
                                    v47 := GenS3{F0: (^(v1)), F1: 413}
                                    _ = v47
                                case 4294967295:
                                    v48 := []int{(-(fi37)), int(v4), (fi37 ^ 32767)}
                                    _ = v48
                                default:
                                    if false {
                                        v5 += v38
                                    }
                                }
                                if false {
                                    var v49 float32 = 1e10
                                    v50 := []int{fi43}
                                    v51 := GenS7{F0: v2, F1: 0, F2: -32768}
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                } else {
                                    var v52 int32 = 1
                                    fmt.Println(v52, fi37, fi43)
                                    v53 := []int{(v38 * (v5 ^ (^(fi37)))), fi37, (fi43 - 32767), v38}
                                    _ = v52
                                    _ = v53
                                }
                                var v54 int = 4294967295
                                _ = fi43
                                _ = v54
                            }
                        }
                        fmt.Println(v5, v1, v1)
                    } else {
                        v5 -= (v5 % (((v38 % ((v38) | 1))) | 1))
                        for fi55 := 0; fi55 < 20; fi55++ {
                            fmt.Println(v4)
                            _ = fi55
                        }
                    }
                case "a":
                    if true {
                        var v56 uint16 = (v3 % ((uint16(v2)) | 1))
                        v39 = (^(v39))
                        var v57 int32 = 42
                        _ = v56
                        _ = v57
                    }
                }
                if false {
                    fmt.Println(fi37, v5)
                    v4 -= uint(v1)
                    if false {
                        v58 := []int{}
                        v58 = append(v58, 0)
                        if true {
                            v39 = (v39 << (((^(v4))) & 63))
                        } else {
                            for fi59 := 0; fi59 < 2; fi59++ {
                                var v60 uint32 = 163
                                v5 = (fi59 - 127)
                                switch (fi59 << ((v4) & 63)) {
                                case 42:
                                    var v61 int32 = -219
                                    _ = v61
                                case 2147483647:
                                    v62 := GenS2{F0: -282}
                                    _ = v62
                                }
                                if len(v58) >= 2 {
                                    v58 = v58[1:len(v58)]
                                }
                                if true {
                                    if (true && (1 != -1)) {
                                        var v63 uint32 = 2
                                        if len(v58) > 0 {
                                            v58[0] = v5
                                        }
                                        var v64 int64 = 1
                                        _ = v63
                                        _ = v64
                                    } else {
                                        v65 := GenS4{F0: "\n\t"}
                                        _ = v65
                                    }
                                    if (!false) {
                                        var v66 bool = (v4 > (^(v4)))
                                        v60 = v60
                                        _ = v66
                                    }
                                    v67 := GenS3{F0: v1, F1: 1}
                                    v3 -= v39
                                    v3 += 1
                                    var v68 uint16 = 220
                                    _ = v67
                                    _ = v68
                                }
                                fmt.Println(fi37, v60)
                                _ = fi59
                                _ = v60
                            }
                        }
                        _ = v58
                    } else {
                        var v69 bool = (v39 > (v39 + (^(v39))))
                        switch 42 {
                        case 747:
                            var v70 float32 = 1e10
                            _ = v70
                        case 705:
                            var v71 string = "foo bar"
                            _ = v71
                        case 127:
                            for fi72 := 0; fi72 < 8; fi72++ {
                                var v73 uint8 = v1
                                fmt.Println(v69, v69, v4)
                                _ = fi72
                                _ = v73
                            }
                        default:
                            for fi74 := 0; fi74 < 20; fi74++ {
                                var v75 uint64 = v2
                                v76 := GenS8{F0: v1, F1: v1, F2: 2147483648}
                                v3 -= (v3 / (((v39 * (v39 + v39))) | 1))
                                _ = fi74
                                _ = v75
                                _ = v76
                            }
                        }
                        v2 = 18446744073709551615
                        _ = v69
                    }
                } else {
                    v1 += v1
                    v39 -= (v3 + v3)
                }
            } else {
                var v77 uint16 = v3
                v78 := GenS6{F0: 313, F1: fi37, F2: 0}
                _ = v77
                _ = v78
            }
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    fmt.Println(v2, v1)
    v79 := GenS5{F0: 74, F1: false, F2: -1}
    v80 := []int{(-(v5)), v5, v5, (v5 / (((v5 - (^(v5)))) | 1))}
    v3 = 932
    v5 -= 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v80:", len(v80), v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v80
    _ = v9
}
