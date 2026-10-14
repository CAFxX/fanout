package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint16; F2 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint; F1 uint8; F2 int32 }
type GenS10 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7)
        v8 := []int{-269, (fi7 + 4294967295), fi7, (fi7 ^ 601)}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 130
    for fi10 := 0; fi10 < 4; fi10++ {
        if true {
            v9 = (^(v9))
            for fi11 := 0; fi11 < 11; fi11++ {
                if ((0.0 == -2.25) && false) {
                    fmt.Println(fi10, fi10)
                    if true {
                        fmt.Println(fi11)
                        var v12 uint16 = 0
                        if false {
                            fmt.Println(fi11, fi11)
                            v12 -= v12
                            v9 = (v9 | (v9 + v9))
                        }
                        v13 := GenS1{F0: 42, F1: 255}
                        fmt.Println(v9, v9, fi11)
                        var v14 uint16 = 340
                        switch 2147483647 {
                        case 32767:
                            var v15 string = ""
                            _ = v15
                        case -806:
                            var v16 float32 = float32(1e10)
                            _ = v16
                        default:
                            var v17 float64 = float64(1e10)
                            v9 += 182
                            _ = v17
                        }
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v18 bool = false
                    _ = v18
                } else {
                    var v19 int64 = -757
                    if true {
                        v20 := GenS1{F0: 0, F1: v9}
                        v19 = (v19 / (((^(v19))) | 1))
                        if (2 > 331) {
                            v21 := GenS2{F0: 1}
                            if (v9 < v9) {
                                var v22 uint16 = 1
                                fmt.Println(fi11, fi11, v19)
                                v23 := GenS3{F0: 4294967295, F1: (v9 + (v9 | 255))}
                                _ = v22
                                _ = v23
                            } else {
                                var v24 uint64 = 790
                                v9 = (v9 << ((190) & 63))
                                v25 := []int{fi10, 32767}
                                _ = v24
                                _ = v25
                            }
                            _ = v21
                        }
                        _ = v20
                    } else {
                        if (true && (18446744073709551615 == 4294967295)) {
                            if true {
                                var v26 uint16 = 0
                                v27 := GenS1{F0: 336, F1: v9}
                                _ = v26
                                _ = v27
                            }
                            v28 := []int{-9223372036854775808, (fi10 & (fi10 | (-(fi11)))), 2147483647, (-(fi11))}
                            var v29 int32 = -1
                            var v30 int16 = 1000
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        v19 = (v19 * 9223372036854775807)
                        v19 = -9223372036854775808
                    }
                    switch (("" + (("\n\t" + ("foo bar" + "foo bar")) + "hello")) + "foo bar") {
                    case "a":
                        v31 := GenS1{F0: -33, F1: (v9 ^ (v9 % (((v9 >> ((1) & 63))) | 1)))}
                        _ = v31
                    }
                    _ = v19
                }
                _ = fi11
            }
            v32 := GenS4{F0: "foo bar", F1: 2}
            _ = v32
        } else {
            v33 := []int{(fi10 % (((fi10 & 42)) | 1)), (-(fi10)), fi10}
            fmt.Println(v9, fi10)
            v9 -= (v9 * (v9 + v9))
            _ = v33
        }
        v34 := GenS4{F0: ("foo bar" + ""), F1: 18446744073709551615}
        _ = fi10
        _ = v34
    }
    var v35 bool = (815 >= 929)
    _ = v35
    _ = v9
    return 488
}
func genFn1() uint {
    var v36 bool = (0 <= 1)
    v36 = false
    var v37 uint32 = 65535
    _ = v36
    _ = v37
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 int16 = 560
    v4 := make(map[uint32]int)
    var v5 uint32 = 58
    v6 := []uint32{(^(v5)), (v5 / ((1) | 1))}
    var v38 int16 = v3
    v1[(v5 >> ((0) & 63))] = (!true)
    v39 := GenS2{F0: -2147483648}
    switch "hello" {
    case "":
        v40 := GenS1{F0: 257, F1: 128}
        _ = v40
    case "foo bar":
        v6 = append(v6, v5)
    }
    switch 4294967295 {
    case 9223372036854775807:
        var v41 bool = true
        _ = v41
    case 42:
        if (!(-647 <= -223)) {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    v38 = v38
                default:
                    v42 := GenS4{F0: "foo bar", F1: 4294967295}
                    _ = v42
                }
            case "foo bar":
                v43 := GenS4{F0: ("\n\t" + ""), F1: 42}
                _ = v43
            case "":
                v44 := v1[(v5 / (((v5 + (^(v5)))) | 1))]
                _ = v44
            default:
                switch 9223372036854775807 {
                case -9223372036854775808:
                    if true {
                        if (("a" + ("hello" + "")) >= ("hello" + (((("a" + "foo bar") + "") + "") + "foo bar"))) {
                            for fi45 := 0; fi45 < 11; fi45++ {
                                if (((fi45 * 32767) > (fi45 >> ((genFn1()) & 63))) || true) {
                                    var v46 uint32 = 2
                                    for fi47 := 0; fi47 < 12; fi47++ {
                                        switch int(v46) {
                                        case 2147483647:
                                            var v48 int64 = 1
                                            _ = v48
                                        case -2147483648:
                                            v5 += 341
                                        case 888:
                                            switch "" {
                                            case "foo bar":
                                                v46 += 2147483647
                                            }
                                        }
                                        _ = fi47
                                    }
                                    fmt.Println(fi45, fi45, v38)
                                    _ = v46
                                }
                                var v49 int32 = 1
                                var v50 uint32 = v5
                                _ = fi45
                                _ = v49
                                _ = v50
                            }
                            var v51 uint16 = 0
                            if true {
                                delete(v2, "a")
                                fmt.Println(v38, v51)
                            } else {
                                v52 := GenS5{F0: v5, F1: 1, F2: (v38 % (((v3 / (((v3 + v38)) | 1))) | 1))}
                                v53 := GenS6{F0: v3}
                                _ = v52
                                _ = v53
                            }
                            _ = v51
                        } else {
                            var v54 uint = 414
                            v3 += v38
                            _ = v54
                        }
                    }
                }
            }
            var v55 int32 = 455
            _ = v55
        }
    default:
        var v56 bool = false
        _ = v56
    }
    if true {
        if true {
            var v57 float64 = -2.25
            switch 0 {
            case -2147483648:
                var v58 uint32 = v5
                _ = v58
            case 1:
                v59 := GenS2{F0: 361}
                _ = v59
            case -577:
                var v60 float64 = ((0.0 * v57) + v57)
                _ = v60
            }
            _ = v57
        } else {
            var v61 int = 999
            _ = v61
        }
        v38 -= 32767
        for fi62 := 0; fi62 < 18; fi62++ {
            switch "foo bar" {
            case "":
                switch ("\n\t" + "hello") {
                case "":
                    if (65535 > 255) {
                        if false {
                            var v63 float32 = -2.25
                            var v64 uint32 = (^(v5))
                            v5 += v64
                            _ = v63
                            _ = v64
                        } else {
                            v5 -= (v5 | 976)
                            v65 := GenS7{F0: "foo bar", F1: uint16(v5), F2: v5}
                            switch "hello" {
                            case "a":
                                v4[2147483648] = fi62
                            case "hello":
                                switch ("\n\t" + ("" + "hello")) {
                                case "hello":
                                    var v66 int16 = (-(v3))
                                    _ = v66
                                case "\n\t":
                                    v67 := GenS8{F0: "foo bar"}
                                    _ = v67
                                case "foo bar":
                                    v5 += 2147483648
                                default:
                                    switch "\n\t" {
                                    case "":
                                        delete(v4, v5)
                                    }
                                }
                            default:
                                var v68 float32 = float32(-2.25)
                                _ = v68
                            }
                            _ = v65
                        }
                        v69 := GenS2{F0: int32(v5)}
                        _ = v69
                    } else {
                        var v70 uint8 = 255
                        switch 4294967295 {
                        case 2147483647:
                            switch ("hello" + (("hello" + "") + ("" + ("hello" + "a")))) {
                            case "":
                                for fi71 := 0; fi71 < 15; fi71++ {
                                    fmt.Println(v3, fi62, v70)
                                    var v72 int64 = -9223372036854775808
                                    v73 := GenS9{F0: 116, F1: v70, F2: -252}
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                            case "a":
                                v3 -= (v38 / (((v3 * v3)) | 1))
                            default:
                                v74 := GenS10{F0: 929, F1: 42, F2: uint16(v5)}
                                _ = v74
                            }
                        default:
                            v70 = (v70 | v70)
                        }
                        _ = v70
                    }
                }
            case "hello":
                var v75 int = 127
                _ = v75
            }
            _ = fi62
        }
    }
    var v76 int64 = 9223372036854775807
    v3 += (v38 / (((v38 - 189)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
