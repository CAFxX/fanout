package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 string; F2 uint64 }
type GenS8 struct { F0 uint8; F1 uint64; F2 bool }
type GenS9 struct { F0 int32; F1 uint; F2 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint; F1 bool; F2 uint }
type GenS12 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 42
    var v3 int32 = -1
    v2 = 42
    var v4 uint64 = 631
    delete(v1, -399)
    if ((!false) && (32767 > 2)) {
        for fi5 := 0; fi5 < 17; fi5++ {
            if true {
                fmt.Println(fi5, v3, v3)
            } else {
                var v6 int = fi5
                var v7 uint = 968
                _ = v6
                _ = v7
            }
            for fi8 := 0; fi8 < 14; fi8++ {
                v4 = v4
                v2 = (v2 >> ((2) & 63))
                var v9 uint32 = (v2 & 2147483648)
                v1[(-(fi8))] = (true && ((3.14159 + ((3.14159 - -2.25) + (0.0 * (0.1 + 0.0)))) > 3.14159))
                _ = fi8
                _ = v9
            }
            v3 = -104
            _ = fi5
        }
    }
    var v10 int32 = -669
    v1[-82] = true
    if (false || (-120 <= 127)) {
        for fi11 := 0; fi11 < 2; fi11++ {
            if (false && true) {
                v12 := []int{(fi11 % ((fi11) | 1)), (fi11 | (fi11 % ((fi11) | 1)))}
                fmt.Println(v4, v3)
                _ = v12
            }
            delete(v1, (fi11 * (^(fi11))))
            var v13 uint64 = 0
            _ = fi11
            _ = v13
        }
        if false {
            for fi14 := 0; fi14 < 12; fi14++ {
                for fi15 := 0; fi15 < 3; fi15++ {
                    v3 -= v10
                    switch (fi15 - fi15) {
                    case 32768:
                        v16 := []int{}
                        _ = v16
                    }
                    v17 := GenS1{F0: fi14, F1: 932}
                    v18 := GenS2{F0: v4, F1: 2}
                    v3 = (v3 / ((v10) | 1))
                    _ = fi15
                    _ = v17
                    _ = v18
                }
                _ = fi14
            }
            fmt.Println(v4)
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 int64 = 0
                var v21 float32 = 1e10
                switch fi19 {
                case 32767:
                    if ((24 < 100) || (!true)) {
                        var v22 uint16 = 42
                        if (677 < v4) {
                            var v23 string = "foo bar"
                            _ = v23
                        }
                        if true {
                            fmt.Println(v22, v4)
                            if ((v2 >> ((4294967295) & 63)) < 4294967295) {
                                v1[fi19] = (4294967295 >= (v2 >> ((881) & 63)))
                                fmt.Println(v21, fi19)
                            }
                        } else {
                            if ((v21 > v21) && true) {
                                if ((false || false) || ("\n\t" != "\n\t")) {
                                    switch (-(fi19)) {
                                    case 9223372036854775807:
                                        v24 := GenS3{F0: 2, F1: true, F2: v4}
                                        _ = v24
                                    case 2147483647:
                                        v4 += v4
                                    case 132:
                                        var v25 bool = true
                                        _ = v25
                                    }
                                    var v26 uint8 = 96
                                    for fi27 := 0; fi27 < 20; fi27++ {
                                        var v28 uint16 = (v22 & v22)
                                        switch (("a" + ("\n\t" + "foo bar")) + ("a" + "a")) {
                                        case "hello":
                                            switch "\n\t" {
                                            case "foo bar":
                                                v26 -= (v26 << ((1) & 63))
                                            case "hello":
                                                switch (^(fi19)) {
                                                case 129:
                                                    switch ("\n\t" + "a") {
                                                    case "\n\t":
                                                        v29 := GenS3{F0: (v4 << ((658) & 63)), F1: true, F2: (v4 ^ (^(v4)))}
                                                        _ = v29
                                                    case "a":
                                                        var v30 int64 = -9223372036854775808
                                                        _ = v30
                                                    case "hello":
                                                        var v31 int32 = -1
                                                        _ = v31
                                                    default:
                                                        var v32 uint64 = (v4 >> ((uint(v2)) & 63))
                                                        _ = v32
                                                    }
                                                case -181:
                                                    fmt.Println(v10, v21, v28)
                                                case 4294967295:
                                                    switch (("\n\t" + "") + ("hello" + ("" + ("\n\t" + ("\n\t" + "\n\t"))))) {
                                                    case "foo bar":
                                                        v3 = v3
                                                    }
                                                default:
                                                    v33 := GenS2{F0: v4, F1: 75}
                                                    _ = v33
                                                }
                                            }
                                        case "a":
                                            v2 = (v2 / ((4294967295) | 1))
                                        default:
                                            var v34 float32 = (1.5 / (v21 / ((float32((0.0 + 1.5)) + 3.14159) / float32(float64((((float32(3.14159) * v21) * v21) / v21))))))
                                            _ = v34
                                        }
                                        var v35 int8 = 33
                                        v36 := GenS4{F0: 65535}
                                        if (1000 == 1) {
                                            v37 := GenS1{F0: fi19, F1: 65535}
                                            v4 = v4
                                            switch "foo bar" {
                                            case "":
                                                var v38 uint = 425
                                                if (2 > (^(v38))) {
                                                    if true {
                                                        fmt.Println(v4, v3, v4)
                                                        var v39 uint64 = v4
                                                        _ = v39
                                                    }
                                                    v40 := GenS5{F0: (v35 / ((v35) | 1)), F1: (^(v26))}
                                                    _ = v40
                                                }
                                                var v41 int32 = -2147483648
                                                _ = v38
                                                _ = v41
                                            case "foo bar":
                                                v42 := GenS6{F0: 9223372036854775807}
                                                _ = v42
                                            case "hello":
                                                var v43 float64 = (-2.25 + 0.1)
                                                _ = v43
                                            default:
                                                var v44 int16 = -1
                                                _ = v44
                                            }
                                            _ = v37
                                        } else {
                                            v45 := []int{}
                                            _ = v45
                                        }
                                        _ = fi27
                                        _ = v28
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v26
                                } else {
                                    if true {
                                        var v46 int32 = v10
                                        _ = v46
                                    } else {
                                        v3 += -337
                                        v22 = v22
                                        var v47 string = "hello"
                                        _ = v47
                                    }
                                }
                            }
                            fmt.Println(v10, v21)
                        }
                        _ = v22
                    }
                case 390:
                    var v48 uint32 = 2147483648
                    _ = v48
                case -9223372036854775808:
                    var v49 float32 = 0.0
                    _ = v49
                default:
                    if (false && (127 <= 128)) {
                        v50 := GenS2{F0: 748, F1: 1}
                        _ = v50
                    }
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            if false {
                v51 := GenS7{F0: -379, F1: ("a" + "foo bar"), F2: (v4 + v4)}
                if (42 > uint16(v4)) {
                    var v52 int16 = 42
                    _ = v52
                }
                _ = v51
            } else {
                for fi53 := 0; fi53 < 7; fi53++ {
                    for fi54 := 0; fi54 < 16; fi54++ {
                        var v55 uint16 = 255
                        switch ("foo bar" + "hello") {
                        case "":
                            v56 := GenS8{F0: 1, F1: (v4 / ((v4) | 1)), F2: ((fi54 << ((18446744073709551615) & 63)) > -762)}
                            _ = v56
                        }
                        _ = fi54
                        _ = v55
                    }
                    v57 := []int{fi53}
                    _ = fi53
                    _ = v57
                }
                var v58 uint = 18446744073709551615
                for fi59 := 0; fi59 < 17; fi59++ {
                    if (-32768 > 1000) {
                        switch fi59 {
                        case 144:
                            if true {
                                if false {
                                    v3 = (v10 * int32(v2))
                                    v60 := GenS8{F0: 2, F1: v4, F2: (-128 == -1)}
                                    _ = v60
                                }
                                v61 := GenS7{F0: 1, F1: ("hello" + "foo bar"), F2: (v4 - 2)}
                                v3 = v10
                                _ = v61
                            }
                        case 608:
                            v62 := GenS9{F0: (v10 + -2147483648), F1: v58, F2: v58}
                            _ = v62
                        case -448:
                            v2 -= 4294967295
                        }
                        delete(v1, fi59)
                        var v63 int64 = 542
                        _ = v63
                    }
                    v64 := GenS2{F0: v4, F1: (v58 << ((v58) & 63))}
                    v4 -= 18446744073709551615
                    _ = fi59
                    _ = v64
                }
                _ = v58
            }
        }
        var v65 int16 = 42
        _ = v65
    } else {
        for fi66 := 0; fi66 < 18; fi66++ {
            v67 := GenS10{F0: 440}
            _ = fi66
            _ = v67
        }
    }
    var v68 int = int(v4)
    for fi69 := 0; fi69 < 11; fi69++ {
        v70 := []int{0, v68, v68, v68}
        var v71 float32 = (((((1.5 * 0.0) - (0.0 - 0.1)) * (float32((float64((-2.25 / 3.14159)) - float64((float32(3.14159) * ((0.0 * 0.1) + (3.14159 - 1.5)))))) + (((0.1 + 3.14159) / 0.1) + (0.1 + ((1.5 * 0.0) * (0.0 / float32((1.5 - ((float64((0.0 + (float32(float64(-2.25)) + ((1e10 * 3.14159) + ((((1.5 - 0.1) * 3.14159) + 0.0) / 0.1))))) / 3.14159) - float64(float32(0.1))))))))))) - (3.14159 - (0.0 * -2.25))) * (1.5 / 1.0))
        for fi72 := 0; fi72 < 3; fi72++ {
            fmt.Println(v10)
            _ = fi72
        }
        _ = fi69
        _ = v70
        _ = v71
    }
    v73 := GenS11{F0: 1, F1: true, F2: 1}
    if ((54 != uint8(v2)) || (127 <= -1)) {
        fmt.Println(v2, v2)
        v74 := GenS3{F0: 4294967295, F1: ((!false) && (true && false)), F2: (v4 % (((v4 >> ((0) & 63))) | 1))}
        if (1 <= 935) {
            for fi75 := 0; fi75 < 3; fi75++ {
                v76 := []int{(fi75 | (v68 ^ 32767)), (v68 / ((fi75) | 1))}
                v4 = 112
                _ = fi75
                _ = v76
            }
            fmt.Println(v3, v2)
            switch (v68 - v68) {
            case 299:
                var v77 int = -492
                _ = v77
            }
        } else {
            v78 := GenS12{F0: 1, F1: ((v68 | v68) == v68)}
            _ = v78
        }
        _ = v74
    }
    for fi79 := 0; fi79 < 20; fi79++ {
        v10 = (v3 + v10)
        var v80 int64 = -9223372036854775808
        _ = fi79
        _ = v80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v73
}
