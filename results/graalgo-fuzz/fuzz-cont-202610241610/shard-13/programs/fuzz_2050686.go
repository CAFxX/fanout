package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 uint; F1 uint16 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int; F1 uint8; F2 uint32 }
func genFn0(p0 uint32) uint32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 42, F1: -1}
        if (!(94 >= 127)) {
            if (142 < 127) {
                var v9 int16 = 1
                switch 638 {
                case 127:
                    fmt.Println(v9)
                case 1:
                    v10 := GenS2{F0: false}
                    _ = v10
                case -695:
                    v11 := GenS2{F0: false}
                    _ = v11
                }
                _ = v9
            }
        }
        switch (fi7 * fi7) {
        case -9223372036854775808:
            var v12 int8 = 42
            _ = v12
        case 641:
            for fi13 := 0; fi13 < 20; fi13++ {
                if false {
                    p0 = (^(p0))
                } else {
                    p0 = (p0 / ((517) | 1))
                    switch (fi7 + -9223372036854775808) {
                    case 4294967295:
                        v14 := []int{(fi13 ^ (fi13 & fi7)), fi7}
                        _ = v14
                    default:
                        switch (("\n\t" + ("" + "a")) + "foo bar") {
                        case "foo bar":
                            if true {
                                var v15 int32 = 1
                                v16 := GenS3{F0: false, F1: 1000}
                                _ = v15
                                _ = v16
                            } else {
                                if (1.5 <= 1e10) {
                                    switch 1 {
                                    case 4294967295:
                                        var v17 float64 = (((0.1 * 1e10) * 0.0) + 1e10)
                                        _ = v17
                                    }
                                } else {
                                    var v18 int16 = 42
                                    _ = v18
                                }
                                v19 := GenS3{F0: (true || false), F1: -945}
                                var v20 int8 = -74
                                _ = v19
                                _ = v20
                            }
                        default:
                            switch "" {
                            case "hello":
                                v21 := []int{fi13, (^(fi7))}
                                _ = v21
                            case "":
                                if true {
                                    fmt.Println(fi7)
                                } else {
                                    var v22 float64 = (-2.25 + 0.1)
                                    if false {
                                        v23 := GenS1{F0: -128, F1: 0}
                                        switch "\n\t" {
                                        case "hello":
                                            v22 = (v22 - (0.0 + float64(1e10)))
                                        }
                                        _ = v23
                                    } else {
                                        v24 := []int{-9223372036854775808, fi13}
                                        var v25 uint8 = 0
                                        _ = v24
                                        _ = v25
                                    }
                                    switch (fi7 >> ((18446744073709551615) & 63)) {
                                    case 519:
                                        var v26 uint16 = 0
                                        _ = v26
                                    case 2147483647:
                                        switch int(p0) {
                                        case 2147483647:
                                            switch "foo bar" {
                                            case "":
                                                switch "a" {
                                                case "foo bar":
                                                    var v27 int64 = -878
                                                    _ = v27
                                                }
                                            case "\n\t":
                                                var v28 int8 = 100
                                                _ = v28
                                            case "foo bar":
                                                var v29 int8 = 100
                                                _ = v29
                                            }
                                        case 0:
                                            v30 := GenS3{F0: (3.14159 != 3.14159), F1: 1000}
                                            _ = v30
                                        case 42:
                                            if false {
                                                v22 = (3.14159 * 3.14159)
                                            } else {
                                                var v31 string = (("foo bar" + "foo bar") + "hello")
                                                if true {
                                                    v32 := []int{(fi7 - fi7), fi13, fi13, 28}
                                                    _ = v32
                                                }
                                                _ = v31
                                            }
                                        }
                                    }
                                    _ = v22
                                }
                            case "\n\t":
                                p0 -= p0
                            }
                        }
                    }
                    switch "" {
                    case "a":
                        v33 := GenS3{F0: true, F1: 1000}
                        _ = v33
                    case "\n\t":
                        var v34 uint16 = 693
                        _ = v34
                    default:
                        var v35 uint16 = 42
                        _ = v35
                    }
                }
                p0 += (p0 ^ (^(p0)))
                v36 := GenS4{F0: 969}
                _ = fi13
                _ = v36
            }
        default:
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 uint16 = uint16(p0)
                _ = fi37
                _ = v38
            }
        }
        _ = fi7
        _ = v8
    }
    p0 -= (p0 % ((65535) | 1))
    v39 := GenS5{F0: 160, F1: -181}
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 string = (("hello" + "hello") + "\n\t")
        _ = fi40
        _ = v41
    }
    _ = p0
    _ = v39
    return (p0 & (p0 | 0))
}
func genFn1(p0 uint8, p1 uint16) uint32 {
    fmt.Println(p0, p0)
    switch -726 {
    case 2147483647:
        v42 := GenS6{F0: -2147483648, F1: genFn0(42), F2: p0}
        _ = v42
    }
    _ = p0
    _ = p1
    return genFn0(genFn0(4294967295))
}

func main() {
    var v1 int16 = 271
    var v2 float64 = 0.0
    v3 := []string{}
    v4 := make(map[string]int)
    var v5 int = 36
    v6 := []uint32{4294967295, 973, 32}
    fmt.Println(v5, v1, v2)
    if true {
        var v43 uint64 = 2
        v44 := GenS5{F0: 42, F1: 455}
        fmt.Println(v5)
        _ = v43
        _ = v44
    } else {
        switch 128 {
        case -48:
            v1 -= (v1 | v1)
        case 665:
            switch "hello" {
            case "":
                for fi45 := 0; fi45 < 11; fi45++ {
                    if false {
                        v46 := GenS6{F0: 42, F1: 685, F2: 255}
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 4; fi47++ {
                            if (genFn0(42) != 65535) {
                                fmt.Println(v2, v1)
                                v48 := GenS7{F0: (!false), F1: 534}
                                _ = v48
                            } else {
                                var v49 int16 = (v1 ^ v1)
                                _ = v49
                            }
                            v5 = (fi45 / ((-1) | 1))
                            var v50 bool = false
                            var v51 float64 = 1.5
                            if v50 {
                                v2 = ((0.1 * v2) / v2)
                                v52 := GenS6{F0: 42, F1: 4294967295, F2: 244}
                                v53 := GenS1{F0: 42, F1: 2147483647}
                                _ = v52
                                _ = v53
                            }
                            _ = fi47
                            _ = v50
                            _ = v51
                        }
                    }
                    var v54 float64 = float64(3.14159)
                    if true {
                        switch "hello" {
                        case "hello":
                            v55 := GenS2{F0: (true && (true && false))}
                            _ = v55
                        case "":
                            for fi56 := 0; fi56 < 14; fi56++ {
                                switch 4294967295 {
                                case 1:
                                    if true {
                                        v1 -= v1
                                    } else {
                                        var v57 int16 = (v1 & v1)
                                        _ = v57
                                    }
                                case 2147483647:
                                    if true {
                                        switch (fi45 >> ((417) & 63)) {
                                        case 80:
                                            v2 += (v2 + 0.0)
                                        default:
                                            if true {
                                                fmt.Println(v54, v2, v54)
                                            }
                                        }
                                        if true {
                                            var v58 uint16 = 42
                                            v59 := GenS8{F0: 638, F1: 834, F2: (v58 & 65535)}
                                            if true {
                                                v60 := v4["hello"]
                                                fmt.Println(v5)
                                                _ = v60
                                            }
                                            _ = v58
                                            _ = v59
                                        } else {
                                            v1 -= 42
                                            if false {
                                                switch "" {
                                                case "a":
                                                    v5 -= fi56
                                                case "\n\t":
                                                    v61 := GenS9{F0: 149, F1: -117}
                                                    _ = v61
                                                case "":
                                                    if false {
                                                        fmt.Println(v54)
                                                        fmt.Println(v54, fi56, fi56)
                                                    }
                                                }
                                            }
                                        }
                                    }
                                case 128:
                                    _ = len(v6)
                                default:
                                    var v62 uint64 = 42
                                    _ = v62
                                }
                                var v63 uint32 = genFn1(195, 32767)
                                var v64 uint64 = 0
                                switch fi56 {
                                case 42:
                                    v54 = (v2 * 1e10)
                                case 2147483647:
                                    v64 += v64
                                default:
                                    var v65 uint16 = 0
                                    _ = v65
                                }
                                _ = fi56
                                _ = v63
                                _ = v64
                            }
                        default:
                            v66 := GenS10{F0: 158, F1: 42}
                            _ = v66
                        }
                        v67 := GenS11{F0: 0, F1: true}
                        _ = v67
                    }
                    _ = fi45
                    _ = v54
                }
            case "foo bar":
                var v68 uint32 = genFn1(255, 255)
                _ = v68
            case "a":
                _ = len(v4)
            }
        default:
            v69 := GenS4{F0: 679}
            _ = v69
        }
    }
    v2 -= 1.5
    var v70 uint = 2
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 146 {
    case 32768:
        var v71 uint64 = 1
        _ = v71
    case 99:
        delete(v4, (("\n\t" + "foo bar") + ""))
    }
    if len(v6) > 0 {
        v6[0] = 859
    }
    v72 := GenS12{F0: 2}
    for fi73 := 0; fi73 < 20; fi73++ {
        var v74 int16 = (v1 >> ((v70) & 63))
        var v75 uint32 = 2147483648
        if false {
            var v76 float64 = float64(1.5)
            var v77 int32 = 0
            _ = v76
            _ = v77
        } else {
            fmt.Println(v2, v5, v75)
            v1 = v74
        }
        _ = fi73
        _ = v74
        _ = v75
    }
    var v78 uint16 = uint16(v70)
    var v79 float64 = -2.25
    if ((!(!false)) && true) {
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "a")
        }
        var v80 int = -1
        v81 := GenS12{F0: 756}
        _ = v80
        _ = v81
    } else {
        if ((v78 >> ((v70) & 63)) > (v78 ^ (v78 ^ 1))) {
            v82 := GenS13{F0: v5, F1: 1, F2: uint32(v78)}
            var v83 uint8 = 2
            _ = v82
            _ = v83
        } else {
            var v84 uint64 = 0
            _ = v84
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v70:", v70)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v72
    _ = v78
    _ = v79
}
