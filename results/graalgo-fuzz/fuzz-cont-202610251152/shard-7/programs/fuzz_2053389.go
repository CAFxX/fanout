package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint32, p1 string, p2 int64) int32 {
    var v6 uint16 = 1
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}
func genFn1(p0 int, p1 int, p2 int32) uint16 {
    switch p0 {
    case 127:
        fmt.Println(p0)
    default:
        p2 = (^(p2))
    }
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 280
}

func main() {
    var v1 int32 = -144
    var v2 uint32 = 800
    v3 := make(map[string]int)
    v4 := []uint32{}
    v5 := make(map[uint32]uint32)
    if (true || true) {
        v8 := GenS1{F0: 255, F1: 42}
        _ = v8
    }
    if false {
        switch 32767 {
        case 9223372036854775807:
            v1 -= v1
        case 127:
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS2{F0: (^(v2)), F1: int8(fi9)}
                _ = fi9
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS2{F0: (v2 >> ((65535) & 63)), F1: -117}
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS2{F0: 65535, F1: int8(fi13)}
                var v15 string = "foo bar"
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 uint64 = 4294967295
            _ = fi11
            _ = v12
            _ = v16
        }
    } else {
        var v17 float64 = 3.14159
        switch 1 {
        case 792:
            for fi18 := 0; fi18 < 17; fi18++ {
                v2 = (v2 | (v2 * 0))
                switch fi18 {
                case 127:
                    v19 := GenS3{F0: 592, F1: -1}
                    _ = v19
                case 32767:
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v21 := GenS2{F0: v2, F1: 47}
                        var v22 float64 = ((v17 - ((1.5 / 0.1) / 1.5)) / 1.5)
                        if ((812 == fi20) && true) {
                            var v23 int8 = 1
                            if (!true) {
                                switch (fi18 % (((-(fi20))) | 1)) {
                                case -1:
                                    v2 -= (v2 / ((2147483648) | 1))
                                case -2147483648:
                                    var v24 string = "\n\t"
                                    v1 = -114
                                    _ = v24
                                case 0:
                                    var v25 int64 = -9223372036854775808
                                    _ = v25
                                default:
                                    v4 = append(v4, 42)
                                    _ = len(v5)
                                    fmt.Println(v22)
                                }
                                v1 -= genFn0(v2, "a", 1)
                                v1 += -39
                            }
                            v1 = (v1 + v1)
                            _ = v23
                        } else {
                            switch fi18 {
                            case -556:
                                var v26 float32 = 0.1
                                _ = v26
                            case -9223372036854775808:
                                v1 = (v1 - (v1 * v1))
                            case 32767:
                                v22 = float64((float32(-2.25) / (0.1 + 3.14159)))
                            default:
                                switch (fi20 / ((fi18) | 1)) {
                                case 32768:
                                    if (!(false && (v1 != (^(v1))))) {
                                        if true {
                                            fmt.Println(v2)
                                        }
                                        var v27 uint16 = 32767
                                        _ = v27
                                    } else {
                                        if true {
                                            v2 = 4294967295
                                        }
                                        if true {
                                            var v28 string = "a"
                                            _ = v28
                                        } else {
                                            var v29 int16 = 32767
                                            _ = v29
                                        }
                                    }
                                case 42:
                                    v17 += v17
                                case -547:
                                    if true {
                                        var v30 uint = 255
                                        var v31 string = ""
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            }
                            _ = len(v5)
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                case 836:
                    for fi32 := 0; fi32 < 11; fi32++ {
                        if ((false && true) || false) {
                            _ = len(v4)
                            v17 -= ((((v17 * -2.25) / 3.14159) * v17) / 1.0)
                            v2 -= 65535
                        } else {
                            v33 := GenS2{F0: (v2 >> ((255) & 63)), F1: 0}
                            v17 = float64(1e10)
                            fmt.Println(fi32, v1, v17)
                            var v34 bool = false
                            var v35 int32 = (v1 + (^(v1)))
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = fi32
                    }
                default:
                    v36 := v5[v2]
                    _ = v36
                }
                var v37 uint8 = 231
                _ = fi18
                _ = v37
            }
        }
        _ = v17
    }
    v1 = v1
    switch -9223372036854775808 {
    case 128:
        for fi38 := 0; fi38 < 13; fi38++ {
            switch ("a" + ("a" + ("\n\t" + "hello"))) {
            case "":
                var v39 int8 = 42
                _ = v39
            }
            v40 := GenS3{F0: 378, F1: -514}
            fmt.Println(v1, fi38)
            _ = fi38
            _ = v40
        }
    case -9223372036854775808:
        v41 := GenS3{F0: 4294967295, F1: 42}
        _ = v41
    case 1:
        switch 0 {
        case 847:
            switch 63 {
            case 32768:
                v1 -= 0
            default:
                for fi42 := 0; fi42 < 18; fi42++ {
                    for fi43 := 0; fi43 < 6; fi43++ {
                        if false {
                            var v44 int = (-(fi43))
                            switch -2147483648 {
                            case 32767:
                                var v45 uint16 = genFn1((fi43 & 485), fi43, (v1 / ((v1) | 1)))
                                _ = v45
                            case -9223372036854775808:
                                switch v44 {
                                case -1:
                                    v46 := GenS1{F0: uint16(v2), F1: 2}
                                    _ = v46
                                case -147:
                                    v44 -= fi43
                                case 4294967295:
                                    switch (v44 | (v44 + (-(fi42)))) {
                                    case 9223372036854775807:
                                        v44 = 2147483647
                                    case -665:
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = 42
                                            }
                                            v2 -= 65535
                                            fmt.Println(v2, v44, v2)
                                            v44 = fi42
                                            switch "hello" {
                                            case "foo bar":
                                                var v47 uint32 = (v2 / (((^(v2))) | 1))
                                                _ = v47
                                            default:
                                                _ = len(v4)
                                                if (-9223372036854775808 < 0) {
                                                    if (true && true) {
                                                        if true {
                                                            v2 += (v2 / (((v2 << ((65535) & 63))) | 1))
                                                            v48 := GenS2{F0: (v2 ^ (^(v2))), F1: 98}
                                                            _ = v48
                                                        } else {
                                                            var v49 int8 = 100
                                                            var v50 float32 = 1.5
                                                            _ = len(v4)
                                                            if (!true) {
                                                                v49 += -125
                                                                v2 -= (v2 % ((v2) | 1))
                                                            }
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    } else {
                                                        v1 = -80
                                                    }
                                                }
                                                v51 := GenS1{F0: 255, F1: 255}
                                                _ = v51
                                            }
                                        } else {
                                            v2 -= v2
                                            var v52 int8 = 53
                                            _ = v52
                                        }
                                    }
                                default:
                                    if true {
                                        var v53 int16 = -307
                                        if ((v2 ^ v2) >= 199) {
                                            v53 += 1
                                            fmt.Println(v44, fi43)
                                        }
                                        if (!true) {
                                            v54 := GenS1{F0: 32, F1: 169}
                                            _ = v54
                                        } else {
                                            v53 = v53
                                        }
                                        v44 += fi43
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v2 = (v2 + 4294967295)
                                        _ = v53
                                    } else {
                                        var v55 string = ("\n\t" + ("a" + "\n\t"))
                                        v55 = v55
                                        v4 = append(v4, 2147483647)
                                        _ = v55
                                    }
                                }
                            case 42:
                                var v56 int64 = -9223372036854775808
                                _ = v56
                            }
                            _ = v44
                        }
                        _ = fi43
                    }
                    _ = fi42
                }
            }
        case 32768:
            var v57 uint = 903
            _ = v57
        }
    }
    var v58 uint = 65535
    switch 277 {
    case 9223372036854775807:
        v59 := GenS2{F0: 42, F1: -13}
        _ = v59
    case -1:
        var v60 int16 = 193
        _ = v60
    }
    var v61 int8 = 0
    for fi62 := 0; fi62 < 12; fi62++ {
        fmt.Println(v2, v2)
        v2 += (v2 - 2147483647)
        for fi63 := 0; fi63 < 6; fi63++ {
            var v64 uint16 = 65535
            _ = fi63
            _ = v64
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v61
}
