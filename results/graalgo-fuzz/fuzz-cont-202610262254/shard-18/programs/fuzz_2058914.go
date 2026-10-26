package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint }
type GenS6 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0(p0 bool, p1 uint) uint {
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    if ("\n\t" > "foo bar") {
        var v5 int32 = 42
        _ = v5
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1() uint16 {
    var v6 int32 = 1
    if ((42 >= 148) || false) {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: 2, F1: 1}
            _ = fi7
            _ = v8
        }
    } else {
        var v9 uint32 = 81
        for fi10 := 0; fi10 < 4; fi10++ {
            switch 4294967295 {
            case -9223372036854775808:
                v9 += 65535
            }
            v11 := GenS2{F0: int8(fi10)}
            fmt.Println(fi10, v6)
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    var v12 uint = 255
    if (v12 == (v12 * v12)) {
        var v13 int64 = 610
        v13 = (v13 - (^(v13)))
        _ = v13
    }
    switch 237 {
    case 42:
        var v14 int32 = (-(v6))
        _ = v14
    case -9223372036854775808:
        switch 128 {
        case 32767:
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 15; fi16++ {
                    if (!(1 < 170)) {
                        v12 = (v12 & (v12 - v12))
                        v17 := GenS1{F0: 324, F1: 1}
                        _ = v17
                    }
                    _ = fi16
                }
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 int8 = -1
                    _ = fi18
                    _ = v19
                }
                var v20 uint16 = 1
                _ = fi15
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 uint16 = 765
                _ = fi21
                _ = v22
            }
        }
    case 393:
        if (v6 < v6) {
            v23 := []int{-2147483648, -458, -1}
            v12 -= (v12 * 4294967295)
            if true {
                var v24 uint16 = 894
                switch ("" + "hello") {
                case "":
                    v25 := GenS1{F0: 477, F1: 2}
                    _ = v25
                }
                _ = v24
            } else {
                _ = len(v23)
                v12 += (v12 | 18446744073709551615)
            }
            _ = v23
        }
    }
    _ = v12
    _ = v6
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 491
    var v3 uint64 = 42
    var v26 float32 = 1e10
    v27 := GenS1{F0: (v3 * 42), F1: 209}
    if (!true) {
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(v26, fi28)
            if ((v2 - (v2 ^ (v2 + 255))) != 255) {
                for fi29 := 0; fi29 < 5; fi29++ {
                    var v30 uint = genFn0(true, 18446744073709551615)
                    if false {
                        v31 := []int{-2147483648, (^(fi28)), (fi29 - fi28), 579}
                        _ = v31
                    }
                    v30 = uint(v3)
                    _ = fi29
                    _ = v30
                }
                if false {
                    v32 := []int{0}
                    if false {
                        if ((0 <= 69) || (((2147483647 >= 65535) && (v3 != v3)) && false)) {
                            var v33 float64 = ((-2.25 + 1e10) - 0.1)
                            fmt.Println(v26)
                            var v34 float64 = (0.1 + 0.0)
                            _ = v33
                            _ = v34
                        } else {
                            switch int(v2) {
                            case -9223372036854775808:
                                v35 := GenS2{F0: 0}
                                _ = v35
                            case -990:
                                var v36 float64 = (float64((v26 + v26)) - (float64(v26) - 3.14159))
                                _ = v36
                            case 32767:
                                v37 := GenS3{F0: 903, F1: -1}
                                _ = v37
                            }
                        }
                    } else {
                        v38 := GenS4{F0: -521}
                        switch (fi28 * (^(fi28))) {
                        case 128:
                            v39 := GenS4{F0: 0}
                            _ = v39
                        case 9223372036854775807:
                            if (1 <= 42) {
                                for fi40 := 0; fi40 < 16; fi40++ {
                                    if ((v26 < -2.25) || false) {
                                        v2 = (v2 & (v2 + (v2 & v2)))
                                        var v41 uint8 = 2
                                        _ = v41
                                    }
                                    v1 += 787
                                    _ = fi40
                                }
                                v1 -= 42
                            }
                        }
                        _ = v38
                    }
                    _ = v32
                } else {
                    fmt.Println(fi28, v1, v2)
                    var v42 int8 = -128
                    _ = v42
                }
                fmt.Println(v26, v26)
            }
            _ = fi28
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            var v44 float32 = (v26 / 0.1)
            v45 := GenS2{F0: 87}
            v26 += float32(0.1)
            _ = fi43
            _ = v44
            _ = v45
        }
    } else {
        fmt.Println(v2, v1)
        v26 = float32((-2.25 / 0.1))
    }
    v46 := GenS2{F0: 82}
    if (-35 != 1) {
        if false {
            var v47 uint64 = (^(v3))
            _ = v47
        } else {
            var v48 float64 = 0.1
            _ = v48
        }
        v3 = (v3 + 2)
        v49 := GenS3{F0: 270, F1: 0}
        _ = v49
    } else {
        var v50 int = 543
        v51 := []int{v50, 0, (v50 & (v50 - v50)), int(v1)}
        for fi52 := 0; fi52 < 1; fi52++ {
            for fi53 := 0; fi53 < 17; fi53++ {
                var v54 int16 = int16(v2)
                fmt.Println(v2, v2)
                _ = fi53
                _ = v54
            }
            _ = fi52
        }
        _ = v50
        _ = v51
    }
    v55 := []int{}
    switch "" {
    case "foo bar":
        if len(v55) > 0 {
            v55[0] = -1
        }
    case "":
        v56 := GenS5{F0: 1000, F1: 187, F2: genFn0(true, 717)}
        _ = v56
    default:
        fmt.Println(v2)
    }
    v57 := GenS1{F0: (v3 ^ v3), F1: 141}
    if len(v55) > 0 {
        v55[0] = 0
    }
    switch "\n\t" {
    case "foo bar":
        v26 = (v26 * 0.1)
    case "":
        switch ("hello" + "") {
        case "\n\t":
            if (-1 != 9223372036854775807) {
                switch "hello" {
                case "a":
                    if false {
                        if (true && true) {
                            var v58 uint16 = 268
                            for fi59 := 0; fi59 < 5; fi59++ {
                                var v60 int8 = -1
                                fmt.Println(v26, v2, v58)
                                for fi61 := 0; fi61 < 16; fi61++ {
                                    v26 -= 0.0
                                    _ = fi61
                                }
                                _ = fi59
                                _ = v60
                            }
                            _ = v58
                        } else {
                            for fi62 := 0; fi62 < 5; fi62++ {
                                fmt.Println(v1)
                                var v63 float32 = v26
                                _ = fi62
                                _ = v63
                            }
                            var v64 bool = (v26 <= (((v26 + (1e10 - float32((3.14159 / 3.14159)))) + 1.5) + v26))
                            v65 := GenS2{F0: 17}
                            _ = v64
                            _ = v65
                        }
                        switch "\n\t" {
                        case "foo bar":
                            var v66 int64 = 9223372036854775807
                            _ = v66
                        case "":
                            switch -236 {
                            case 32768:
                                var v67 int32 = -427
                                _ = v67
                            case 128:
                                switch 1 {
                                case -342:
                                    var v68 bool = true
                                    _ = v68
                                case 9223372036854775807:
                                    if (-771 != 42) {
                                        for fi69 := 0; fi69 < 9; fi69++ {
                                            v70 := GenS4{F0: 1}
                                            var v71 float32 = 3.14159
                                            _ = fi69
                                            _ = v70
                                            _ = v71
                                        }
                                        v72 := GenS6{F0: (v1 & v1), F1: 13, F2: -9223372036854775808}
                                        _ = len(v55)
                                        _ = v72
                                    }
                                }
                            default:
                                var v73 float32 = (v26 / float32(0.1))
                                _ = v73
                            }
                        default:
                            var v74 bool = true
                            _ = v74
                        }
                    } else {
                        v75 := GenS1{F0: v3, F1: uint8(v2)}
                        var v76 string = "hello"
                        _ = v75
                        _ = v76
                    }
                case "\n\t":
                    _ = len(v55)
                case "":
                    if false {
                        _ = len(v55)
                    }
                }
            }
        default:
            v26 = 1e10
        }
    case "a":
        v1 = v1
    default:
        if false {
            for fi77 := 0; fi77 < 14; fi77++ {
                var v78 int64 = int64(fi77)
                for fi79 := 0; fi79 < 7; fi79++ {
                    if false {
                        fmt.Println(fi79)
                        fmt.Println(v1)
                        if ((v1 + (v1 / (((v1 << ((1) & 63))) | 1))) != v1) {
                            v55 = append(v55, fi77)
                            v1 += v1
                            fmt.Println(v2)
                            var v80 int64 = (-(v78))
                            _ = v80
                        } else {
                            switch 241 {
                            case 1:
                                v81 := GenS6{F0: v1, F1: 255, F2: -1}
                                _ = v81
                            }
                        }
                    } else {
                        if (255 == 0) {
                            switch "a" {
                            case "foo bar":
                                if len(v55) >= 2 {
                                    v55 = v55[1:len(v55)]
                                }
                            default:
                                fmt.Println(v3, v26)
                                var v82 uint32 = 2147483647
                                var v83 uint8 = 127
                                _ = v82
                                _ = v83
                            }
                        }
                        _ = len(v55)
                        fmt.Println(v3, v3)
                    }
                    _ = len(v55)
                    var v84 uint32 = 271
                    _ = fi79
                    _ = v84
                }
                _ = fi77
                _ = v78
            }
            if len(v55) >= 2 {
                v55 = v55[1:len(v55)]
            }
        } else {
            var v85 int16 = -950
            switch -9223372036854775808 {
            case 32768:
                v86 := GenS2{F0: 1}
                _ = v86
            case 128:
                var v87 float64 = float64(v26)
                _ = v87
            case 552:
                var v88 uint8 = 0
                _ = v88
            default:
                v89 := GenS6{F0: v1, F1: 127, F2: 9223372036854775807}
                _ = v89
            }
            for fi90 := 0; fi90 < 16; fi90++ {
                if false {
                    v2 -= v2
                } else {
                    v55 = append(v55, (fi90 + (fi90 / ((fi90) | 1))))
                    v55 = append(v55, fi90)
                    for fi91 := 0; fi91 < 1; fi91++ {
                        switch (fi90 / (((fi90 >> ((65535) & 63))) | 1)) {
                        case -816:
                            switch "a" {
                            case "a":
                                if false {
                                    var v92 uint8 = 1
                                    var v93 uint8 = 0
                                    _ = v92
                                    _ = v93
                                }
                            case "hello":
                                fmt.Println(v1, v85)
                            }
                        case 1:
                            switch "hello" {
                            case "a":
                                if len(v55) > 0 {
                                    v55[0] = (fi91 / ((fi90) | 1))
                                }
                            default:
                                fmt.Println(v1)
                            }
                        default:
                            fmt.Println(fi90, v3)
                        }
                        _ = fi91
                    }
                }
                _ = fi90
            }
            _ = v85
        }
    }
    v3 = 0
    var v94 int16 = 173
    var v95 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v46
    _ = v55
    _ = v57
    _ = v94
    _ = v95
}
