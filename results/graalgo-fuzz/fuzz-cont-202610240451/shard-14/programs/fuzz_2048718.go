package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint8; F1 int }
func genFn0() uint8 {
    v6 := GenS1{F0: 853, F1: (true || (((65535 <= 132) && true) || false)), F2: ("hello" + "\n\t")}
    v7 := []int{-1, 9223372036854775807, 0, -9223372036854775808}
    var v8 uint64 = 18446744073709551615
    _ = len(v7)
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    _ = v6
    _ = v7
    _ = v8
    return 220
}
func genFn1() int8 {
    v9 := GenS1{F0: 42, F1: true, F2: "foo bar"}
    var v10 uint = 467
    switch 9223372036854775807 {
    case 383:
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(fi11, fi11)
            v10 += v10
            v12 := []int{2147483647, 9223372036854775807, (fi11 % (((^(fi11))) | 1))}
            _ = fi11
            _ = v12
        }
    case 581:
        switch 2147483647 {
        case 1:
            v10 -= 431
        default:
            if false {
                switch 0 {
                case 4294967295:
                    var v13 string = "\n\t"
                    _ = v13
                }
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := []int{(^(fi14))}
                    if true {
                        v16 := GenS1{F0: -415, F1: ((0.1 * 1.5) < -2.25), F2: "\n\t"}
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v10 += (v10 / (((v10 & (v10 + (v10 ^ v10)))) | 1))
                            _ = fi17
                        }
                    }
                    var v18 int = 896
                    _ = fi14
                    _ = v15
                    _ = v18
                }
            }
        }
    case 2147483647:
        var v19 string = "a"
        _ = v19
    }
    _ = v10
    _ = v9
    return -123
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 int8 = 26
    var v5 string = "foo bar"
    v1 = (v1 << ((819) & 63))
    for fi20 := 0; fi20 < 12; fi20++ {
        v1 = (^(v1))
        var v21 uint16 = 255
        _ = fi20
        _ = v21
    }
    fmt.Println(v4)
    v22 := GenS1{F0: 820, F1: (!(0 >= 1)), F2: v5}
    fmt.Println(v5)
    var v23 bool = true
    v24 := GenS1{F0: 2147483647, F1: v23, F2: (((v5 + v5) + v5) + "hello")}
    var v25 uint8 = (^(v1))
    var v26 int8 = -17
    if v23 {
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS1{F0: 715, F1: v23, F2: v5}
            switch (fi27 | 2147483647) {
            case -2147483648:
                fmt.Println(v5, v5, v5)
            }
            v26 = 0
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 uint = 255
            _ = fi29
            _ = v30
        }
        if (0 != 65535) {
            var v31 int32 = 1
            for fi32 := 0; fi32 < 19; fi32++ {
                if true {
                    v23 = (v23 && ((1.5 / (3.14159 * 3.14159)) == -2.25))
                } else {
                    delete(v3, v5)
                    var v33 int32 = (v31 >> ((18446744073709551615) & 63))
                    _ = v33
                }
                fmt.Println(fi32, v4, v23)
                _ = fi32
            }
            _ = v31
        } else {
            switch "" {
            case "\n\t":
                _ = len(v3)
            }
            v4 = v26
            var v34 uint64 = 2
            _ = v34
        }
    }
    v35 := GenS1{F0: 0, F1: v23, F2: ((("\n\t" + ((v5 + (v5 + "hello")) + (v5 + v5))) + "hello") + ("hello" + (v5 + v5)))}
    v36 := GenS1{F0: -657, F1: (!(9223372036854775807 >= 32767)), F2: v5}
    for fi37 := 0; fi37 < 11; fi37++ {
        if (false && false) {
            v23 = v23
            if false {
                v4 = 0
                v3[v5] = 4294967295
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 int8 = (v4 ^ 0)
                    fmt.Println(v25, fi38, fi38)
                    _ = fi38
                    _ = v39
                }
            } else {
                delete(v3, "foo bar")
            }
            var v40 float32 = float32((3.14159 * 1.5))
            _ = v40
        }
        if ((^(fi37)) <= 0) {
            v4 += (^(v4))
            switch ("a" + v5) {
            case "a":
                for fi41 := 0; fi41 < 2; fi41++ {
                    if true {
                        switch v5 {
                        case "a":
                            var v42 uint8 = (v25 % ((255) | 1))
                            _ = v42
                        case "":
                            _ = len(v2)
                        default:
                            switch fi37 {
                            case 32767:
                                v5 = v5
                            default:
                                var v43 int = (fi37 & (fi37 + fi41))
                                _ = v43
                            }
                        }
                        if v23 {
                            v44 := GenS2{F0: v5, F1: 42}
                            var v45 uint16 = 255
                            _ = v44
                            _ = v45
                        } else {
                            v25 -= (v25 / ((2) | 1))
                        }
                        if (!true) {
                            fmt.Println(v1, v25, v23)
                            v26 = v26
                            if (false && v23) {
                                v23 = ((!((fi41 ^ fi37) >= (fi41 - (fi41 - (fi37 | (fi41 >> ((255) & 63))))))) && true)
                            }
                            fmt.Println(fi41, v5)
                            v46 := GenS1{F0: 1, F1: (33 < 65), F2: (v5 + "foo bar")}
                            var v47 bool = v23
                            _ = v46
                            _ = v47
                        }
                    } else {
                        switch fi37 {
                        case 543:
                            var v48 int16 = -781
                            _ = v48
                        }
                    }
                    if (!(v5 >= v5)) {
                        var v49 uint = 718
                        v50 := GenS3{F0: 2147483647}
                        v1 = (v1 - uint8(v49))
                        _ = v49
                        _ = v50
                    }
                    if (-9223372036854775808 == 9223372036854775807) {
                        v25 = 127
                        if v23 {
                            fmt.Println(v1, v5, fi37)
                            if v23 {
                                var v51 int8 = 100
                                _ = v51
                            }
                        } else {
                            switch v5 {
                            case "foo bar":
                                if ("hello" > v5) {
                                    var v52 float32 = 1e10
                                    if v23 {
                                        var v53 float32 = 0.1
                                        var v54 float64 = 3.14159
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v25 = (^(v25))
                                        v26 = (-(v26))
                                    }
                                    v55 := []int{(fi41 - 679)}
                                    _ = v52
                                    _ = v55
                                }
                            }
                        }
                        v56 := []int{(fi37 << ((255) & 63)), fi37, (^(fi37))}
                        _ = v56
                    }
                    _ = fi41
                }
            case "\n\t":
                for fi57 := 0; fi57 < 12; fi57++ {
                    if (v1 < genFn0()) {
                        v58 := GenS4{F0: (v1 << ((0) & 63)), F1: 0}
                        _ = v58
                    }
                    _ = fi57
                }
            }
        } else {
            var v59 int32 = 1
            for fi60 := 0; fi60 < 3; fi60++ {
                v5 = v5
                switch (fi60 - -690) {
                case 656:
                    delete(v2, 2)
                case 9223372036854775807:
                    switch (fi37 << ((42) & 63)) {
                    case 436:
                        v4 += 42
                    case 2147483647:
                        switch (fi37 << ((4294967295) & 63)) {
                        case -586:
                            v23 = v23
                        default:
                            v5 = ("a" + ("a" + v5))
                        }
                    case 0:
                        v1 -= (^(v1))
                    }
                case 2147483647:
                    var v61 uint16 = 0
                    _ = v61
                default:
                    var v62 int8 = int8(v59)
                    _ = v62
                }
                _ = fi60
            }
            if (v23 && ((true && v23) || false)) {
                for fi63 := 0; fi63 < 14; fi63++ {
                    v64 := GenS5{F0: (v25 % ((42) | 1)), F1: fi37}
                    v1 -= v1
                    var v65 int8 = v4
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            } else {
                v4 = v26
                for fi66 := 0; fi66 < 1; fi66++ {
                    var v67 int64 = -416
                    if (531 >= 0) {
                        var v68 uint64 = 972
                        _ = v68
                    } else {
                        v23 = v23
                        switch "a" {
                        case "hello":
                            fmt.Println(v1)
                            switch v5 {
                            case "\n\t":
                                if false {
                                    v26 = v4
                                }
                            default:
                                var v69 uint32 = 42
                                _ = v69
                            }
                            v67 += v67
                        }
                        v4 -= (v26 / (((v4 / ((-58) | 1))) | 1))
                    }
                    _ = fi66
                    _ = v67
                }
                for fi70 := 0; fi70 < 11; fi70++ {
                    var v71 uint = 296
                    if true {
                        v72 := GenS4{F0: (^(v25)), F1: -388}
                        _ = v72
                    } else {
                        var v73 int8 = (v26 % (((v4 | v26)) | 1))
                        v74 := []int{fi37, (fi70 << (((v71 & (^(v71)))) & 63)), int(v4), fi70}
                        v71 = (v71 / (((v71 / ((v71) | 1))) | 1))
                        _ = v73
                        _ = v74
                    }
                    _ = fi70
                    _ = v71
                }
            }
            _ = v59
        }
        _ = fi37
    }
    v75 := []int{0, 1, -439}
    v25 -= (v1 / (((v25 % ((v25) | 1))) | 1))
    if (-1 > 2147483647) {
        if (!(-572 == -9223372036854775808)) {
            v23 = v23
        } else {
            var v76 uint32 = 366
            _ = v76
        }
        var v77 int64 = -9223372036854775808
        _ = v77
    }
    switch 0 {
    case 128:
        var v78 float32 = (3.14159 * 0.0)
        _ = v78
    case 564:
        fmt.Println(v26, v5)
    case 505:
        fmt.Println(v5, v23, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v75:", len(v75), v75)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v75
}
