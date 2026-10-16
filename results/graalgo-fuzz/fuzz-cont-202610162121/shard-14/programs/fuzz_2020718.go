package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 uint; F1 uint; F2 uint32 }
type GenS10 struct { F0 uint }

func main() {
    var v1 uint = 4294967295
    var v2 string = "hello"
    var v3 float32 = float32(3.14159)
    var v4 string = v2
    if false {
        v4 = "hello"
    }
    v5 := GenS1{F0: (v1 * (^(v1))), F1: -1, F2: -2147483648}
    if (true || (!false)) {
        for fi6 := 0; fi6 < 7; fi6++ {
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 uint16 = 700
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        if false {
            for fi9 := 0; fi9 < 9; fi9++ {
                switch fi9 {
                case -2147483648:
                    fmt.Println(fi9, v2)
                case 383:
                    var v10 uint32 = 813
                    _ = v10
                }
                _ = fi9
            }
        } else {
            if false {
                for fi11 := 0; fi11 < 9; fi11++ {
                    switch 127 {
                    case -9223372036854775808:
                        v12 := []int{fi11, fi11, fi11}
                        _ = v12
                    case 2147483647:
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 int32 = -2147483648
                            var v15 bool = false
                            var v16 float32 = v3
                            switch v2 {
                            case "":
                                v17 := GenS2{F0: (fi13 << (((v1 & (v1 ^ 4294967295))) & 63)), F1: 872, F2: 65535}
                                _ = v17
                            case "\n\t":
                                if false {
                                    v18 := []int{fi13, (fi13 - -221), (fi13 % ((-416) | 1)), (fi13 / ((fi11) | 1))}
                                    v19 := GenS3{F0: ((fi13 & (fi13 * (fi13 << ((v1) & 63)))) != (-(fi11)))}
                                    _ = v18
                                    _ = v19
                                } else {
                                    v20 := GenS3{F0: v15}
                                    var v21 uint32 = 2147483647
                                    v22 := []int{(-(fi11)), fi11, (-(fi13))}
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                            default:
                                var v23 uint = v1
                                _ = v23
                            }
                            var v24 uint64 = 865
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v24
                        }
                    default:
                        var v25 string = v2
                        _ = v25
                    }
                    var v26 float64 = float64(1e10)
                    _ = fi11
                    _ = v26
                }
                var v27 uint8 = 140
                _ = v27
            } else {
                for fi28 := 0; fi28 < 14; fi28++ {
                    if true {
                        switch (fi28 / ((812) | 1)) {
                        case -882:
                            if false {
                                for fi29 := 0; fi29 < 20; fi29++ {
                                    var v30 float64 = 0.0
                                    v31 := GenS1{F0: 65535, F1: (fi29 / ((fi28) | 1)), F2: -2147483648}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                        case -9223372036854775808:
                            var v32 int = (-(fi28))
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 17; fi33++ {
                            if (!true) {
                                fmt.Println(fi33, fi33, fi28)
                                switch (fi33 | (fi28 / (((-(fi33))) | 1))) {
                                case 4294967295:
                                    var v34 float32 = float32(float64(v3))
                                    _ = v34
                                case -633:
                                    v1 = 137
                                case -216:
                                    v1 = (v1 << ((921) & 63))
                                default:
                                    var v35 uint32 = 202
                                    _ = v35
                                }
                                v36 := []int{fi28}
                                fmt.Println(fi28)
                                _ = v36
                            } else {
                                v2 = (v4 + ((v2 + "a") + "a"))
                            }
                            v37 := []int{fi28, fi33, 4294967295}
                            _ = fi33
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 1; fi38++ {
                            v3 = v3
                            _ = fi38
                        }
                    } else {
                        v39 := []int{}
                        for fi40 := 0; fi40 < 6; fi40++ {
                            var v41 float64 = (0.0 + (3.14159 * 0.1))
                            var v42 uint64 = 4294967295
                            var v43 float64 = (v41 - v41)
                            var v44 float64 = (v41 - float64((((v3 * (3.14159 + (v3 / v3))) + v3) + v3)))
                            _ = fi40
                            _ = v41
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        v45 := GenS4{F0: 42, F1: 18446744073709551615}
                        _ = v39
                        _ = v45
                    }
                    var v46 string = (("foo bar" + v4) + v4)
                    fmt.Println(v46)
                    _ = fi28
                    _ = v46
                }
                var v47 uint64 = 0
                _ = v47
            }
            fmt.Println(v2)
        }
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        v49 := GenS5{F0: uint32(v1)}
        _ = fi48
        _ = v49
    }
    if (2 > 28) {
        fmt.Println(v1, v1)
    } else {
        v50 := []int{4294967295, 860, -163, 875}
        fmt.Println(v4)
        _ = v50
    }
    v51 := GenS6{F0: 2, F1: 42}
    v2 = "foo bar"
    for fi52 := 0; fi52 < 2; fi52++ {
        fmt.Println(v3, v4, v2)
        var v53 float64 = (0.1 / 0.1)
        var v54 int = -1
        _ = fi52
        _ = v53
        _ = v54
    }
    for fi55 := 0; fi55 < 11; fi55++ {
        for fi56 := 0; fi56 < 17; fi56++ {
            var v57 float64 = (((1.5 / 1.0) * 0.0) / ((0.0 + float64(float32(0.1))) - 1e10))
            var v58 int8 = -128
            var v59 string = v4
            _ = fi56
            _ = v57
            _ = v58
            _ = v59
        }
        _ = fi55
    }
    v1 += 452
    var v60 string = (("" + (("foo bar" + "hello") + v4)) + (((v4 + "a") + v4) + v2))
    v61 := []int{}
    if (2147483647 != -316) {
        if true {
            v62 := GenS4{F0: 0, F1: 18446744073709551615}
            if (2 >= 127) {
                v4 = ((v2 + "foo bar") + v60)
                v60 = v4
            }
            v4 = ("a" + "\n\t")
            _ = v62
        } else {
            v63 := GenS1{F0: (v1 / ((647) | 1)), F1: 128, F2: 0}
            if len(v61) > 0 {
                v61[0] = 417
            }
            for fi64 := 0; fi64 < 13; fi64++ {
                if (false || (!(42 > 32767))) {
                    if (v1 > (v1 | v1)) {
                        v1 = v1
                    }
                    for fi65 := 0; fi65 < 13; fi65++ {
                        switch ("a" + ((v2 + (((v60 + v4) + (v4 + ("foo bar" + ""))) + "a")) + (("a" + v60) + (v4 + ((v60 + v60) + "a"))))) {
                        case "":
                            if len(v61) >= 2 {
                                v61 = v61[1:len(v61)]
                            }
                        }
                        v2 = "hello"
                        v66 := GenS7{F0: 127, F1: 599}
                        _ = fi65
                        _ = v66
                    }
                    fmt.Println(v4, v4)
                } else {
                    switch int(v1) {
                    case -147:
                        fmt.Println(v1, v60)
                    default:
                        if ((false && (255 <= 263)) || (158 <= 9223372036854775807)) {
                            if len(v61) > 0 {
                                v61[0] = (^(fi64))
                            }
                        } else {
                            v2 = v60
                            if true {
                                switch ("a" + v4) {
                                case "\n\t":
                                    var v67 int64 = 11
                                    _ = v67
                                case "":
                                    var v68 int32 = 42
                                    _ = v68
                                }
                                var v69 int64 = -1
                                v70 := GenS6{F0: 2, F1: -1}
                                _ = v69
                                _ = v70
                            }
                            var v71 uint32 = 0
                            _ = v71
                        }
                    }
                    if len(v61) > 0 {
                        v61[0] = -873
                    }
                    fmt.Println(v60, v1)
                }
                for fi72 := 0; fi72 < 11; fi72++ {
                    fmt.Println(v4, v4, v1)
                    _ = fi72
                }
                if len(v61) >= 2 {
                    v61 = v61[1:len(v61)]
                }
                _ = fi64
            }
            _ = v63
        }
        if false {
            if (-788 <= 32767) {
                if false {
                    switch -569 {
                    case -376:
                        v73 := GenS8{F0: -87, F1: 43}
                        _ = v73
                    case 2147483647:
                        switch v4 {
                        case "\n\t":
                            var v74 uint16 = 32767
                            _ = v74
                        default:
                            switch 32768 {
                            case -2147483648:
                                v75 := GenS9{F0: v1, F1: 0, F2: 2147483648}
                                _ = v75
                            case 32767:
                                for fi76 := 0; fi76 < 8; fi76++ {
                                    v77 := GenS6{F0: 65535, F1: 42}
                                    _ = fi76
                                    _ = v77
                                }
                            }
                        }
                    case 9223372036854775807:
                        var v78 int = 32767
                        _ = v78
                    default:
                        v79 := GenS10{F0: v1}
                        _ = v79
                    }
                }
            }
            var v80 uint64 = 18446744073709551615
            _ = v80
        }
    } else {
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v60
    _ = v61
}
