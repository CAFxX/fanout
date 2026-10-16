package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint; F2 int }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 int8, p1 int, p2 int) uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = 42
        v8 := GenS1{F0: 0}
        switch "a" {
        case "\n\t":
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 int16 = int16(p1)
                v11 := GenS1{F0: 116}
                _ = fi9
                _ = v10
                _ = v11
            }
        case "":
            v12 := []int{4294967295, (-(fi6)), fi6, (p1 | p2)}
            _ = v12
        case "foo bar":
            var v13 int16 = 411
            _ = v13
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v14 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 0
}

func main() {
    var v1 string = (("hello" + "") + ((("" + "a") + ("a" + "a")) + ""))
    v2 := []string{}
    var v3 uint = 330
    var v4 int16 = 42
    var v5 uint32 = 0
    switch v1 {
    case "foo bar":
        if false {
            v3 -= (v3 - 254)
        }
    case "hello":
        fmt.Println(v1, v1)
    case "\n\t":
        var v15 string = "a"
        _ = v15
    default:
        fmt.Println(v1, v1)
    }
    v2 = append(v2, v1)
    v16 := GenS1{F0: 128}
    var v17 float64 = 0.0
    v18 := GenS1{F0: 0}
    v3 = (v3 ^ (v3 & (^(v3))))
    switch 127 {
    case -188:
        if (v3 != v3) {
            v3 = (v3 ^ (v3 / ((65535) | 1)))
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 int32 = 2147483647
                _ = fi19
                _ = v20
            }
        } else {
            v21 := GenS1{F0: 210}
            switch -86 {
            case -1:
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 uint = (v3 << (((v3 / ((v3) | 1))) & 63))
                    _ = len(v2)
                    _ = fi22
                    _ = v23
                }
            case -2147483648:
                var v24 int32 = -2147483648
                _ = v24
            case 42:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v25 int64 = -554
                _ = v25
            }
            _ = v21
        }
    case 32768:
        for fi26 := 0; fi26 < 19; fi26++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v27 uint32 = v5
            _ = fi26
            _ = v27
        }
    default:
        var v28 int64 = 0
        _ = v28
    }
    if (-632 == -4) {
        if (82 > 218) {
            switch -1 {
            case 848:
                if ((v4 | (-(v4))) >= v4) {
                    if ((v4 % (((-(v4))) | 1)) == (v4 + 128)) {
                        fmt.Println(v1)
                    }
                } else {
                    v29 := GenS2{F0: (v4 / (((v4 + v4)) | 1)), F1: 151, F2: 1}
                    if (42 > 42) {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            v5 += genFn0(0, 392, (fi30 % ((fi30) | 1)))
                            _ = fi30
                        }
                        if false {
                            v17 += v17
                            if (v17 > v17) {
                                if (v5 >= v5) {
                                    v3 -= (^(v3))
                                    v31 := GenS3{F0: -9223372036854775808, F1: 42}
                                    v32 := GenS4{F0: true}
                                    _ = v31
                                    _ = v32
                                }
                                var v33 float32 = 0.1
                                _ = v33
                            }
                            var v34 int = 2147483647
                            _ = v34
                        } else {
                            var v35 float32 = -2.25
                            var v36 uint16 = 379
                            fmt.Println(v5)
                            _ = v35
                            _ = v36
                        }
                        _ = len(v2)
                    }
                    v17 -= 1e10
                    _ = v29
                }
            case 493:
                fmt.Println(v1)
            default:
                switch (v1 + "foo bar") {
                case "":
                    switch v1 {
                    case "\n\t":
                        v37 := GenS1{F0: 127}
                        _ = v37
                    case "a":
                        v4 = v4
                    default:
                        v38 := GenS5{F0: 127, F1: (v4 * v4), F2: v5}
                        _ = v38
                    }
                case "a":
                    switch v1 {
                    case "foo bar":
                        var v39 uint16 = 903
                        _ = v39
                    }
                case "\n\t":
                    var v40 int8 = 127
                    _ = v40
                }
            }
        } else {
            var v41 float64 = ((v17 / v17) - (v17 + float64(1.5)))
            _ = v41
        }
        switch -198 {
        case 868:
            var v42 uint16 = 42
            _ = v42
        default:
            v17 = float64(1e10)
        }
        v3 += 283
    } else {
        for fi43 := 0; fi43 < 17; fi43++ {
            v44 := GenS4{F0: false}
            _ = fi43
            _ = v44
        }
    }
    var v45 int64 = 9223372036854775807
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v46 uint16 = 274
    fmt.Println(v5, v17, v45)
    var v47 uint = (v3 >> ((4294967295) & 63))
    if false {
        v48 := GenS2{F0: 42, F1: v47, F2: -280}
        if true {
            if true {
                if (!true) {
                    if ((v17 + v17) >= v17) {
                        if true {
                            for fi49 := 0; fi49 < 3; fi49++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(fi49, v3, v47)
                                for fi50 := 0; fi50 < 11; fi50++ {
                                    var v51 float32 = (3.14159 / (1.5 / (-2.25 + 0.0)))
                                    _ = fi50
                                    _ = v51
                                }
                                _ = fi49
                            }
                        }
                        v52 := GenS2{F0: v4, F1: (v47 / (((v47 - (v47 << (((v47 | (v47 * 255))) & 63)))) | 1)), F2: -2147483648}
                        _ = v52
                    } else {
                        var v53 bool = false
                        var v54 int32 = -85
                        v2 = append(v2, v1)
                        _ = v53
                        _ = v54
                    }
                    v17 = v17
                } else {
                    switch v1 {
                    case "foo bar":
                        if ("" >= (v1 + "foo bar")) {
                            var v55 uint8 = 42
                            var v56 int8 = 1
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v55
                            _ = v56
                        } else {
                            var v57 uint32 = (v5 * (v5 ^ uint32(v46)))
                            _ = v57
                        }
                    case "hello":
                        switch 605 {
                        case -2147483648:
                            v58 := GenS2{F0: v4, F1: v3, F2: 127}
                            _ = v58
                        }
                    case "\n\t":
                        fmt.Println(v45, v46, v45)
                    default:
                        var v59 uint32 = 65535
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 6; fi60++ {
                        var v61 float64 = v17
                        for fi62 := 0; fi62 < 17; fi62++ {
                            fmt.Println(v45, v17, fi62)
                            if (1.5 <= v61) {
                                var v63 int16 = -748
                                v47 += v3
                                switch -1 {
                                case 2147483647:
                                    fmt.Println(v5)
                                case -1:
                                    if (((^(v46)) > (v46 / ((v46) | 1))) || false) {
                                        if (genFn0(100, (fi60 * (-(fi60))), (fi60 * (-(fi60)))) == v5) {
                                            var v64 uint64 = uint64(v46)
                                            _ = v64
                                        } else {
                                            var v65 uint64 = 1
                                            v66 := GenS1{F0: 128}
                                            v67 := GenS1{F0: 191}
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        }
                                        v68 := GenS5{F0: 255, F1: v4, F2: (v5 + 683)}
                                        _ = v68
                                    }
                                case 42:
                                    var v69 uint8 = 208
                                    _ = v69
                                }
                                _ = v63
                            }
                            _ = fi62
                        }
                        fmt.Println(v5, v46)
                        _ = fi60
                        _ = v61
                    }
                    for fi70 := 0; fi70 < 10; fi70++ {
                        for fi71 := 0; fi71 < 2; fi71++ {
                            v72 := GenS2{F0: 66, F1: v3, F2: -664}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v45 = -367
                            var v73 int16 = (v4 + (^(v4)))
                            if false {
                                v3 = v3
                                var v74 bool = (42 == 2147483647)
                                v2 = append(v2, (v1 + v1))
                                _ = v74
                            }
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                        _ = fi70
                    }
                }
                var v75 uint64 = 396
                _ = len(v2)
                _ = v75
            }
        }
        _ = v48
    } else {
        var v76 float64 = v17
        v77 := GenS5{F0: 249, F1: -1, F2: v5}
        var v78 int = -444
        _ = v76
        _ = v77
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
