package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 uint8; F1 bool; F2 int }
type GenS4 struct { F0 int64; F1 uint32; F2 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint32, p1 int8, p2 uint64) int32 {
    v5 := GenS1{F0: 42, F1: (p0 >> ((uint(p2)) & 63))}
    if false {
        var v6 int16 = 771
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := []int{}
                var v10 string = "a"
                switch fi7 {
                case 42:
                    var v11 int8 = -1
                    _ = v11
                }
                switch v10 {
                case "a":
                    switch fi8 {
                    case -487:
                        if (0 > 0) {
                            if (p1 <= (p1 + 42)) {
                                if true {
                                    fmt.Println(v10)
                                    var v12 int64 = 978
                                    _ = v12
                                } else {
                                    v10 = (("a" + ("\n\t" + "")) + "")
                                }
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                var v13 int16 = (v6 << ((255) & 63))
                                _ = v13
                            } else {
                                v9 = append(v9, -1)
                                if true {
                                    v14 := GenS2{F0: 92, F1: (true && false), F2: uint(p2)}
                                    p2 = (^(p2))
                                    v15 := GenS2{F0: 83, F1: false, F2: 2}
                                    _ = v14
                                    _ = v15
                                } else {
                                    var v16 uint32 = (p0 | p0)
                                    v16 = 2
                                    v10 = ("" + v10)
                                    _ = v16
                                }
                            }
                            fmt.Println(p2)
                        }
                    case 127:
                        p2 = 0
                    default:
                        var v17 float32 = (1.5 - 0.1)
                        _ = v17
                    }
                    p0 = 1
                case "foo bar":
                    v18 := GenS3{F0: 127, F1: (0.0 == 0.1), F2: -2147483648}
                    _ = v18
                case "\n\t":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                default:
                    if false {
                        if len(v9) > 0 {
                            v9[0] = fi7
                        }
                    }
                }
                var v19 uint16 = 388
                _ = fi8
                _ = v10
                _ = v19
                _ = v9
            }
            for fi20 := 0; fi20 < 11; fi20++ {
                p0 -= 0
                _ = fi20
            }
            _ = fi7
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 bool = ((true && false) || true)
            v6 = v6
            v6 = -912
            _ = fi21
            _ = v22
        }
        _ = v6
    } else {
        v23 := []int{}
        _ = v23
    }
    if (121 < 128) {
        switch 128 {
        case 42:
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 bool = false
                v26 := GenS3{F0: 42, F1: (v25 && (v25 && (!((!(v25 || (42 == uint(p2)))) || false)))), F2: fi24}
                _ = fi24
                _ = v25
                _ = v26
            }
        case 604:
            switch 2147483647 {
            case -9223372036854775808:
                p0 += (^(p0))
            case 128:
                p0 += p0
            case -2147483648:
                var v27 int8 = p1
                _ = v27
            }
        default:
            var v28 float64 = -2.25
            _ = v28
        }
        p1 += p1
    } else {
        if true {
            var v29 string = ("foo bar" + ("\n\t" + ""))
            _ = v29
        }
        v30 := GenS2{F0: 1, F1: ((p1 % ((-128) | 1)) <= p1), F2: 2}
        _ = v30
    }
    if false {
        p0 = p0
        var v31 float32 = 3.14159
        _ = v31
    } else {
        fmt.Println(p0, p0, p1)
        var v32 uint32 = (p0 + p0)
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1() uint32 {
    var v33 float64 = 1e10
    var v34 int64 = 0
    var v35 uint16 = 0
    _ = v33
    _ = v34
    _ = v35
    return 42
}

func main() {
    var v1 uint16 = 0
    var v2 int = 152
    var v3 uint = 0
    v4 := []int{(v2 - (^(v2)))}
    var v36 int32 = -1
    v3 = (v3 | v3)
    if (true && (v2 != v2)) {
        v3 = v3
    }
    v37 := GenS1{F0: 42, F1: 2147483647}
    var v38 int32 = genFn0(0, 0, 2)
    v39 := GenS2{F0: 2, F1: false, F2: v3}
    if (((!false) && true) && true) {
        v4 = append(v4, (v2 << ((v3) & 63)))
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        switch "a" {
        case "\n\t":
            if true {
                fmt.Println(v38, v38, v2)
                v38 = (v36 + (v38 / (((v36 | (^(v36)))) | 1)))
            }
        case "a":
            var v41 int32 = v36
            _ = v41
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        switch "hello" {
        case "hello":
            var v42 int8 = -1
            _ = v42
        }
        if false {
            var v43 float32 = 3.14159
            if (286 < 191) {
                var v44 uint64 = 0
                _ = v44
            } else {
                v1 += 65535
            }
            for fi45 := 0; fi45 < 11; fi45++ {
                if (-519 != -1) {
                    switch 4294967295 {
                    case -509:
                        var v46 float64 = 1.5
                        _ = v46
                    default:
                        var v47 int32 = v36
                        _ = v47
                    }
                    if true {
                        v36 = v36
                        var v48 uint16 = 764
                        switch 32767 {
                        case -803:
                            var v49 uint32 = 42
                            if false {
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v48)
                                default:
                                    v50 := GenS2{F0: 253, F1: (v1 == 255), F2: v3}
                                    _ = v50
                                }
                                var v51 uint64 = 946
                                _ = v51
                            }
                            _ = v49
                        case 4294967295:
                            v2 -= v2
                        default:
                            v48 += 65535
                        }
                        _ = v48
                    } else {
                        if (((v1 | 65535) == (^(v1))) && true) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v52 := GenS2{F0: 175, F1: true, F2: v3}
                            _ = v52
                        } else {
                            fmt.Println(v2)
                        }
                        _ = len(v4)
                        var v53 int16 = 806
                        _ = v53
                    }
                    var v54 float32 = (v43 * v43)
                    if (981 < 42) {
                        var v55 int = v2
                        v3 = v3
                        if false {
                            fmt.Println(v38)
                            switch "\n\t" {
                            case "":
                                v56 := GenS4{F0: 900, F1: uint32(v3), F2: 1}
                                _ = v56
                            case "\n\t":
                                if (((float32(((float64(float32((-2.25 / float64((1e10 + v43))))) / (0.1 * (3.14159 + (1.5 / 1.0)))) + 3.14159)) + -2.25) > 0.0) && false) {
                                    v57 := GenS2{F0: 0, F1: ("\n\t" > "\n\t"), F2: (v3 + v3)}
                                    _ = v57
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = fi40
                                    }
                                    switch "" {
                                    case "a":
                                        var v58 bool = false
                                        _ = v58
                                    default:
                                        v55 += 275
                                    }
                                    var v59 uint32 = 397
                                    _ = v59
                                }
                            case "a":
                                v60 := GenS5{F0: 0}
                                _ = v60
                            }
                        } else {
                            var v61 int64 = -9223372036854775808
                            var v62 float32 = v54
                            v63 := GenS4{F0: (v61 >> ((0) & 63)), F1: genFn1(), F2: v61}
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        _ = v55
                    } else {
                        switch 2147483647 {
                        case 4294967295:
                            _ = len(v4)
                        case -880:
                            if true {
                                if true {
                                    v64 := GenS6{F0: 875, F1: fi40}
                                    if (0 < 23) {
                                        v65 := GenS7{F0: 2147483648}
                                        var v66 uint = (v3 >> ((v3) & 63))
                                        if false {
                                            var v67 bool = (true || ("" != "hello"))
                                            _ = v67
                                        } else {
                                            var v68 float64 = float64(float32(-2.25))
                                            v68 -= (v68 - v68)
                                            v69 := GenS3{F0: 42, F1: (false || false), F2: v2}
                                            _ = v68
                                            _ = v69
                                        }
                                        _ = v65
                                        _ = v66
                                    } else {
                                        if (float32(0.0) != (1e10 * v43)) {
                                            if len(v4) > 0 {
                                                v4[0] = fi45
                                            }
                                            v70 := GenS5{F0: -192}
                                            _ = v70
                                        } else {
                                            var v71 int16 = 1000
                                            _ = v71
                                        }
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v2, v3, fi45)
                                    case "\n\t":
                                        var v72 int32 = v38
                                        _ = v72
                                    case "foo bar":
                                        var v73 uint32 = 676
                                        _ = v73
                                    }
                                    _ = v64
                                }
                                v36 = v36
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 += fi40
                                v74 := GenS8{F0: 0, F1: v1}
                                _ = v74
                            }
                        }
                    }
                    v4 = append(v4, 127)
                    v3 = (v3 >> (((v3 << (((v3 << (((v3 & (^(v3)))) & 63))) & 63))) & 63))
                    _ = v54
                } else {
                    v3 = v3
                }
                v36 -= v36
                _ = fi45
            }
            _ = v43
        } else {
            v4 = append(v4, fi40)
            if ((3.14159 / 0.1) < 3.14159) {
                for fi75 := 0; fi75 < 2; fi75++ {
                    var v76 int = v2
                    v36 = (v36 % ((v38) | 1))
                    _ = fi75
                    _ = v76
                }
                _ = len(v4)
            } else {
                var v77 int = -331
                _ = v77
            }
        }
        _ = fi40
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
