package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 bool; F1 string; F2 int32 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16; F1 int32; F2 uint }
func genFn0() uint16 {
    v7 := []int{-604}
    var v8 int = 32768
    fmt.Println(v8)
    v8 = (^(v8))
    if ((((true && true) || true) || (2147483647 < 2147483647)) && (!false)) {
        var v9 uint = 42
        if (false || false) {
            var v10 float32 = 0.0
            var v11 int8 = -128
            fmt.Println(v8)
            _ = v10
            _ = v11
        }
        _ = len(v7)
        _ = v9
    }
    _ = v7
    _ = v8
    return 442
}
func genFn1(p0 uint16, p1 int64, p2 bool) int16 {
    switch -1 {
    case 234:
        v12 := []int{857, 381, 42}
        _ = v12
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int = fi13
            _ = fi13
            _ = v14
        }
    }
    p2 = p2
    if (4294967295 == 4294967295) {
        var v15 int64 = p1
        if true {
            for fi16 := 0; fi16 < 6; fi16++ {
                if (42 <= 2) {
                    var v17 int64 = p1
                    var v18 int64 = (v15 & v17)
                    _ = v17
                    _ = v18
                } else {
                    switch fi16 {
                    case -1:
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 int8 = 77
                            _ = fi19
                            _ = v20
                        }
                    default:
                        v21 := GenS1{F0: v15, F1: genFn0(), F2: 1}
                        _ = v21
                    }
                    var v22 uint = 255
                    v22 = 1
                    _ = v22
                }
                var v23 uint8 = 184
                _ = fi16
                _ = v23
            }
            v24 := GenS1{F0: p1, F1: 85, F2: 1}
            v25 := []int{}
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: p1, F1: p0, F2: 276}
        _ = v15
        _ = v26
    }
    v27 := GenS1{F0: (p1 ^ p1), F1: p0, F2: 878}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return -391
}

func main() {
    v1 := []uint32{42, 473}
    var v2 bool = (false || ((((float32(((3.14159 - 0.0) / 1.0)) + 3.14159) >= (1.5 - ((((float32((float64(0.0) + float64((0.0 + 1e10)))) * 1e10) + (float32(1.5) + (-2.25 - 3.14159))) - ((-2.25 + 1.5) - (1e10 + 1.5))) - (3.14159 * 1.5)))) || false) && true))
    var v3 int16 = 32767
    var v4 uint64 = 4294967295
    v5 := make(map[int]int)
    var v6 uint = 211
    v1 = append(v1, 164)
    var v28 uint32 = 2
    if false {
        if (!v2) {
            v29 := GenS2{F0: 728}
            fmt.Println(v4)
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        v5[fi30] = fi30
        var v31 string = "hello"
        v1 = append(v1, (v28 & v28))
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        var v33 uint64 = 18446744073709551615
        v34 := GenS3{F0: 107, F1: 9223372036854775807}
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v3 = genFn1(802, 1, ((0 > v28) || ((^(v6)) != (v6 ^ (v6 >> (((^(v6))) & 63))))))
        switch (fi35 ^ fi35) {
        case 1:
            var v36 int = (^(fi35))
            _ = v36
        case 42:
            for fi37 := 0; fi37 < 2; fi37++ {
                v38 := GenS4{F0: (1 < 25), F1: ("foo bar" + "\n\t"), F2: 42}
                switch (fi35 % ((9223372036854775807) | 1)) {
                case 0:
                    v4 = (v4 + v4)
                case 128:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 32767:
                    if true {
                        if false {
                            v39 := GenS5{F0: "hello", F1: 42}
                            if len(v1) > 0 {
                                v1[0] = v28
                            }
                            _ = v39
                        } else {
                            v40 := v5[42]
                            if v2 {
                                _ = len(v1)
                                v28 += (^(v28))
                                _ = len(v5)
                            } else {
                                if v2 {
                                    if (14 < 1) {
                                        var v41 int8 = 100
                                        _ = v41
                                    }
                                    v40 = fi35
                                    var v42 uint32 = 65535
                                    _ = v42
                                }
                            }
                            v43 := GenS5{F0: "\n\t", F1: 1}
                            _ = v40
                            _ = v43
                        }
                        v44 := GenS5{F0: (("hello" + "") + "foo bar"), F1: -518}
                        v6 = (v6 + v6)
                        v6 = v6
                        _ = v44
                    } else {
                        delete(v5, fi35)
                        var v45 int = 382
                        _ = v45
                    }
                default:
                    v46 := GenS5{F0: "", F1: 0}
                    _ = v46
                }
                _ = fi37
                _ = v38
            }
        case -14:
            var v47 int16 = v3
            _ = v47
        }
        _ = fi35
    }
    var v48 uint16 = 2
    fmt.Println(v4, v6, v28)
    for fi49 := 0; fi49 < 13; fi49++ {
        if (v6 != 1) {
            v3 += (v3 * genFn1(v48, -911, ((3.14159 + 1.5) >= (-2.25 - 0.1))))
        } else {
            if v2 {
                v50 := GenS6{F0: 109, F1: 42, F2: fi49}
                _ = v50
            } else {
                for fi51 := 0; fi51 < 9; fi51++ {
                    v5[fi51] = fi51
                    var v52 int = -2147483648
                    _ = fi51
                    _ = v52
                }
                v3 += v3
                v2 = v2
            }
            switch "a" {
            case "":
                if ((0.1 <= -2.25) || (!v2)) {
                    switch "hello" {
                    case "\n\t":
                        if (!true) {
                            v53 := GenS7{F0: -1, F1: 891}
                            v4 -= 18446744073709551615
                            _ = v53
                        } else {
                            for fi54 := 0; fi54 < 6; fi54++ {
                                v55 := GenS8{F0: (3.14159 > 0.1), F1: v2}
                                v56 := v5[(-(fi54))]
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                            v57 := v5[int(v48)]
                            _ = len(v1)
                            _ = v57
                        }
                    case "a":
                        fmt.Println(v3, v6, v4)
                    case "":
                        var v58 int64 = -9223372036854775808
                        _ = v58
                    default:
                        var v59 int = int(v48)
                        _ = v59
                    }
                    fmt.Println(v6, v6, v48)
                    v28 += uint32(v4)
                }
            case "foo bar":
                if (true && v2) {
                    for fi60 := 0; fi60 < 12; fi60++ {
                        v61 := GenS9{F0: v2}
                        v28 -= 124
                        if false {
                            v62 := GenS2{F0: (fi60 ^ (fi60 >> ((v6) & 63)))}
                            switch 576 {
                            case 0:
                                if v2 {
                                    var v63 float32 = 1e10
                                    _ = v63
                                } else {
                                    var v64 float32 = (3.14159 + 0.1)
                                    v65 := GenS5{F0: "\n\t", F1: 1}
                                    v3 += 42
                                    _ = v64
                                    _ = v65
                                }
                            case 509:
                                if v2 {
                                    var v66 float32 = -2.25
                                    if (!v2) {
                                        fmt.Println(v2, v6)
                                        var v67 bool = false
                                        if len(v1) > 0 {
                                            v1[0] = v28
                                        }
                                        _ = v67
                                    }
                                    fmt.Println(fi60, v6)
                                    _ = v66
                                } else {
                                    var v68 int16 = (^(v3))
                                    if ((v28 / (((v28 & uint32(v48))) | 1)) > (^(v28))) {
                                        var v69 uint16 = v48
                                        _ = v69
                                    }
                                    if v2 {
                                        var v70 uint64 = (v4 >> (((v6 / ((4294967295) | 1))) & 63))
                                        var v71 int32 = -2147483648
                                        if false {
                                            if false {
                                                var v72 int64 = -185
                                                fmt.Println(v71, v70)
                                                v3 -= (v68 / (((v3 - -1)) | 1))
                                                _ = v72
                                            } else {
                                                if v2 {
                                                    fmt.Println(v48, fi60, v48)
                                                    fmt.Println(fi60, v68)
                                                }
                                            }
                                        }
                                        _ = v70
                                        _ = v71
                                    }
                                    _ = v68
                                }
                                v6 += 4294967295
                            default:
                                var v73 uint16 = uint16(v28)
                                _ = v73
                            }
                            v4 = uint64(v6)
                            v74 := GenS3{F0: 0, F1: 383}
                            _ = v62
                            _ = v74
                        }
                        _ = fi60
                        _ = v61
                    }
                } else {
                    if v2 {
                        v75 := GenS10{F0: v48}
                        switch fi49 {
                        case 128:
                            if true {
                                v76 := GenS11{F0: int(v3)}
                                for fi77 := 0; fi77 < 12; fi77++ {
                                    switch "hello" {
                                    case "foo bar":
                                        if (v2 || v2) {
                                            v78 := GenS2{F0: (fi77 - (fi49 % ((fi49) | 1)))}
                                            _ = v78
                                        }
                                    default:
                                        var v79 float32 = 1e10
                                        _ = v79
                                    }
                                    _ = fi77
                                }
                                _ = v76
                            } else {
                                var v80 int8 = int8(v48)
                                _ = v80
                            }
                        }
                        if v2 {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v75
                    } else {
                        for fi81 := 0; fi81 < 19; fi81++ {
                            var v82 string = ((("hello" + "\n\t") + ("\n\t" + "hello")) + "foo bar")
                            fmt.Println(fi81, v3)
                            _ = fi81
                            _ = v82
                        }
                        v83 := GenS12{F0: v3, F1: -109, F2: 48}
                        _ = v83
                    }
                    v4 -= 983
                }
            }
        }
        v48 -= (v48 / (((^(v48))) | 1))
        v2 = v2
        _ = fi49
    }
    switch "" {
    case "foo bar":
        v84 := GenS3{F0: 0, F1: 9223372036854775807}
        _ = v84
    case "a":
        var v85 float64 = 1.5
        _ = v85
    case "hello":
        var v86 uint = 0
        _ = v86
    }
    fmt.Println(v48, v4)
    var v87 float64 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v87
}
