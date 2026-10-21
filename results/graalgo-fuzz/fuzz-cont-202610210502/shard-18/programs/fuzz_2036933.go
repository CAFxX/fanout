package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint; F1 int; F2 uint32 }
func genFn0(p0 int32) int64 {
    v4 := GenS1{F0: false, F1: 9223372036854775807, F2: 18446744073709551615}
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 -= (p0 ^ 49)
        v6 := []int{(fi5 << ((2) & 63))}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 42
}
func genFn1() uint {
    switch 32768 {
    case -9223372036854775808:
        if (-2.25 < float32(0.1)) {
            var v7 int32 = -303
            _ = v7
        } else {
            if ((!true) && (float64(1.5) == 0.0)) {
                if (1 < 595) {
                    v8 := GenS1{F0: true, F1: 127, F2: 568}
                    var v9 uint64 = 18446744073709551615
                    _ = v8
                    _ = v9
                } else {
                    var v10 int = 9223372036854775807
                    v11 := GenS2{F0: true}
                    _ = v10
                    _ = v11
                }
                v12 := []int{0, -2147483648, 42}
                _ = v12
            }
        }
    case 127:
        var v13 uint64 = 511
        _ = v13
    }
    v14 := []int{127, 127, 4294967295}
    var v15 uint8 = 255
    for fi16 := 0; fi16 < 10; fi16++ {
        if (-9223372036854775808 <= 1) {
            if true {
                v17 := GenS2{F0: true}
                var v18 uint64 = 4294967295
                _ = v17
                _ = v18
            }
            fmt.Println(v15, fi16)
        }
        v15 += v15
        _ = fi16
    }
    if true {
        v19 := GenS3{F0: 1, F1: 42}
        _ = v19
    } else {
        switch 712 {
        case 128:
            if (true && (1 >= 248)) {
                for fi20 := 0; fi20 < 14; fi20++ {
                    if true {
                        var v21 int = (fi20 * fi20)
                        if (!(((false || true) || true) || (false && false))) {
                            fmt.Println(v21, v15, v15)
                            v22 := GenS4{F0: 32767, F1: 625}
                            for fi23 := 0; fi23 < 1; fi23++ {
                                if (1 == -1) {
                                    var v24 uint64 = 2
                                    _ = v24
                                }
                                _ = fi23
                            }
                            _ = v22
                        } else {
                            fmt.Println(fi20, v15, fi20)
                            if ((v15 * (v15 - v15)) < v15) {
                                if (1 == 696) {
                                    var v25 uint32 = 65535
                                    fmt.Println(v15, v21, v25)
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    _ = v25
                                }
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    v27 := GenS4{F0: 32767, F1: 4294967295}
                                    fmt.Println(v15, fi20, fi26)
                                    _ = fi26
                                    _ = v27
                                }
                                v21 = (-(v21))
                            } else {
                                fmt.Println(fi20, fi20, fi20)
                                v21 -= (^(v21))
                            }
                            var v28 int64 = 42
                            _ = v28
                        }
                        _ = v21
                    }
                    _ = fi20
                }
                for fi29 := 0; fi29 < 8; fi29++ {
                    v15 -= 0
                    v15 += v15
                    _ = fi29
                }
                var v30 float32 = -2.25
                _ = v30
            }
        default:
            var v31 int16 = 32767
            _ = v31
        }
        v15 = 255
        fmt.Println(v15)
    }
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 bool = false
    var v2 float32 = 3.14159
    var v3 int = 0
    if (!(1000 == -32768)) {
        if (v1 && false) {
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 uint8 = 9
                v1 = false
                _ = fi32
                _ = v33
            }
        } else {
            if false {
                if v1 {
                    if false {
                        v34 := GenS3{F0: (v3 - 0), F1: 4294967295}
                        v35 := []int{v3}
                        _ = v34
                        _ = v35
                    } else {
                        v36 := GenS3{F0: v3, F1: 2}
                        var v37 uint32 = 4294967295
                        _ = v36
                        _ = v37
                    }
                    v1 = true
                    for fi38 := 0; fi38 < 1; fi38++ {
                        v39 := GenS1{F0: (255 == 1), F1: (-(v3)), F2: 255}
                        _ = fi38
                        _ = v39
                    }
                }
                switch 1 {
                case 42:
                    v40 := []int{(v3 | v3), v3, (^(v3)), (v3 & v3)}
                    _ = v40
                }
                var v41 uint = genFn1()
                _ = v41
            } else {
                v42 := []int{32768}
                _ = v42
            }
            v3 += (v3 - -270)
            v43 := []int{v3, -1, -803}
            _ = v43
        }
    }
    v44 := []int{v3}
    var v45 bool = v1
    if v1 {
        if (v1 || true) {
            v1 = (88 > 349)
        } else {
            if v1 {
                if (v45 && (9223372036854775807 == genFn0(238))) {
                    if false {
                        switch ("foo bar" + "") {
                        case "hello":
                            if (true && (v2 >= ((1e10 * float32((0.0 / (-2.25 * -2.25)))) - v2))) {
                                v46 := GenS3{F0: (v3 + (v3 % (((v3 | (v3 + v3))) | 1))), F1: 2}
                                v3 -= v3
                                _ = v46
                            }
                        case "":
                            var v47 uint8 = 2
                            _ = v47
                        case "a":
                            switch 2147483647 {
                            case 42:
                                v48 := GenS2{F0: ((!v45) || v1)}
                                _ = v48
                            case 9223372036854775807:
                                for fi49 := 0; fi49 < 3; fi49++ {
                                    fmt.Println(fi49, v45)
                                    var v50 uint32 = 2147483648
                                    if ((v2 * (v2 + v2)) >= float32(0.0)) {
                                        switch v3 {
                                        case -238:
                                            v51 := GenS1{F0: (1 != 2), F1: (^(v3)), F2: 794}
                                            _ = v51
                                        case 4294967295:
                                            var v52 uint = 1
                                            _ = v52
                                        case -1:
                                            fmt.Println(v3)
                                        }
                                        v45 = (2 >= 4294967295)
                                    }
                                    _ = fi49
                                    _ = v50
                                }
                            case -197:
                                v53 := GenS5{F0: (230 != 0), F1: 532}
                                _ = v53
                            }
                        }
                        fmt.Println(v2, v3)
                    } else {
                        v1 = (v1 && v1)
                    }
                    if (1 < -120) {
                        var v54 int16 = 298
                        v44 = append(v44, (v3 | 127))
                        _ = v54
                    }
                } else {
                    v55 := GenS4{F0: 2, F1: 1}
                    var v56 float64 = -2.25
                    fmt.Println(v2, v2, v45)
                    _ = v55
                    _ = v56
                }
            } else {
                if v1 {
                    if v1 {
                        v45 = (2 != 649)
                        var v57 float64 = (0.0 - 1e10)
                        var v58 int32 = -1
                        _ = v57
                        _ = v58
                    }
                    if len(v44) >= 2 {
                        v44 = v44[1:len(v44)]
                    }
                    v59 := GenS6{F0: 18446744073709551615, F1: 221, F2: 2}
                    _ = v59
                }
                if len(v44) >= 2 {
                    v44 = v44[1:len(v44)]
                }
                for fi60 := 0; fi60 < 8; fi60++ {
                    var v61 int8 = -91
                    _ = fi60
                    _ = v61
                }
            }
            v44 = append(v44, 1)
        }
    } else {
        fmt.Println(v45, v45, v45)
        for fi62 := 0; fi62 < 1; fi62++ {
            v3 += (fi62 << ((65535) & 63))
            var v63 int8 = int8(v3)
            _ = fi62
            _ = v63
        }
        var v64 int = v3
        _ = v64
    }
    fmt.Println(v3)
    v1 = false
    v44 = append(v44, v3)
    v44 = append(v44, v3)
    v45 = (768 != 2)
    var v65 float64 = float64(1.5)
    v3 -= 356
    v45 = v1
    for fi66 := 0; fi66 < 16; fi66++ {
        v67 := GenS4{F0: 2, F1: 308}
        if len(v44) >= 2 {
            v44 = v44[1:len(v44)]
        }
        _ = fi66
        _ = v67
    }
    fmt.Println(v2)
    var v68 uint = 889
    v44 = append(v44, v3)
    switch (v3 << (((v68 >> ((v68) & 63))) & 63)) {
    case 9223372036854775807:
        v68 = (v68 >> (((v68 % ((65535) | 1))) & 63))
    case 548:
        for fi69 := 0; fi69 < 6; fi69++ {
            var v70 float32 = (v2 / v2)
            if (4294967295 == 907) {
                fmt.Println(v2)
                for fi71 := 0; fi71 < 16; fi71++ {
                    switch fi71 {
                    case 479:
                        if (2147483647 != 0) {
                            if true {
                                var v72 uint16 = 32767
                                _ = v72
                            } else {
                                if len(v44) > 0 {
                                    v44[0] = fi69
                                }
                                var v73 float64 = 3.14159
                                _ = v73
                            }
                        }
                    case 128:
                        var v74 int64 = 0
                        _ = v74
                    case 2147483647:
                        _ = len(v44)
                    default:
                        fmt.Println(fi71, fi71)
                    }
                    if ((genFn0(2147483647) != -62) || v1) {
                        if ((v1 && true) && (!false)) {
                            var v75 int8 = 0
                            v76 := GenS6{F0: 788, F1: 1, F2: 0}
                            _ = v75
                            _ = v76
                        }
                        var v77 bool = v45
                        var v78 uint16 = 32767
                        if (v77 || (v45 || false)) {
                            fmt.Println(fi71)
                            var v79 uint32 = 451
                            _ = v79
                        } else {
                            fmt.Println(fi69, v78, v45)
                        }
                        if true {
                            var v80 int8 = 127
                            var v81 int16 = int16(v80)
                            if v45 {
                                var v82 int = -9223372036854775808
                                fmt.Println(v1, v82)
                                _ = v82
                            } else {
                                var v83 uint16 = v78
                                _ = v83
                            }
                            _ = v80
                            _ = v81
                        }
                        _ = v77
                        _ = v78
                    }
                    switch "a" {
                    case "\n\t":
                        v84 := GenS7{F0: (v68 ^ v68), F1: -756, F2: 1}
                        v68 = v68
                        _ = v84
                    }
                    _ = fi71
                }
                if (true && v1) {
                    fmt.Println(v45, v68, v2)
                    if (!false) {
                        v85 := GenS3{F0: fi69, F1: 42}
                        var v86 int16 = 1000
                        _ = v85
                        _ = v86
                    }
                    v68 += genFn1()
                }
            }
            _ = fi69
            _ = v70
        }
    }
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v65
    _ = v68
}
