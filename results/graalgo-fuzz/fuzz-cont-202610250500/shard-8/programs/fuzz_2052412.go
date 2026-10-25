package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string; F1 uint8; F2 int }
type GenS5 struct { F0 int; F1 int8; F2 uint16 }
func genFn0(p0 int16, p1 int16, p2 int32) uint32 {
    v4 := []int{}
    p1 = int16(p2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 float64 = -2.25
    if (255 == 0) {
        var v6 int8 = 127
        if false {
            for fi7 := 0; fi7 < 11; fi7++ {
                p1 = (p0 / ((-32768) | 1))
                _ = fi7
            }
            v5 = v5
            for fi8 := 0; fi8 < 16; fi8++ {
                if false {
                    p2 = -1
                } else {
                    v9 := GenS1{F0: (p0 >> ((0) & 63))}
                    switch "a" {
                    case "\n\t":
                        for fi10 := 0; fi10 < 17; fi10++ {
                            fmt.Println(p1, p2)
                            var v11 int16 = (p0 & (p1 / ((42) | 1)))
                            _ = fi10
                            _ = v11
                        }
                    case "hello":
                        p1 -= p1
                    case "":
                        p0 = (^(p0))
                    }
                    var v12 int32 = (p2 >> ((51) & 63))
                    _ = v12
                    _ = v9
                }
                v13 := GenS1{F0: p0}
                _ = fi8
                _ = v13
            }
        }
        _ = v6
    } else {
        v5 -= -2.25
        for fi14 := 0; fi14 < 1; fi14++ {
            p2 = 737
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 328
}
func genFn1(p0 uint) int32 {
    var v15 uint = (p0 % ((p0) | 1))
    p0 = (^(p0))
    if false {
        p0 = (^(p0))
        var v16 uint64 = 154
        _ = v16
    }
    if (134 > 943) {
        switch -1 {
        case -531:
            fmt.Println(v15, v15)
        case -2147483648:
            if false {
                v17 := GenS1{F0: 1}
                fmt.Println(p0, p0)
                fmt.Println(v15, p0)
                _ = v17
            } else {
                var v18 string = ""
                if true {
                    p0 -= (^(v15))
                } else {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        p0 -= v15
                        var v20 int = 127
                        v21 := GenS1{F0: -728}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v18 = (v18 + (v18 + "hello"))
                }
                _ = v18
            }
        case 127:
            switch 128 {
            case 32768:
                var v22 string = "a"
                _ = v22
            case -197:
                v23 := GenS1{F0: -432}
                _ = v23
            case -9223372036854775808:
                var v24 uint = (^(p0))
                _ = v24
            default:
                p0 += 18446744073709551615
            }
        default:
            switch (("a" + ("hello" + ("foo bar" + "\n\t"))) + "\n\t") {
            case "":
                var v25 uint64 = 928
                _ = v25
            case "foo bar":
                v15 -= v15
            default:
                v15 += (^(p0))
            }
        }
    } else {
        for fi26 := 0; fi26 < 11; fi26++ {
            p0 = p0
            var v27 uint = v15
            _ = fi26
            _ = v27
        }
    }
    _ = p0
    _ = v15
    return -1
}

func main() {
    var v1 uint8 = 2
    var v2 int8 = 1
    v3 := make(map[uint32]uint32)
    var v28 float32 = (1e10 / 3.14159)
    if false {
        for fi29 := 0; fi29 < 13; fi29++ {
            if false {
                delete(v3, 2147483648)
                if (2147483647 > 42) {
                    v30 := []int{fi29, (^(fi29)), (fi29 * fi29), 4294967295}
                    if false {
                        for fi31 := 0; fi31 < 6; fi31++ {
                            if (genFn1(56) == genFn1(416)) {
                                if (genFn0(42, -1, 42) >= 350) {
                                    if (!false) {
                                        fmt.Println(v1, fi31, v28)
                                    }
                                } else {
                                    if false {
                                        var v32 int32 = 0
                                        _ = v32
                                    } else {
                                        if (!(32767 <= 42)) {
                                            var v33 string = "foo bar"
                                            if false {
                                                var v34 uint8 = 128
                                                fmt.Println(v33, v33, v33)
                                                _ = v34
                                            } else {
                                                v35 := GenS2{F0: 1, F1: 65535, F2: -78}
                                                var v36 float64 = float64((1.5 - 1.5))
                                                v37 := GenS3{F0: (v1 ^ 7), F1: 42}
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            }
                                            _ = v33
                                        } else {
                                            switch ("" + "\n\t") {
                                            case "hello":
                                                var v38 int = fi29
                                                _ = v38
                                            }
                                            v39 := GenS4{F0: "hello", F1: (v1 / ((v1) | 1)), F2: (fi31 | 0)}
                                            v40 := GenS1{F0: 1000}
                                            _ = v39
                                            _ = v40
                                        }
                                        v30 = append(v30, (fi31 | fi29))
                                    }
                                }
                                if (true && false) {
                                    switch (("a" + "foo bar") + "hello") {
                                    case "a":
                                        if ((-2.25 * 1e10) > 1e10) {
                                            v28 = (v28 * ((v28 + ((float32((1.5 / 3.14159)) / v28) / v28)) / v28))
                                            _ = len(v30)
                                        } else {
                                            v1 = (v1 % ((v1) | 1))
                                            fmt.Println(v1, v1, v28)
                                            if true {
                                                var v41 bool = false
                                                var v42 string = "\n\t"
                                                fmt.Println(v42)
                                                _ = v41
                                                _ = v42
                                            } else {
                                                if false {
                                                    if true {
                                                        v43 := v3[65535]
                                                        if true {
                                                            var v44 int = 42
                                                            _ = len(v30)
                                                            _ = v44
                                                        }
                                                        _ = v43
                                                    }
                                                    _ = len(v3)
                                                    switch ("hello" + ("" + "foo bar")) {
                                                    case "\n\t":
                                                        delete(v3, 2147483647)
                                                    case "":
                                                        v45 := GenS2{F0: -962, F1: 2, F2: 42}
                                                        if (!(float64(v28) >= (1e10 * ((0.0 + -2.25) - 3.14159)))) {
                                                            var v46 int32 = -2147483648
                                                            _ = v46
                                                        }
                                                        v28 += (3.14159 - v28)
                                                        _ = v45
                                                    case "hello":
                                                        if len(v30) > 0 {
                                                            v30[0] = 724
                                                        }
                                                    }
                                                    var v47 int16 = -32768
                                                    fmt.Println(v2, fi29)
                                                    _ = v47
                                                }
                                                if len(v30) >= 2 {
                                                    v30 = v30[1:len(v30)]
                                                }
                                            }
                                        }
                                    case "":
                                        _ = len(v3)
                                    case "hello":
                                        v28 = 3.14159
                                    }
                                }
                                if true {
                                    var v48 int64 = 910
                                    var v49 uint16 = 361
                                    v3[118] = 65535
                                    _ = v48
                                    _ = v49
                                }
                            }
                            v2 = 91
                            v3[2] = 2147483648
                            _ = fi31
                        }
                    } else {
                        var v50 uint16 = 381
                        _ = v50
                    }
                    _ = v30
                }
            } else {
                fmt.Println(v1, v28, v2)
                for fi51 := 0; fi51 < 8; fi51++ {
                    v52 := GenS4{F0: "foo bar", F1: v1, F2: 42}
                    _ = fi51
                    _ = v52
                }
            }
            var v53 int32 = -2147483648
            _ = fi29
            _ = v53
        }
        v54 := GenS5{F0: -774, F1: -5, F2: 255}
        _ = v54
    } else {
        var v55 int = 42
        _ = v55
    }
    for fi56 := 0; fi56 < 17; fi56++ {
        for fi57 := 0; fi57 < 12; fi57++ {
            v1 = v1
            if false {
                var v58 uint8 = (v1 - v1)
                _ = v58
            }
            var v59 bool = true
            _ = fi57
            _ = v59
        }
        _ = fi56
    }
    v1 = 42
    v28 += v28
    v60 := GenS1{F0: 0}
    switch ("hello" + ("hello" + "")) {
    case "\n\t":
        v28 = v28
    default:
        for fi61 := 0; fi61 < 6; fi61++ {
            var v62 int32 = -2147483648
            _ = fi61
            _ = v62
        }
    }
    var v63 uint16 = 65535
    switch "a" {
    case "foo bar":
        var v64 int8 = 127
        _ = v64
    }
    fmt.Println(v28, v28, v63)
    v65 := []int{128, 88, 4294967295, -858}
    for fi66 := 0; fi66 < 19; fi66++ {
        delete(v3, 42)
        for fi67 := 0; fi67 < 9; fi67++ {
            v3[358] = 2147483648
            switch -3 {
            case 1:
                switch (((("a" + "hello") + "hello") + "hello") + "foo bar") {
                case "hello":
                    v68 := GenS3{F0: (v1 + v1), F1: 255}
                    _ = v68
                case "foo bar":
                    var v69 uint8 = 128
                    _ = v69
                }
            case 32768:
                _ = len(v3)
            }
            v1 = v1
            _ = fi67
        }
        if false {
            var v70 int64 = 42
            _ = v70
        }
        _ = fi66
    }
    _ = len(v65)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v60
    _ = v63
    _ = v65
}
