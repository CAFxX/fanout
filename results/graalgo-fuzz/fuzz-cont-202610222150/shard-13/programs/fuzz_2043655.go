package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int32; F2 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 uint16; F1 int8 }
func genFn0(p0 int, p1 int32, p2 uint8) int64 {
    var v4 int32 = 805
    for fi5 := 0; fi5 < 11; fi5++ {
        switch (p0 - fi5) {
        case 4294967295:
            v6 := GenS1{F0: fi5, F1: 4294967295}
            _ = v6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        if (42 > 1) {
            v8 := []int{fi7, fi7, p0, (fi7 >> ((255) & 63))}
            _ = v8
        } else {
            fmt.Println(v4, fi7, p2)
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(fi7, v4, p2)
                if (!false) {
                    fmt.Println(p2)
                    var v10 int32 = (p1 ^ (v4 / ((v4) | 1)))
                    _ = v10
                }
                switch "" {
                case "\n\t":
                    var v11 float64 = 3.14159
                    _ = v11
                case "foo bar":
                    var v12 bool = (!false)
                    _ = v12
                case "hello":
                    if true {
                        fmt.Println(fi7, fi9, v4)
                        fmt.Println(p1, fi7)
                        v13 := GenS2{F0: -97, F1: 395, F2: (p1 % ((42) | 1))}
                        _ = v13
                    }
                default:
                    if (!true) {
                        switch 4294967295 {
                        case 32768:
                            if (42 > 0) {
                                v4 = v4
                                v4 = 1
                                v14 := GenS3{F0: -1}
                                v15 := GenS3{F0: 9223372036854775807}
                                _ = v14
                                _ = v15
                            }
                        }
                        if (1 <= 127) {
                            var v16 bool = (((p2 * (p2 & (^(p2)))) < (^(p2))) || true)
                            p2 = 0
                            _ = v16
                        } else {
                            fmt.Println(p0, p2, p2)
                        }
                    } else {
                        var v17 uint = 662
                        _ = v17
                    }
                }
                var v18 uint64 = 955
                var v19 uint32 = 674
                var v20 float32 = (-2.25 / 1.0)
                _ = fi9
                _ = v18
                _ = v19
                _ = v20
            }
            v21 := GenS4{F0: 2147483647}
            _ = v21
        }
        if (!true) {
            switch ("\n\t" + "foo bar") {
            case "a":
                if (42 != 9223372036854775807) {
                    p2 -= 174
                }
            case "\n\t":
                p1 += (v4 % (((p1 + (p1 * v4))) | 1))
            }
            v22 := []int{-583, p0}
            switch "a" {
            case "hello":
                v22 = append(v22, -9223372036854775808)
            case "foo bar":
                switch -1 {
                case -9223372036854775808:
                    fmt.Println(v4, p2, v4)
                case 355:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 int32 = (p1 << ((1) & 63))
                        var v25 uint32 = 4294967295
                        switch (-(fi23)) {
                        case -1:
                            var v26 int64 = 0
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case 9223372036854775807:
                    v27 := GenS1{F0: (p0 >> ((4294967295) & 63)), F1: 42}
                    _ = v27
                }
            case "\n\t":
                p2 += p2
            default:
                var v28 int8 = -90
                _ = v28
            }
            _ = v22
        } else {
            v29 := GenS4{F0: 65535}
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS5{F0: 235, F1: false}
                v32 := GenS5{F0: p2, F1: true}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 0
    var v2 float64 = 0.0
    v3 := make(map[uint32]string)
    v33 := GenS3{F0: -995}
    if false {
        switch 168 {
        case 476:
            for fi34 := 0; fi34 < 15; fi34++ {
                v35 := []int{(fi34 + -643), -9223372036854775808, fi34}
                fmt.Println(v1)
                _ = fi34
                _ = v35
            }
        }
    }
    v36 := GenS6{F0: 0}
    var v37 int16 = 32767
    for fi38 := 0; fi38 < 2; fi38++ {
        v37 -= -1
        var v39 int32 = -1
        _ = fi38
        _ = v39
    }
    v1 = 441
    var v40 int8 = -1
    switch -2147483648 {
    case 9223372036854775807:
        v41 := []int{32768, 32767, 629, 1}
        _ = v41
    default:
        var v42 string = "hello"
        _ = v42
    }
    v43 := v3[(^(v1))]
    v43 = "\n\t"
    v44 := []int{962, 42, 2147483647}
    for fi45 := 0; fi45 < 2; fi45++ {
        var v46 float64 = float64(((float32(v2) + ((1e10 / 1e10) - 1.5)) - 0.0))
        _ = fi45
        _ = v46
    }
    v44 = append(v44, 0)
    for fi47 := 0; fi47 < 13; fi47++ {
        for fi48 := 0; fi48 < 12; fi48++ {
            switch 1 {
            case 367:
                if len(v44) > 0 {
                    v44[0] = (fi47 + (fi47 % (((fi47 ^ fi47)) | 1)))
                }
            case 594:
                v43 = v43
            case 0:
                if true {
                    fmt.Println(v43, v1)
                    var v49 int32 = 1
                    _ = v49
                } else {
                    var v50 uint = 974
                    _ = v50
                }
                if (false && false) {
                    switch v43 {
                    case "foo bar":
                        _ = len(v44)
                    }
                    v51 := GenS7{F0: (42 != 42), F1: -34, F2: fi48}
                    _ = v51
                } else {
                    var v52 float32 = 1e10
                    if (-1 == -1) {
                        if false {
                            if true {
                                if true {
                                    var v53 float64 = 1e10
                                    var v54 float32 = 0.1
                                    var v55 uint32 = v1
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                } else {
                                    switch 9223372036854775807 {
                                    case 1:
                                        v56 := GenS8{F0: (fi47 | (-(fi48)))}
                                        _ = v56
                                    case 42:
                                        var v57 float64 = (v2 / float64(float32(((((v2 * v2) * 1.5) * 1e10) - 0.1))))
                                        _ = v57
                                    default:
                                        var v58 uint = 580
                                        _ = v58
                                    }
                                    v59 := GenS9{F0: fi47}
                                    _ = v59
                                }
                                v40 += (v40 | (v40 << ((255) & 63)))
                                v40 -= v40
                                v37 = v37
                                v40 += v40
                            }
                            var v60 float32 = 0.0
                            _ = v60
                        } else {
                            fmt.Println(v52)
                        }
                        v37 -= -53
                        v40 += v40
                        if (true || (fi48 == fi47)) {
                            if true {
                                var v61 int = fi48
                                _ = v61
                            }
                            if (!((!(true || (42 > 2147483647))) && false)) {
                                var v62 int64 = 454
                                var v63 uint64 = 185
                                var v64 int32 = 42
                                if true {
                                    v37 += (v37 | v37)
                                }
                                v52 -= v52
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                            if true {
                                delete(v3, v1)
                                v65 := GenS10{F0: fi48, F1: 2}
                                v66 := GenS8{F0: (fi48 % ((fi48) | 1))}
                                v1 -= 65535
                                if len(v44) > 0 {
                                    v44[0] = fi47
                                }
                                if false {
                                    fmt.Println(v37, fi47, v37)
                                    v52 = -2.25
                                }
                                switch fi48 {
                                case 30:
                                    v67 := GenS1{F0: 0, F1: 59}
                                    _ = v67
                                case 42:
                                    var v68 uint = 18446744073709551615
                                    var v69 int = fi47
                                    fmt.Println(v1, v52)
                                    _ = v68
                                    _ = v69
                                default:
                                    if true {
                                        switch fi48 {
                                        case 127:
                                            var v70 float32 = v52
                                            _ = v70
                                        }
                                        var v71 uint = 4294967295
                                        switch 42 {
                                        case 4294967295:
                                            v52 = v52
                                            v72 := GenS2{F0: 100, F1: v1, F2: -1}
                                            v73 := GenS2{F0: v40, F1: v1, F2: -2147483648}
                                            _ = v72
                                            _ = v73
                                        default:
                                            fmt.Println(v37, v71, v71)
                                        }
                                        _ = v71
                                    } else {
                                        var v74 int32 = 206
                                        v3[(v1 >> ((1) & 63))] = v43
                                        var v75 uint32 = v1
                                        _ = v74
                                        _ = v75
                                    }
                                }
                                _ = v65
                                _ = v66
                            }
                            var v76 uint16 = 255
                            _ = v76
                        } else {
                            if true {
                                fmt.Println(v1, v1)
                                fmt.Println(fi47)
                            } else {
                                var v77 uint32 = v1
                                v1 = v1
                                _ = v77
                            }
                        }
                    }
                    _ = v52
                }
                var v78 int64 = -9223372036854775808
                _ = v78
            }
            var v79 int = int(v1)
            if (uint64(v1) >= 0) {
                delete(v3, (^(v1)))
                v40 = (^(v40))
                var v80 uint32 = (v1 | (v1 % ((v1) | 1)))
                var v81 float32 = 0.0
                _ = v80
                _ = v81
            }
            _ = fi48
            _ = v79
        }
        v43 = ("\n\t" + "")
        _ = fi47
    }
    var v82 int16 = (v37 ^ v37)
    if (true && true) {
        v83 := v3[v1]
        _ = v83
    }
    if (true && true) {
        v2 += 0.1
    }
    for fi84 := 0; fi84 < 19; fi84++ {
        var v85 float32 = -2.25
        var v86 int64 = 720
        switch "a" {
        case "foo bar":
            v87 := GenS11{F0: 0, F1: v40}
            _ = v87
        default:
            if (v43 <= (v43 + (v43 + ((v43 + (v43 + "\n\t")) + ((((v43 + (((v43 + v43) + ((v43 + ("foo bar" + "")) + v43)) + v43)) + v43) + "\n\t") + "hello"))))) {
                if ((v2 / -2.25) >= (v2 - float64(1.5))) {
                    v1 = 266
                    if (441 >= 4294967295) {
                        var v88 float64 = 1e10
                        _ = v88
                    } else {
                        for fi89 := 0; fi89 < 16; fi89++ {
                            v90 := GenS8{F0: (fi89 % ((-848) | 1))}
                            if (v2 > ((v2 / v2) * float64(-2.25))) {
                                v3[(^(v1))] = v43
                            }
                            var v91 int64 = 42
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                    }
                }
            }
        }
        _ = fi84
        _ = v85
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v40
    _ = v43
    _ = v44
    _ = v82
}
