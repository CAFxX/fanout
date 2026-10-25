package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint8, p1 int, p2 int16) bool {
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(p2)
        _ = fi7
    }
    var v8 bool = true
    v9 := GenS1{F0: 2147483647, F1: -2147483648}
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return ((42 < 1) || ("" >= "\n\t"))
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 int16 = 104
    var v4 bool = v1
    var v5 uint64 = 1
    var v6 uint32 = 105
    var v11 uint = 614
    v12 := []int{-405}
    if true {
        fmt.Println(v11, v3, v4)
        switch "\n\t" {
        case "":
            var v13 string = ("foo bar" + "hello")
            _ = v13
        case "foo bar":
            v4 = v1
        case "hello":
            for fi14 := 0; fi14 < 1; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    fmt.Println(fi15)
                    var v16 string = "foo bar"
                    var v17 string = (v16 + "foo bar")
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                fmt.Println(v11, v4)
                var v18 uint16 = 295
                _ = fi14
                _ = v18
            }
        default:
            var v19 bool = (0.1 > -2.25)
            _ = v19
        }
        if true {
            var v20 int8 = 37
            if (true && ((1 >= 170) && (!(2 < 144)))) {
                var v21 uint8 = 1
                var v22 float32 = (0.0 * ((0.1 + 0.1) / (-2.25 / -2.25)))
                for fi23 := 0; fi23 < 14; fi23++ {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v2[936] = fi23
                        _ = fi24
                    }
                    var v25 float64 = 1.5
                    _ = fi23
                    _ = v25
                }
                _ = v21
                _ = v22
            }
            var v26 uint16 = 577
            _ = v20
            _ = v26
        } else {
            v27 := GenS1{F0: int32(v11), F1: -284}
            _ = v27
        }
    }
    switch "a" {
    case "":
        v1 = false
    case "a":
        v6 -= (v6 << (((v11 + (v11 & v11))) & 63))
    }
    v28 := GenS2{F0: 128, F1: 255}
    var v29 float32 = float32((0.0 - (0.1 + (0.1 - 0.1))))
    v30 := GenS3{F0: 42, F1: 42}
    if v1 {
        var v31 uint8 = 42
        for fi32 := 0; fi32 < 4; fi32++ {
            if (1.5 > (1e10 + float64((v29 + v29)))) {
                v11 += v11
            }
            var v33 float32 = v29
            var v34 int64 = int64(v6)
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    } else {
        v29 = v29
        v3 = (v3 - (-(v3)))
        v35 := GenS4{F0: 631, F1: (v5 - (v5 * (v5 / ((v5) | 1)))), F2: (v6 << ((uint(v5)) & 63))}
        _ = v35
    }
    if (42 < 0) {
        var v36 uint64 = v5
        if false {
            if true {
                v37 := GenS3{F0: 2, F1: 2147483647}
                _ = v37
            }
            v3 = (v3 % ((v3) | 1))
        } else {
            if false {
                if v4 {
                    var v38 int8 = -97
                    for fi39 := 0; fi39 < 20; fi39++ {
                        _ = len(v12)
                        _ = fi39
                    }
                    _ = v38
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
            v40 := v2[v6]
            var v41 int8 = 37
            _ = v40
            _ = v41
        }
        _ = v36
    }
    if true {
        for fi42 := 0; fi42 < 5; fi42++ {
            var v43 string = ""
            v44 := GenS3{F0: 42, F1: 787}
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = len(v12)
        if len(v12) > 0 {
            v12[0] = 1
        }
    }
    if v4 {
        for fi45 := 0; fi45 < 1; fi45++ {
            if (!v1) {
                v1 = ((!v4) && ((65535 == (^(v11))) || true))
                var v46 bool = v1
                _ = v46
            } else {
                v12 = append(v12, 128)
                _ = len(v12)
            }
            v11 -= 418
            if (360 <= -387) {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                for fi47 := 0; fi47 < 7; fi47++ {
                    v11 = (v11 / ((269) | 1))
                    v5 = v5
                    _ = fi47
                }
                fmt.Println(v3, v4)
            }
            _ = fi45
        }
        v3 = v3
        if v1 {
            v5 += v5
        }
    } else {
        var v48 string = "\n\t"
        if v1 {
            delete(v2, v6)
        } else {
            v49 := GenS4{F0: 2, F1: (v5 % (((^(v5))) | 1)), F2: (v6 % ((v6) | 1))}
            v50 := GenS3{F0: 192, F1: -447}
            if v1 {
                for fi51 := 0; fi51 < 6; fi51++ {
                    if (v3 > v3) {
                        v5 = (v5 << ((v11) & 63))
                        var v52 uint32 = (^(v6))
                        var v53 uint8 = uint8(v11)
                        _ = v52
                        _ = v53
                    } else {
                        var v54 uint64 = (^(v5))
                        if (true && v1) {
                            v55 := GenS5{F0: 0, F1: (^(v6))}
                            for fi56 := 0; fi56 < 10; fi56++ {
                                var v57 uint = 18446744073709551615
                                if (-1 == 127) {
                                    v6 = v6
                                } else {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                                _ = fi56
                                _ = v57
                            }
                            v4 = v4
                            _ = v55
                        } else {
                            var v58 bool = (!(v29 > float32(float64(v29))))
                            switch v48 {
                            case "hello":
                                if v58 {
                                    var v59 uint64 = (v54 - (v5 & v54))
                                    switch (fi51 & fi51) {
                                    case 9223372036854775807:
                                        for fi60 := 0; fi60 < 6; fi60++ {
                                            if v4 {
                                                v58 = (v1 && v1)
                                                if v4 {
                                                    _ = len(v12)
                                                } else {
                                                    _ = len(v12)
                                                    var v61 uint64 = 18446744073709551615
                                                    _ = v61
                                                }
                                                _ = len(v2)
                                            } else {
                                                fmt.Println(v54, v59)
                                                v12 = append(v12, fi60)
                                            }
                                            fmt.Println(fi51, v59, v4)
                                            var v62 int = (fi51 ^ (fi51 ^ fi51))
                                            _ = fi60
                                            _ = v62
                                        }
                                    default:
                                        if len(v12) > 0 {
                                            v12[0] = 127
                                        }
                                    }
                                    if v1 {
                                        var v63 float32 = 1e10
                                        if false {
                                            for fi64 := 0; fi64 < 2; fi64++ {
                                                v3 += 1
                                                switch fi51 {
                                                case 32768:
                                                    _ = len(v2)
                                                }
                                                var v65 float64 = -2.25
                                                _ = fi64
                                                _ = v65
                                            }
                                        }
                                        _ = v63
                                    }
                                    _ = v59
                                }
                            default:
                                var v66 int32 = 42
                                _ = v66
                            }
                            var v67 uint64 = (v54 % ((18446744073709551615) | 1))
                            _ = v58
                            _ = v67
                        }
                        _ = v54
                    }
                    for fi68 := 0; fi68 < 12; fi68++ {
                        var v69 int32 = -357
                        _ = fi68
                        _ = v69
                    }
                    var v70 uint16 = 0
                    _ = fi51
                    _ = v70
                }
            }
            _ = v49
            _ = v50
        }
        _ = v48
    }
    for fi71 := 0; fi71 < 14; fi71++ {
        for fi72 := 0; fi72 < 3; fi72++ {
            var v73 float64 = 1.5
            _ = fi72
            _ = v73
        }
        v29 = (0.1 + (v29 / float32(3.14159)))
        _ = fi71
    }
    v5 += (v5 << ((364) & 63))
    v4 = (false || v1)
    var v74 float64 = (1e10 / -2.25)
    v5 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
