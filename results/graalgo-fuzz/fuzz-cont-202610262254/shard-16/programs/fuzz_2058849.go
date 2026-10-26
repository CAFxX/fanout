package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 string; F2 uint }
type GenS5 struct { F0 int64; F1 int64; F2 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint; F1 bool }
func genFn0(p0 int16, p1 int32) int64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 = p0
        v6 := GenS1{F0: 15, F1: 342, F2: 0}
        p0 = 1
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 100, F1: 127, F2: (p0 & (p0 & p0))}
    _ = p0
    _ = p1
    _ = v7
    return -482
}
func genFn1(p0 uint8, p1 int64) uint16 {
    v8 := GenS1{F0: 2, F1: 3, F2: 65535}
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 892
            if false {
                v11 := GenS1{F0: 679, F1: v10, F2: 0}
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        if true {
            fmt.Println(p0, p1)
            var v12 string = ("" + "a")
            if true {
                var v13 int32 = -933
                var v14 uint64 = 1
                var v15 int64 = -9223372036854775808
                _ = v13
                _ = v14
                _ = v15
            } else {
                var v16 int16 = -282
                _ = v16
            }
            _ = v12
        } else {
            p0 = (^(p0))
        }
    } else {
        v17 := GenS2{F0: 127, F1: 127, F2: -1}
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 uint64 = 596
            _ = fi18
            _ = v19
        }
        v20 := []int{9223372036854775807, -2147483648}
        _ = v17
        _ = v20
    }
    v21 := GenS1{F0: 255, F1: 32, F2: 65535}
    p1 = p1
    var v22 int64 = p1
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v8
    return 65535
}

func main() {
    v1 := []string{"\n\t", ("foo bar" + ("hello" + "hello")), "hello"}
    var v2 int8 = -31
    var v3 bool = (1000 == -388)
    v4 := []uint32{247, 1, 2, 2}
    if v3 {
        fmt.Println(v3, v3)
        if v3 {
            for fi23 := 0; fi23 < 8; fi23++ {
                v4 = append(v4, 42)
                v24 := GenS1{F0: 42, F1: 4294967295, F2: 65535}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                switch 0 {
                case 507:
                    v3 = (!true)
                case -766:
                    v2 = v2
                case 0:
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v2, fi25, v3)
                    case "":
                        v2 += v2
                    default:
                        var v26 int64 = -9223372036854775808
                        _ = v26
                    }
                default:
                    if false {
                        if false {
                            v2 = v2
                            v4 = append(v4, 0)
                        } else {
                            v2 = -128
                        }
                        if v3 {
                            if (1 <= (-(v2))) {
                                var v27 int64 = 1
                                _ = v27
                            } else {
                                var v28 uint16 = 821
                                _ = v28
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v3 = true
                        }
                        if v3 {
                            v2 -= (v2 * (v2 - 50))
                            v29 := GenS2{F0: int8(fi25), F1: v2, F2: 134}
                            _ = v29
                        }
                    }
                }
                for fi30 := 0; fi30 < 18; fi30++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (!false) {
                        fmt.Println(fi30, v3)
                        v31 := GenS3{F0: "foo bar"}
                        _ = v31
                    } else {
                        v3 = v3
                        var v32 uint = 18446744073709551615
                        v4 = append(v4, 4294967295)
                        _ = v32
                    }
                    var v33 int32 = 42
                    _ = fi30
                    _ = v33
                }
                _ = fi25
            }
            v4 = append(v4, 4294967295)
        }
    } else {
        fmt.Println(v3)
        v3 = (!(2147483647 <= 744))
        switch (("hello" + "hello") + "a") {
        case "hello":
            for fi34 := 0; fi34 < 14; fi34++ {
                if v3 {
                    switch "" {
                    case "foo bar":
                        v2 = 0
                    case "\n\t":
                        var v35 int64 = 1
                        _ = v35
                    case "":
                        var v36 int16 = 42
                        _ = v36
                    default:
                        if len(v4) > 0 {
                            v4[0] = 248
                        }
                    }
                } else {
                    switch (fi34 + (fi34 ^ fi34)) {
                    case 4294967295:
                        var v37 int16 = -110
                        _ = v37
                    case 0:
                        switch fi34 {
                        case 32768:
                            var v38 int8 = v2
                            _ = v38
                        case 573:
                            var v39 int = fi34
                            _ = v39
                        default:
                            if false {
                                var v40 uint64 = 799
                                var v41 bool = (1.5 < (float64((1e10 - 1.5)) * 1.5))
                                _ = v40
                                _ = v41
                            }
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        v2 = v2
                    case "":
                        fmt.Println(v3, v3)
                    case "a":
                        var v42 uint32 = 65535
                        _ = v42
                    default:
                        v2 = 127
                    }
                }
                _ = fi34
            }
        }
    }
    v3 = v3
    if (42 >= genFn1(211, int64(v2))) {
        var v43 uint8 = 28
        _ = v43
    } else {
        if len(v1) > 0 {
            v1[0] = ""
        }
        v2 = v2
    }
    v3 = true
    v2 = 0
    v44 := GenS4{F0: 52, F1: "\n\t", F2: 4294967295}
    switch 4294967295 {
    case 127:
        if v3 {
            fmt.Println(v3, v2)
            var v45 uint8 = 2
            _ = v45
        }
    case 2147483647:
        var v46 bool = ((-2.25 / 3.14159) >= ((0.0 * 3.14159) - 1.5))
        _ = v46
    case -43:
        for fi47 := 0; fi47 < 14; fi47++ {
            if (18446744073709551615 <= 18446744073709551615) {
                for fi48 := 0; fi48 < 17; fi48++ {
                    v3 = (2 <= 42)
                    if (255 < 127) {
                        var v49 uint64 = 103
                        var v50 int16 = 42
                        _ = v49
                        _ = v50
                    } else {
                        var v51 uint64 = 4294967295
                        v51 += (^(v51))
                        var v52 uint64 = (v51 ^ (^(v51)))
                        v51 = (v52 << ((65535) & 63))
                        v52 = (v51 * 4294967295)
                        _ = len(v4)
                        _ = v51
                        _ = v52
                    }
                    var v53 int16 = 42
                    _ = fi48
                    _ = v53
                }
                for fi54 := 0; fi54 < 19; fi54++ {
                    switch 4294967295 {
                    case 42:
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            var v55 uint = 324
                            _ = v55
                        default:
                            var v56 uint8 = 39
                            _ = v56
                        }
                    case 1:
                        var v57 int8 = v2
                        _ = v57
                    case -22:
                        var v58 float64 = 0.1
                        _ = v58
                    default:
                        if v3 {
                            v59 := GenS2{F0: v2, F1: (v2 & 75), F2: 981}
                            fmt.Println(v3, v2, v3)
                            _ = v59
                        }
                    }
                    v60 := GenS5{F0: -9223372036854775808, F1: 331, F2: 18446744073709551615}
                    _ = fi54
                    _ = v60
                }
                switch (-(fi47)) {
                case 2147483647:
                    switch fi47 {
                    case 1:
                        switch fi47 {
                        case 4294967295:
                            fmt.Println(v2, v2)
                        case 798:
                            var v61 int = (fi47 ^ -780)
                            _ = v61
                        case 127:
                            var v62 uint32 = 4294967295
                            _ = v62
                        default:
                            v2 = (v2 * v2)
                        }
                    case -53:
                        for fi63 := 0; fi63 < 12; fi63++ {
                            fmt.Println(v3)
                            fmt.Println(v2, fi63, fi47)
                            v64 := GenS6{F0: 128}
                            _ = fi63
                            _ = v64
                        }
                    case 32768:
                        v2 += v2
                    }
                }
            }
            var v65 uint64 = 0
            _ = fi47
            _ = v65
        }
    default:
        for fi66 := 0; fi66 < 18; fi66++ {
            v3 = v3
            var v67 bool = ((1 > 0) || v3)
            v2 = -1
            _ = fi66
            _ = v67
        }
    }
    var v68 uint = 33
    v2 += (v2 - 51)
    v68 += (^(v68))
    v69 := GenS4{F0: -9223372036854775808, F1: "\n\t", F2: 65535}
    var v70 uint16 = 514
    v68 -= (^(v68))
    if (v2 > 127) {
        v71 := GenS7{F0: (v68 | v68), F1: v3}
        if (!((v2 - (^(v2))) >= (^(v2)))) {
            _ = len(v1)
        } else {
            if ((v70 - v70) < (v70 << ((2) & 63))) {
                var v72 int32 = -163
                _ = v72
            } else {
                _ = len(v1)
                var v73 bool = (-2147483648 >= int32(v2))
                _ = v73
            }
        }
        _ = v71
    } else {
        _ = len(v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v68
    _ = v69
    _ = v70
}
