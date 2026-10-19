package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int16; F1 uint16; F2 uint16 }
func genFn0(p0 int8, p1 int8, p2 uint8) uint8 {
    v4 := []int{32768, 0, 0}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}
func genFn1(p0 uint64, p1 uint16, p2 int64) bool {
    v5 := GenS1{F0: 42}
    p0 = (p0 ^ p0)
    switch ("" + "") {
    case "":
        var v6 string = ""
        _ = v6
    case "hello":
        v7 := GenS1{F0: genFn0(-1, 54, 0)}
        _ = v7
    case "foo bar":
        if ((((float64(((1.5 - 0.1) * (0.1 / 1.0))) / (1e10 + -2.25)) >= (3.14159 - float64(-2.25))) || true) || (!((1e10 >= ((1.5 * (0.1 - -2.25)) - ((float32(0.0) * -2.25) * 1.5))) && (0 <= 0)))) {
            v8 := []int{-1, 32768, 398}
            _ = v8
        } else {
            v9 := GenS1{F0: 53}
            v10 := GenS2{F0: ((("\n\t" + "foo bar") + "") != ""), F1: 0, F2: 0}
            v11 := GenS3{F0: 42}
            _ = v10
            _ = v11
            _ = v9
        }
    default:
        var v12 string = "\n\t"
        _ = v12
    }
    p2 -= (p2 << ((1) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (!((((true && false) || false) || ((p0 & (p0 << ((42) & 63))) != (p0 % ((p0) | 1)))) && (42 == -1)))
}

func main() {
    var v1 uint32 = 0
    var v2 int16 = -757
    var v3 bool = (!true)
    var v13 int = -1
    v14 := []int{v13, v13, (^(v13)), 127}
    var v15 uint8 = genFn0(127, -1, 0)
    for fi16 := 0; fi16 < 9; fi16++ {
        _ = len(v14)
        if (478 < 635) {
            v3 = v3
            var v17 uint64 = 270
            _ = v17
        }
        v18 := GenS2{F0: v3, F1: (v2 + v2), F2: 214}
        _ = fi16
        _ = v18
    }
    var v19 bool = v3
    var v20 int = (v13 << ((42) & 63))
    var v21 int8 = 1
    if len(v14) > 0 {
        v14[0] = (v13 * (v13 >> ((18446744073709551615) & 63)))
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        switch "" {
        case "foo bar":
            v23 := GenS3{F0: v2}
            _ = v23
        case "hello":
            v1 = (v1 >> ((1) & 63))
        }
        _ = fi22
    }
    v24 := GenS3{F0: v2}
    v15 = (v15 & v15)
    v19 = v3
    v19 = (!v19)
    if (804 != 42) {
        if (int16(v15) == v2) {
            var v25 int64 = 1
            var v26 int64 = v25
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 6; fi27++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                v15 += v15
                _ = fi27
            }
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            switch (^(v20)) {
            case -2147483648:
                v2 -= v2
            case 620:
                v29 := GenS4{F0: -1, F1: (^(v15)), F2: v21}
                _ = v29
            default:
                var v30 uint = 0
                _ = v30
            }
            _ = fi28
        }
        var v31 int32 = 42
        _ = v31
    } else {
        for fi32 := 0; fi32 < 16; fi32++ {
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 uint8 = v15
                _ = fi33
                _ = v34
            }
            if v19 {
                for fi35 := 0; fi35 < 19; fi35++ {
                    if false {
                        var v36 int16 = v2
                        fmt.Println(v13, v15, v21)
                        var v37 int8 = v21
                        _ = v36
                        _ = v37
                    } else {
                        if (2 > 1) {
                            fmt.Println(v13, v3, v1)
                        }
                        _ = len(v14)
                        switch (fi32 * (^(v20))) {
                        case 550:
                            var v38 int8 = 100
                            _ = v38
                        case 32768:
                            var v39 float32 = (1.5 * -2.25)
                            _ = v39
                        case -9223372036854775808:
                            v40 := GenS1{F0: v15}
                            _ = v40
                        default:
                            if v19 {
                                v41 := GenS3{F0: v2}
                                v21 -= v21
                                _ = v41
                            } else {
                                if v3 {
                                    var v42 uint16 = 0
                                    if ((255 > 435) || (v19 && v19)) {
                                        if ((584 != 723) || v3) {
                                            var v43 bool = v19
                                            v42 -= v42
                                            fmt.Println(fi32, fi32, fi35)
                                            _ = v43
                                        }
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    } else {
                                        v42 += v42
                                        v3 = v19
                                        var v44 uint8 = (^(v15))
                                        _ = v44
                                    }
                                    if (42 < v21) {
                                        if (v19 && false) {
                                            switch 127 {
                                            case -9223372036854775808:
                                                fmt.Println(v13, v1, v19)
                                            case 32768:
                                                v3 = true
                                            case 128:
                                                v45 := GenS5{F0: 1}
                                                _ = v45
                                            default:
                                                var v46 bool = (!true)
                                                _ = v46
                                            }
                                        } else {
                                            fmt.Println(v13, v2)
                                            v47 := GenS6{F0: (v13 / ((fi32) | 1)), F1: 18446744073709551615}
                                            var v48 uint16 = v42
                                            switch "a" {
                                            case "hello":
                                                var v49 bool = v19
                                                _ = v49
                                            case "a":
                                                var v50 uint64 = 1
                                                _ = v50
                                            }
                                            v19 = false
                                            _ = v47
                                            _ = v48
                                        }
                                        var v51 uint8 = (v15 << ((255) & 63))
                                        var v52 bool = v19
                                        _ = v51
                                        _ = v52
                                    }
                                    v13 = -842
                                    _ = v42
                                }
                            }
                        }
                    }
                    _ = len(v14)
                    switch "\n\t" {
                    case "a":
                        var v53 uint = 42
                        _ = v53
                    case "hello":
                        var v54 uint8 = v15
                        _ = v54
                    }
                    _ = fi35
                }
                v55 := GenS7{F0: v2, F1: 1, F2: 32767}
                _ = v55
            } else {
                switch "foo bar" {
                case "":
                    v15 -= v15
                default:
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v14 = append(v14, v20)
                        fmt.Println(fi56, v20, v21)
                        switch "foo bar" {
                        case "\n\t":
                            if (("\n\t" + ("" + "a")) != "\n\t") {
                                var v57 uint64 = 18446744073709551615
                                var v58 uint = 687
                                _ = v57
                                _ = v58
                            }
                        case "foo bar":
                            var v59 int8 = (^(v21))
                            _ = v59
                        default:
                            v60 := GenS2{F0: (32767 != 42), F1: v2, F2: 255}
                            _ = v60
                        }
                        _ = fi56
                    }
                }
                switch ((("foo bar" + "foo bar") + "") + "\n\t") {
                case "hello":
                    switch 4294967295 {
                    case -940:
                        var v61 string = ""
                        _ = v61
                    case 509:
                        if (v21 == v21) {
                            for fi62 := 0; fi62 < 19; fi62++ {
                                if ((v1 & 727) != (v1 / ((v1) | 1))) {
                                    if true {
                                        fmt.Println(v13, v13, fi32)
                                        v20 -= (v13 * -459)
                                    } else {
                                        switch ("hello" + "foo bar") {
                                        case "\n\t":
                                            v14 = append(v14, (fi62 + 9223372036854775807))
                                        case "a":
                                            v1 -= 939
                                        case "hello":
                                            if (791 != 2147483647) {
                                                if ((42 <= -2147483648) || (false && true)) {
                                                    v3 = true
                                                    var v63 float32 = 3.14159
                                                    if ((^(v1)) <= (v1 * v1)) {
                                                        v3 = v19
                                                        v63 = float32(((((float64(v63) / -2.25) - ((-2.25 + 0.0) + 0.1)) * (-2.25 / float64(3.14159))) - 0.1))
                                                    } else {
                                                        v2 = -851
                                                        if ((v21 * 100) != (v21 / (((v21 % (((v21 << ((134) & 63))) | 1))) | 1))) {
                                                            v1 -= (v1 >> ((18446744073709551615) & 63))
                                                            var v64 int64 = 9223372036854775807
                                                            _ = v64
                                                        }
                                                        var v65 uint = 530
                                                        _ = v65
                                                    }
                                                    _ = v63
                                                } else {
                                                    fmt.Println(fi32, v3)
                                                    _ = len(v14)
                                                }
                                                if (18446744073709551615 >= 377) {
                                                    v14 = append(v14, (-(fi62)))
                                                    _ = len(v14)
                                                }
                                                v13 += (v20 & v20)
                                            }
                                        }
                                    }
                                    if (!((((^(v20)) != (v13 ^ (v13 / ((128) | 1)))) && (0.0 != (3.14159 - 0.0))) || false)) {
                                        var v66 uint16 = 65535
                                        v13 = fi62
                                        _ = v66
                                    }
                                }
                                _ = len(v14)
                                var v67 uint64 = 1
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                v2 = (^(v2))
                                _ = fi62
                                _ = v67
                            }
                            fmt.Println(v21, v15, v19)
                        }
                    default:
                        var v68 int = (fi32 - 32768)
                        _ = v68
                    }
                case "foo bar":
                    if (4294967295 == 0) {
                        for fi69 := 0; fi69 < 18; fi69++ {
                            v3 = v3
                            _ = fi69
                        }
                        if (42 != -1) {
                            for fi70 := 0; fi70 < 10; fi70++ {
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                switch ("" + "hello") {
                                case "a":
                                    var v71 int16 = v2
                                    _ = v71
                                case "":
                                    v2 += int16(v13)
                                case "foo bar":
                                    var v72 uint = 0
                                    _ = v72
                                }
                                _ = fi70
                            }
                            switch "foo bar" {
                            case "foo bar":
                                var v73 int64 = -1
                                _ = v73
                            case "\n\t":
                                v1 = (v1 * v1)
                            case "":
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            default:
                                v20 -= (v13 - (v20 / ((v20) | 1)))
                            }
                        } else {
                            for fi74 := 0; fi74 < 10; fi74++ {
                                switch (((("hello" + ("hello" + "\n\t")) + "") + "foo bar") + "") {
                                case "hello":
                                    v3 = false
                                }
                                fmt.Println(v1, v19)
                                v2 -= 42
                                _ = fi74
                            }
                        }
                        for fi75 := 0; fi75 < 13; fi75++ {
                            fmt.Println(v1)
                            v1 = 0
                            var v76 uint8 = v15
                            _ = fi75
                            _ = v76
                        }
                    } else {
                        v15 = v15
                    }
                }
                v19 = (0.0 > 0.0)
            }
            _ = fi32
        }
    }
    var v77 int16 = v2
    fmt.Println(v15)
    v78 := GenS2{F0: (4294967295 < 879), F1: (^(v2)), F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v77
    _ = v78
}
