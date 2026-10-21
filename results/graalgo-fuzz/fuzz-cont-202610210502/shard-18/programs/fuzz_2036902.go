package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 uint64; F1 int64 }
func genFn0(p0 int64, p1 bool) uint {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 uint = 2
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        p1 = p1
        _ = fi6
    }
    if (p1 || false) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{fi7, fi7}
            _ = fi7
            _ = v8
        }
        var v9 int64 = -9223372036854775808
        _ = v9
    }
    v10 := GenS1{F0: 265}
    var v11 string = (("a" + "hello") + "\n\t")
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 2
}
func genFn1(p0 uint8, p1 int32, p2 uint) uint8 {
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS1{F0: 255}
            v15 := []int{(-(fi13))}
            switch fi12 {
            case 4294967295:
                var v16 bool = true
                _ = v16
            case -9223372036854775808:
                _ = len(v15)
            default:
                if (276 == 540) {
                    v17 := GenS2{F0: 42, F1: 465}
                    var v18 int = (fi13 * (fi12 | fi13))
                    var v19 int8 = 63
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    switch "hello" {
                    case "hello":
                        v15 = append(v15, fi12)
                    case "\n\t":
                        p2 += (p2 >> (((p2 & p2)) & 63))
                    default:
                        fmt.Println(p2, p2, fi13)
                    }
                    if true {
                        if (3.14159 == (0.1 - 1e10)) {
                            v20 := GenS3{F0: 1, F1: -991}
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            var v21 uint64 = 1
                            _ = v20
                            _ = v21
                        }
                    } else {
                        var v22 uint8 = p0
                        fmt.Println(p2, p0, p2)
                        _ = v22
                    }
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        v23 := []int{fi12}
        p2 = 18446744073709551615
        _ = fi12
        _ = v23
    }
    p2 = p2
    v24 := []int{-9223372036854775808, 593}
    for fi25 := 0; fi25 < 11; fi25++ {
        p1 = 2147483647
        for fi26 := 0; fi26 < 13; fi26++ {
            fmt.Println(p1, p2, fi26)
            _ = fi26
        }
        fmt.Println(p2, p2, fi25)
        _ = fi25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return p0
}

func main() {
    var v1 string = ("foo bar" + ("a" + ""))
    v2 := []uint32{58, 278, 65535, 263}
    v3 := make(map[int]string)
    switch "" {
    case "":
        fmt.Println(v1)
    default:
        fmt.Println(v1)
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 string = (v1 + v1)
            v29 = ""
            var v30 float32 = 0.0
            _ = fi28
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            if (!false) {
                fmt.Println(fi27, fi27, fi31)
            } else {
                fmt.Println(fi31, fi27)
                fmt.Println(fi31)
            }
            if (true && (!true)) {
                v3[fi31] = ((v1 + (((v1 + v1) + v1) + v1)) + "")
                switch "\n\t" {
                case "\n\t":
                    var v32 int64 = 1
                    _ = v32
                case "foo bar":
                    switch v1 {
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                    case "\n\t":
                        switch (-(fi27)) {
                        case 431:
                            _ = len(v2)
                            var v33 uint8 = 255
                            var v34 uint8 = genFn1((v33 >> ((18446744073709551615) & 63)), 2147483647, genFn0(-9223372036854775808, true))
                            _ = v33
                            _ = v34
                        case 582:
                            var v35 int32 = -1
                            _ = v35
                        case -1:
                            var v36 int16 = 32767
                            v1 = v1
                            _ = v36
                        default:
                            v1 = (((v1 + ("\n\t" + ("foo bar" + v1))) + "hello") + v1)
                        }
                    case "hello":
                        var v37 int8 = 0
                        _ = v37
                    default:
                        var v38 int32 = 42
                        _ = v38
                    }
                case "a":
                    var v39 bool = (494 <= 0)
                    _ = v39
                default:
                    v40 := v3[(fi27 % (((fi27 << ((0) & 63))) | 1))]
                    _ = v40
                }
                if (57 != 1) {
                    v1 = v1
                } else {
                    v1 = v1
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                }
            }
            v1 = ((v1 + v1) + v1)
            _ = fi31
        }
        _ = fi27
    }
    if (true || (255 <= 65535)) {
        if true {
            if (false || (true || (!true))) {
                v1 = ""
            }
        }
        var v41 uint16 = 227
        if (!true) {
            v42 := GenS4{F0: 100, F1: (v1 + v1)}
            _ = v42
        }
        _ = v41
    } else {
        for fi43 := 0; fi43 < 9; fi43++ {
            v1 = "foo bar"
            _ = fi43
        }
        var v44 int8 = 53
        for fi45 := 0; fi45 < 18; fi45++ {
            v44 = (v44 | (v44 + v44))
            _ = fi45
        }
        _ = v44
    }
    if true {
        var v46 int = 2147483647
        _ = v46
    }
    v1 = "\n\t"
    v47 := v3[200]
    for fi48 := 0; fi48 < 5; fi48++ {
        fmt.Println(fi48, v1)
        if false {
            v49 := GenS3{F0: 0, F1: 42}
            fmt.Println(v1, v1, v47)
            v50 := GenS5{F0: -971}
            _ = v49
            _ = v50
        }
        v1 = ((v1 + v47) + v1)
        _ = fi48
    }
    var v51 uint8 = 128
    switch v47 {
    case "":
        if ((794 > 4294967295) || (-2147483648 >= 1)) {
            var v52 float64 = 3.14159
            fmt.Println(v51, v47)
            if false {
                var v53 int8 = 42
                v54 := v3[-675]
                for fi55 := 0; fi55 < 9; fi55++ {
                    switch fi55 {
                    case 4294967295:
                        if (235 < 850) {
                            for fi56 := 0; fi56 < 15; fi56++ {
                                v51 -= v51
                                var v57 int64 = -9223372036854775808
                                var v58 float32 = float32(0.0)
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                            v59 := GenS6{F0: true}
                            _ = v59
                        } else {
                            v60 := GenS7{F0: (v53 >= 42), F1: 1}
                            _ = v60
                        }
                    }
                    v61 := GenS8{F0: -1, F1: -295}
                    _ = fi55
                    _ = v61
                }
                _ = v53
                _ = v54
            } else {
                v51 += 0
                var v62 uint64 = 2
                v51 = (v51 << ((uint(v51)) & 63))
                _ = v62
            }
            _ = v52
        } else {
            v63 := GenS5{F0: -405}
            _ = v63
        }
    }
    if (!false) {
        v64 := GenS9{F0: 0, F1: 0}
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi65 := 0; fi65 < 10; fi65++ {
                for fi66 := 0; fi66 < 3; fi66++ {
                    v47 = (v47 + "hello")
                    _ = fi66
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi65
            }
        case 127:
            fmt.Println(v1, v1, v47)
        case 389:
            v1 = v1
        }
        var v67 bool = (663 <= 683)
        _ = v64
        _ = v67
    }
    switch v47 {
    case "":
        fmt.Println(v51, v47)
    case "a":
        v51 -= v51
    case "foo bar":
        _ = len(v3)
    default:
        switch 32767 {
        case 1:
            var v68 uint = 176
            _ = v68
        case -9223372036854775808:
            var v69 float32 = (0.0 / (0.1 + 1e10))
            _ = v69
        }
    }
    var v70 string = v47
    v71 := v3[-275]
    switch 4294967295 {
    case 128:
        _ = len(v2)
    case 0:
        var v72 uint32 = 0
        _ = v72
    case -2147483648:
        var v73 int = 483
        _ = v73
    default:
        for fi74 := 0; fi74 < 4; fi74++ {
            if true {
                for fi75 := 0; fi75 < 15; fi75++ {
                    if (!(false && (((1 < -442) || true) && (127 <= (fi75 << ((18446744073709551615) & 63)))))) {
                        switch fi75 {
                        case 198:
                            var v76 int16 = -32768
                            _ = v76
                        case -291:
                            if (1e10 == 0.0) {
                                if (127 > 127) {
                                    var v77 int = (fi75 * (fi75 | fi75))
                                    _ = v77
                                }
                            }
                        case -406:
                            var v78 float64 = ((3.14159 / 0.1) * float64(1.5))
                            _ = v78
                        }
                        var v79 int = (fi75 * fi75)
                        if (v51 != (^(v51))) {
                            var v80 int32 = 2147483647
                            _ = v80
                        }
                        _ = v79
                    }
                    v71 = (v47 + v1)
                    _ = fi75
                }
                v1 = ((v70 + v71) + (v71 + (v70 + (v71 + v1))))
                for fi81 := 0; fi81 < 20; fi81++ {
                    if (v70 < "") {
                        var v82 uint64 = 578
                        _ = v82
                    } else {
                        v70 = "a"
                        var v83 string = v47
                        var v84 int32 = 216
                        v83 = ("" + ("foo bar" + ("a" + v70)))
                        _ = len(v3)
                        _ = v83
                        _ = v84
                    }
                    if false {
                        v71 = v71
                        _ = len(v3)
                        var v85 int32 = 0
                        _ = v85
                    }
                    var v86 uint = 4294967295
                    var v87 int8 = 127
                    var v88 uint8 = v51
                    var v89 uint16 = 0
                    _ = fi81
                    _ = v86
                    _ = v87
                    _ = v88
                    _ = v89
                }
            }
            fmt.Println(fi74, v51, v71)
            _ = fi74
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v51
    _ = v70
    _ = v71
}
