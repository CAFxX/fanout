package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 string }
type GenS4 struct { F0 bool; F1 bool; F2 bool }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int; F1 uint64; F2 uint64 }
func genFn0() uint32 {
    switch "foo bar" {
    case "":
        var v7 uint64 = 97
        _ = v7
    case "hello":
        var v8 int16 = 624
        _ = v8
    case "foo bar":
        if false {
            var v9 int32 = 160
            fmt.Println(v9)
            if (false && true) {
                v9 = (^(v9))
                var v10 uint = 18446744073709551615
                v11 := []int{-9223372036854775808, 0, 2147483647}
                _ = v10
                _ = v11
            }
            _ = v9
        } else {
            var v12 int32 = 578
            v12 += -437
            _ = v12
        }
    default:
        v13 := []int{617, -1, 2147483647, 128}
        _ = v13
    }
    var v14 float32 = -2.25
    _ = v14
    return 2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 186
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 uint32 = 2
    var v6 uint16 = 781
    v2 -= -421
    if (18446744073709551615 >= 453) {
        fmt.Println(v6)
        v2 -= 2147483647
    }
    v5 -= (^(v5))
    var v15 float64 = 1e10
    switch (("foo bar" + (("a" + "hello") + "foo bar")) + "hello") {
    case "\n\t":
        var v16 uint8 = 1
        _ = v16
    case "hello":
        v5 = genFn0()
    default:
        fmt.Println(v2)
    }
    var v17 uint32 = v5
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := []int{(fi18 >> ((113) & 63)), (fi18 / ((0) | 1)), -254, fi18}
        fmt.Println(v17)
        var v20 uint16 = (v6 - (v6 << ((198) & 63)))
        _ = fi18
        _ = v19
        _ = v20
    }
    v5 = (v17 ^ v5)
    var v21 uint16 = 1
    v22 := GenS1{F0: (v2 | (v2 & v2)), F1: "hello", F2: 100}
    switch "\n\t" {
    case "\n\t":
        if true {
            v23 := GenS2{F0: -128}
            var v24 uint = 429
            v21 -= 1
            _ = v23
            _ = v24
        }
    case "":
        if true {
            switch 32767 {
            case 50:
                v21 = (v6 + (v21 + (v21 % ((v21) | 1))))
            case -816:
                v25 := []int{}
                _ = v25
            case 724:
                switch ("" + "foo bar") {
                case "foo bar":
                    v26 := GenS2{F0: -128}
                    _ = v26
                case "a":
                    delete(v1, "foo bar")
                }
            }
            if false {
                switch 32768 {
                case 1:
                    v27 := GenS1{F0: v2, F1: ("hello" + "hello"), F2: 115}
                    _ = v27
                case 128:
                    var v28 int16 = 1000
                    _ = v28
                case 0:
                    var v29 int32 = v2
                    _ = v29
                default:
                    var v30 int8 = -1
                    _ = v30
                }
                if true {
                    var v31 int32 = v2
                    v31 -= 0
                    var v32 float64 = v15
                    _ = v31
                    _ = v32
                } else {
                    v33 := GenS3{F0: (!(false || false)), F1: 1, F2: "\n\t"}
                    v34 := GenS4{F0: false, F1: (!(true || ("\n\t" > "a"))), F2: true}
                    fmt.Println(v6, v15, v2)
                    _ = v33
                    _ = v34
                }
                if (!false) {
                    var v35 uint64 = 42
                    _ = v35
                }
            } else {
                v36 := GenS1{F0: (^(v2)), F1: ((("foo bar" + "foo bar") + "foo bar") + "\n\t"), F2: 100}
                _ = v36
            }
            if (v21 >= (^(v6))) {
                if (!(678 <= 135)) {
                    var v37 float32 = (0.0 * ((0.0 / (0.1 - 3.14159)) + (1.5 + float32(float64(-2.25)))))
                    var v38 int = 0
                    _ = v37
                    _ = v38
                }
                v39 := []int{}
                var v40 int = 128
                _ = v39
                _ = v40
            } else {
                v2 -= (v2 + v2)
                if true {
                    switch "foo bar" {
                    case "a":
                        switch "\n\t" {
                        case "\n\t":
                            v41 := GenS5{F0: 0, F1: 2, F2: (v5 & (v17 ^ (v5 << ((4294967295) & 63))))}
                            _ = v41
                        case "a":
                            _ = len(v3)
                        case "foo bar":
                            var v42 int32 = -1
                            _ = v42
                        default:
                            if false {
                                var v43 uint = 612
                                var v44 int8 = 127
                                _ = v43
                                _ = v44
                            } else {
                                var v45 uint16 = v6
                                delete(v3, 816)
                                _ = v45
                            }
                        }
                    case "\n\t":
                        var v46 int64 = -1
                        _ = v46
                    case "foo bar":
                        for fi47 := 0; fi47 < 8; fi47++ {
                            var v48 uint = 1
                            _ = fi47
                            _ = v48
                        }
                    default:
                        v15 = (v15 / (v15 - v15))
                    }
                    for fi49 := 0; fi49 < 1; fi49++ {
                        switch (fi49 + (^(fi49))) {
                        case 2147483647:
                            v50 := GenS1{F0: 42, F1: ("\n\t" + ("\n\t" + ("\n\t" + "hello"))), F2: -23}
                            _ = v50
                        case 127:
                            for fi51 := 0; fi51 < 12; fi51++ {
                                var v52 int64 = -9223372036854775808
                                var v53 uint = 255
                                v54 := []int{-1, fi51, (fi49 >> (((v53 / ((525) | 1))) & 63))}
                                _ = fi51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        case 128:
                            if false {
                                v15 = (float64(0.0) + (v15 / v15))
                            }
                        default:
                            _ = len(v1)
                        }
                        _ = fi49
                    }
                } else {
                    v55 := GenS6{F0: 2}
                    if (true || true) {
                        var v56 float64 = (v15 - v15)
                        if true {
                            v57 := []int{-2147483648, 1, 344}
                            switch "foo bar" {
                            case "a":
                                var v58 uint32 = 2147483648
                                _ = v58
                            case "":
                                var v59 int16 = -939
                                _ = v59
                            }
                            _ = v57
                        }
                        _ = v56
                    } else {
                        if (!(149 >= -367)) {
                            var v60 float64 = (((float64(float32((0.1 + v15))) - (v15 + (1.5 * (v15 / v15)))) + v15) - v15)
                            fmt.Println(v2)
                            fmt.Println(v15, v21)
                            _ = v60
                        } else {
                            v2 = (v2 >> ((255) & 63))
                        }
                    }
                    _ = v55
                }
            }
        }
    case "foo bar":
        var v61 int16 = -32768
        _ = v61
    }
    if (!((false || false) && (-2.25 > 3.14159))) {
        v62 := []int{127, 4294967295}
        var v63 uint64 = 67
        _ = v62
        _ = v63
    }
    if false {
        switch 522 {
        case -465:
            for fi64 := 0; fi64 < 15; fi64++ {
                v65 := []int{}
                switch "" {
                case "\n\t":
                    for fi66 := 0; fi66 < 6; fi66++ {
                        var v67 int8 = -128
                        _ = fi66
                        _ = v67
                    }
                case "foo bar":
                    v21 += v6
                case "hello":
                    v68 := GenS3{F0: ((62 == 38) && false), F1: 0, F2: "foo bar"}
                    _ = v68
                default:
                    for fi69 := 0; fi69 < 5; fi69++ {
                        v21 = v21
                        v65 = append(v65, 32767)
                        _ = fi69
                    }
                }
                var v70 uint64 = 735
                _ = fi64
                _ = v65
                _ = v70
            }
        case 0:
            v71 := GenS5{F0: 255, F1: 2, F2: (v17 ^ uint32(v21))}
            _ = v71
        case -433:
            v5 = v17
        default:
            switch "foo bar" {
            case "a":
                v21 = v21
            }
        }
        for fi72 := 0; fi72 < 20; fi72++ {
            for fi73 := 0; fi73 < 11; fi73++ {
                v6 = (^(v6))
                if true {
                    v74 := GenS7{F0: -1, F1: 0}
                    _ = v74
                } else {
                    v75 := GenS5{F0: 127, F1: 42, F2: (v5 * (v5 + v5))}
                    var v76 float64 = float64(1e10)
                    _ = v75
                    _ = v76
                }
                _ = fi73
            }
            v77 := GenS1{F0: (v2 % ((v2) | 1)), F1: ("" + "\n\t"), F2: 42}
            _ = fi72
            _ = v77
        }
        var v78 bool = (99 < uint8(v21))
        _ = v78
    } else {
        if (1 == 738) {
            switch 135 {
            case 9223372036854775807:
                v6 = v21
            }
            for fi79 := 0; fi79 < 16; fi79++ {
                _ = len(v1)
                _ = fi79
            }
            fmt.Println(v2, v21)
        }
        fmt.Println(v17)
        for fi80 := 0; fi80 < 12; fi80++ {
            var v81 float64 = float64(0.0)
            if (!(18446744073709551615 < 18446744073709551615)) {
                fmt.Println(v6, v15)
                for fi82 := 0; fi82 < 12; fi82++ {
                    v83 := []int{(fi80 | (fi80 >> ((255) & 63))), 4294967295, (fi82 << ((954) & 63))}
                    v83 = append(v83, fi82)
                    fmt.Println(v5, v15)
                    _ = fi82
                    _ = v83
                }
                var v84 int64 = 604
                _ = v84
            } else {
                var v85 int64 = -1
                _ = v85
            }
            _ = fi80
            _ = v81
        }
    }
    fmt.Println(v5)
    if false {
        var v86 bool = (!false)
        _ = v86
    } else {
        if true {
            v5 = v5
        }
        if (2147483647 > 343) {
            var v87 uint32 = (v17 * 177)
            _ = v87
        }
        fmt.Println(v2, v2, v5)
    }
    v17 = v5
    v88 := v3[128]
    v89 := GenS8{F0: v88, F1: 4294967295, F2: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
    _ = v89
}
