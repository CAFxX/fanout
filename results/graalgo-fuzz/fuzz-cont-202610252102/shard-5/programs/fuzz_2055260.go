package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 uint32, p1 int) int32 {
    if true {
        v6 := GenS1{F0: 65535, F1: 128, F2: 224}
        switch p1 {
        case -2147483648:
            p1 = p1
        case 4294967295:
            p1 += -82
        }
        p0 += (^(p0))
        _ = v6
    } else {
        v7 := []int{}
        var v8 int16 = 1
        var v9 uint8 = 42
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 1
    _ = p0
    _ = p1
    _ = v10
    return 42
}
func genFn1(p0 bool, p1 string, p2 uint32) uint32 {
    var v11 int64 = -9223372036854775808
    v12 := GenS2{F0: p0, F1: 248}
    v11 = (v11 << ((0) & 63))
    p2 += (^(p2))
    var v13 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    var v4 string = "hello"
    v5 := make(map[int]int)
    v14 := []int{}
    for fi15 := 0; fi15 < 7; fi15++ {
        v1 -= v1
        v16 := v5[fi15]
        _ = fi15
        _ = v16
    }
    var v17 uint8 = 255
    v3 -= v3
    _ = len(v2)
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int8 = v1
        if false {
            if false {
                if true {
                    v20 := GenS2{F0: true, F1: 0}
                    _ = v20
                }
            } else {
                var v21 uint8 = 128
                if (!false) {
                    switch -42 {
                    case 840:
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v23 := GenS1{F0: 65535, F1: fi22, F2: (v3 % (((v3 / ((v3) | 1))) | 1))}
                            fmt.Println(v1, v3)
                            _ = fi22
                            _ = v23
                        }
                    case 127:
                        var v24 int64 = int64(v21)
                        _ = v24
                    }
                    v17 = v17
                    fmt.Println(v21, v1, v19)
                } else {
                    var v25 string = v4
                    v26 := GenS2{F0: true, F1: (v21 / (((v17 / (((^(v17))) | 1))) | 1))}
                    v3 -= (v3 - v3)
                    _ = v25
                    _ = v26
                }
                _ = v21
            }
            if len(v14) > 0 {
                v14[0] = (fi18 ^ fi18)
            }
            fmt.Println(v1)
        } else {
            switch 42 {
            case 42:
                if (!(!false)) {
                    for fi27 := 0; fi27 < 8; fi27++ {
                        v3 = 0
                        if false {
                            v19 -= (v19 / (((v1 % ((v19) | 1))) | 1))
                            v4 = ""
                        } else {
                            v28 := GenS3{F0: v4}
                            _ = v28
                        }
                        v19 = 77
                        _ = fi27
                    }
                    _ = len(v2)
                    v29 := GenS4{F0: (fi18 << ((441) & 63)), F1: 2147483648}
                    _ = v29
                } else {
                    delete(v2, (("" + "") + v4))
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 bool = (false || (!false))
                        _ = fi30
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 14; fi32++ {
                        switch (v4 + (v4 + (v4 + ("a" + "foo bar")))) {
                        case "\n\t":
                            v2[(("a" + ("a" + (("" + ("hello" + "")) + "hello"))) + (v4 + ("a" + "a")))] = true
                            v5[fi32] = fi18
                        case "":
                            v33 := v5[-405]
                            _ = v33
                        }
                        _ = fi32
                    }
                }
            case -589:
                var v34 uint8 = 22
                _ = v34
            case 429:
                v3 = 2
            default:
                if true {
                    v4 = v4
                    var v35 uint = 18446744073709551615
                    _ = v35
                } else {
                    v36 := GenS2{F0: false, F1: (^(v17))}
                    v1 += -58
                    _ = v36
                }
            }
            var v37 int32 = 0
            v17 = (v17 - v17)
            _ = v37
        }
        v17 = 0
        _ = fi18
        _ = v19
    }
    fmt.Println(v1, v1)
    v1 = (v1 ^ (v1 + (v1 >> ((255) & 63))))
    fmt.Println(v17, v3)
    v3 += v3
    if (!false) {
        var v38 uint32 = 638
        if (((!true) || true) && false) {
            var v39 uint = 268
            _ = v39
        } else {
            switch ("hello" + v4) {
            case "foo bar":
                for fi40 := 0; fi40 < 8; fi40++ {
                    v41 := GenS5{F0: -9223372036854775808}
                    _ = fi40
                    _ = v41
                }
            case "\n\t":
                for fi42 := 0; fi42 < 20; fi42++ {
                    var v43 float64 = float64(1.5)
                    v44 := GenS3{F0: (("hello" + v4) + v4)}
                    if ((0.0 * 1e10) >= (1.5 / 1e10)) {
                        v14 = append(v14, fi42)
                        if false {
                            switch ((v4 + "") + v4) {
                            case "foo bar":
                                var v45 bool = ((uint32(v3) > v38) && true)
                                _ = v45
                            case "a":
                                var v46 bool = (9223372036854775807 != 9223372036854775807)
                                _ = v46
                            case "":
                                v38 -= v38
                            }
                        } else {
                            switch v4 {
                            case "":
                                v2[("" + (((v4 + "foo bar") + ("foo bar" + v4)) + v4))] = (v17 > (v17 | v17))
                            case "a":
                                v47 := GenS6{F0: 65535, F1: 255}
                                _ = v47
                            }
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                        if true {
                            if len(v14) > 0 {
                                v14[0] = fi42
                            }
                            var v48 float32 = (1e10 - -2.25)
                            v49 := GenS6{F0: 477, F1: 2}
                            _ = v48
                            _ = v49
                        }
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            case "hello":
                var v50 int64 = 1
                _ = v50
            default:
                fmt.Println(v17, v38)
            }
            if false {
                if (4294967295 < 18446744073709551615) {
                    if (false || true) {
                        switch v4 {
                        case "hello":
                            var v51 uint8 = v17
                            _ = v51
                        case "foo bar":
                            v1 += v1
                        default:
                            if false {
                                if true {
                                    switch v4 {
                                    case "a":
                                        fmt.Println(v3, v17)
                                    }
                                    v52 := GenS1{F0: 635, F1: 32768, F2: v3}
                                    if (false || (42 == 1)) {
                                        var v53 uint64 = v3
                                        switch int(v38) {
                                        case 117:
                                            if (2 < 1) {
                                                for fi54 := 0; fi54 < 15; fi54++ {
                                                    _ = len(v14)
                                                    v55 := v2[(v4 + v4)]
                                                    _ = fi54
                                                    _ = v55
                                                }
                                                var v56 int16 = int16(v17)
                                                fmt.Println(v17, v17)
                                                _ = v56
                                            } else {
                                                v53 -= v3
                                                v1 -= v1
                                                v57 := GenS7{F0: ((int8(v53) < (v1 / (((v1 | (v1 << ((18446744073709551615) & 63)))) | 1))) && (4294967295 == (v3 + v53)))}
                                                _ = v57
                                            }
                                        case 1:
                                            if (42 >= 210) {
                                                fmt.Println(v17, v53, v3)
                                            }
                                        case 878:
                                            _ = len(v14)
                                        default:
                                            switch -344 {
                                            case -9223372036854775808:
                                                fmt.Println(v38)
                                            }
                                        }
                                        var v58 uint32 = (^(v38))
                                        _ = v53
                                        _ = v58
                                    }
                                    _ = v52
                                }
                                var v59 int32 = 2147483647
                                _ = v59
                            }
                        }
                    } else {
                        v60 := GenS1{F0: 255, F1: 4294967295, F2: v3}
                        var v61 float64 = ((((0.0 / (1e10 / -2.25)) + 1e10) + float64(float32(3.14159))) + 0.1)
                        v62 := GenS6{F0: 57, F1: 255}
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                } else {
                    var v63 float32 = (1.5 - ((0.0 / 1e10) + (1e10 + ((0.0 + 0.0) / float32(3.14159)))))
                    if false {
                        var v64 string = (v4 + (v4 + (v4 + "\n\t")))
                        v65 := GenS4{F0: 1, F1: 357}
                        _ = v64
                        _ = v65
                    } else {
                        v63 = 0.0
                        v66 := GenS2{F0: true, F1: v17}
                        v63 += v63
                        _ = v66
                    }
                    _ = v63
                }
            }
            v67 := GenS7{F0: ((^(v3)) > v3)}
            _ = v67
        }
        switch 9223372036854775807 {
        case 2147483647:
            fmt.Println(v4)
        case 738:
            fmt.Println(v38, v3)
        case 156:
            var v68 uint8 = (v17 ^ v17)
            _ = v68
        }
        _ = v38
    }
    _ = len(v2)
    if (4294967295 > 515) {
        v2[v4] = ((v1 ^ v1) <= v1)
        for fi69 := 0; fi69 < 7; fi69++ {
            var v70 uint16 = 1
            _ = fi69
            _ = v70
        }
        for fi71 := 0; fi71 < 16; fi71++ {
            v72 := GenS8{F0: true, F1: 1, F2: 1}
            _ = fi71
            _ = v72
        }
    } else {
        var v73 bool = true
        v73 = (-2147483648 >= 830)
        _ = v73
    }
    var v74 string = "\n\t"
    for fi75 := 0; fi75 < 6; fi75++ {
        for fi76 := 0; fi76 < 8; fi76++ {
            if (true && (true || (1 >= 2))) {
                var v77 uint16 = 507
                delete(v5, fi76)
                _ = v77
            } else {
                v17 = (v17 | (v17 & (v17 >> ((0) & 63))))
                v4 = "foo bar"
            }
            _ = fi76
        }
        fmt.Println(v3, v1)
        _ = fi75
    }
    var v78 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v74:", v74)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v78
}
