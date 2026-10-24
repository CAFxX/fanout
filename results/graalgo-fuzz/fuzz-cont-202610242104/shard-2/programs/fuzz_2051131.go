package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int8; F1 int32 }
func genFn0(p0 int64, p1 int8, p2 uint32) string {
    for fi4 := 0; fi4 < 8; fi4++ {
        switch ("foo bar" + "\n\t") {
        case "a":
            v5 := []int{fi4, (fi4 / ((fi4) | 1)), -416}
            _ = v5
        case "":
            p2 = 42
        }
        _ = fi4
    }
    v6 := []int{0, -1, -9223372036854775808, 118}
    if (true || (-2147483648 < -9223372036854775808)) {
        var v7 uint32 = 1
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("foo bar" + "a")
}
func genFn1(p0 uint64, p1 uint, p2 uint32) int64 {
    if ((^(p0)) <= 1) {
        v8 := GenS1{F0: 4294967295, F1: 20}
        _ = v8
    } else {
        v9 := []int{127}
        switch 127 {
        case 1:
            v10 := GenS1{F0: -922, F1: 127}
            _ = v10
        }
        _ = v9
    }
    p2 += p2
    if false {
        v11 := []int{4294967295}
        v12 := GenS1{F0: 127, F1: -128}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (0.1 + 0.1)
    v3 := make(map[uint32]uint32)
    fmt.Println(v2)
    v2 = 1.5
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: (fi13 - (fi13 ^ 32768)), F1: 0}
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 785, F1: -128}
    switch -1 {
    case -258:
        v16 := GenS1{F0: 32768, F1: 127}
        _ = v16
    case 2147483647:
        fmt.Println(v2)
    case 420:
        v2 = (0.1 + 1e10)
    default:
        v2 = -2.25
    }
    fmt.Println(v2)
    var v17 string = "\n\t"
    var v18 bool = (0 == 0)
    if v18 {
        v18 = v18
        v19 := []int{-225, 798}
        _ = v19
    } else {
        fmt.Println(v2, v2, v2)
        var v20 int8 = -128
        v17 = v17
        _ = v20
    }
    v17 = v17
    if v18 {
        var v21 uint8 = 2
        fmt.Println(v21, v17)
        _ = v21
    } else {
        var v22 bool = (v18 && (-892 >= genFn1(42, 42, 559)))
        if v22 {
            switch "\n\t" {
            case "a":
                if ((true && false) && true) {
                    v2 = 0.0
                    if (-9223372036854775808 == -957) {
                        for fi23 := 0; fi23 < 20; fi23++ {
                            if (!v22) {
                                v22 = v22
                            } else {
                                var v24 bool = (0 <= 1)
                                var v25 bool = v22
                                var v26 uint8 = 127
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            delete(v3, 871)
                            var v27 uint32 = 332
                            _ = fi23
                            _ = v27
                        }
                        v28 := []int{}
                        _ = v28
                    }
                    var v29 string = "\n\t"
                    _ = v29
                } else {
                    v22 = (!(0 >= 0))
                }
            case "":
                if v18 {
                    if (0.0 >= (((((-2.25 + float64(0.1)) - 1.5) * 0.1) + v2) + v2)) {
                        fmt.Println(v18)
                        if (-75 < 42) {
                            var v30 uint8 = 127
                            switch 42 {
                            case 22:
                                v31 := GenS1{F0: -565, F1: -116}
                                _ = v31
                            }
                            _ = v30
                        }
                    }
                    var v32 bool = v18
                    _ = v32
                }
            case "\n\t":
                var v33 int8 = -128
                _ = v33
            default:
                v34 := GenS2{F0: -128, F1: 719, F2: (v17 + ("foo bar" + ""))}
                _ = v34
            }
        }
        _ = v22
    }
    var v35 int16 = 1
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := GenS3{F0: 100, F1: -2147483648}
        for fi38 := 0; fi38 < 12; fi38++ {
            v39 := GenS1{F0: (fi38 / (((fi38 % ((2147483647) | 1))) | 1)), F1: 42}
            fmt.Println(v2, fi36)
            if v18 {
                var v40 int16 = 1
                v41 := GenS3{F0: 42, F1: 653}
                _ = v40
                _ = v41
            } else {
                if (0 < 97) {
                    v35 = 1
                    v2 = ((0.1 + v2) - (v2 + 0.0))
                    v18 = true
                } else {
                    var v42 uint8 = 204
                    v43 := GenS4{F0: 65535, F1: 0}
                    if true {
                        v17 = "foo bar"
                        if v18 {
                            v44 := GenS3{F0: 42, F1: -2147483648}
                            v42 = v42
                            var v45 float32 = (((0.0 / 1.0) + 0.0) - 0.1)
                            _ = v44
                            _ = v45
                        } else {
                            var v46 float64 = (0.1 / v2)
                            _ = v46
                        }
                    } else {
                        var v47 int16 = 1000
                        _ = v47
                    }
                    fmt.Println(v18, v2)
                    switch (v17 + (("" + (v17 + v17)) + v17)) {
                    case "":
                        var v48 int64 = 9223372036854775807
                        _ = v48
                    case "a":
                        if v18 {
                            var v49 string = v17
                            _ = v49
                        }
                    case "foo bar":
                        fmt.Println(fi36, v18)
                    }
                    _ = v42
                    _ = v43
                }
                _ = len(v3)
                var v50 int32 = 912
                v35 = v35
                _ = v50
            }
            var v51 bool = v18
            switch v17 {
            case "a":
                fmt.Println(v51, v35, fi36)
            case "":
                v52 := []int{fi36, fi36, fi38}
                _ = v52
            case "foo bar":
                var v53 int16 = v35
                _ = v53
            default:
                v2 = (v2 * (v2 * v2))
            }
            _ = fi38
            _ = v39
            _ = v51
        }
        var v54 float64 = v2
        _ = fi36
        _ = v37
        _ = v54
    }
    v55 := GenS4{F0: 65535, F1: -745}
    if v18 {
        for fi56 := 0; fi56 < 4; fi56++ {
            switch fi56 {
            case -872:
                var v57 string = (v17 + v17)
                _ = v57
            }
            v58 := []int{(fi56 << ((862) & 63)), 127}
            _ = fi56
            _ = v58
        }
    } else {
        var v59 int8 = 58
        if (!(!(false || true))) {
            v60 := GenS5{F0: (v59 - (v59 / ((v59) | 1))), F1: 42}
            for fi61 := 0; fi61 < 3; fi61++ {
                for fi62 := 0; fi62 < 14; fi62++ {
                    v35 -= v35
                    v63 := GenS1{F0: fi61, F1: -73}
                    _ = fi62
                    _ = v63
                }
                var v64 int8 = (v59 >> ((1) & 63))
                var v65 float64 = float64(0.0)
                _ = fi61
                _ = v64
                _ = v65
            }
            _ = v60
        } else {
            v66 := GenS3{F0: 42, F1: -2147483648}
            v67 := GenS2{F0: (v59 / ((v59) | 1)), F1: 676, F2: ("" + v17)}
            for fi68 := 0; fi68 < 2; fi68++ {
                var v69 string = ""
                if (!false) {
                    if false {
                        if true {
                            v2 -= 3.14159
                        } else {
                            for fi70 := 0; fi70 < 8; fi70++ {
                                v2 = float64((0.1 / (0.1 - (0.1 / 0.1))))
                                switch 32768 {
                                case 42:
                                    switch v69 {
                                    case "\n\t":
                                        v69 = v17
                                    default:
                                        fmt.Println(v17, v17, v2)
                                    }
                                default:
                                    var v71 uint8 = 1
                                    _ = v71
                                }
                                fmt.Println(fi68, v2)
                                _ = fi70
                            }
                            fmt.Println(v59)
                        }
                        switch (fi68 >> ((152) & 63)) {
                        case -505:
                            switch v17 {
                            case "hello":
                                v59 = v59
                            default:
                                if v18 {
                                    v72 := GenS4{F0: 1, F1: 0}
                                    _ = v72
                                } else {
                                    v73 := v1[""]
                                    fmt.Println(v73)
                                    _ = v73
                                }
                            }
                        case 4294967295:
                            v74 := GenS1{F0: fi68, F1: -32}
                            _ = v74
                        default:
                            v75 := []int{-1, fi68, 2147483647, fi68}
                            _ = v75
                        }
                        v2 = 3.14159
                    }
                    var v76 uint64 = 2
                    _ = v76
                } else {
                    var v77 int = fi68
                    v78 := []int{32768, (fi68 % ((-2147483648) | 1))}
                    _ = v77
                    _ = v78
                }
                _ = fi68
                _ = v69
            }
            _ = v66
            _ = v67
        }
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v55
}
