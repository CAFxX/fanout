package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int8; F1 int; F2 int64 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint64; F1 int64; F2 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 uint64; F2 int32 }
type GenS11 struct { F0 int32 }
func genFn0() bool {
    v5 := GenS1{F0: -1, F1: ((("" + ("foo bar" + "")) + ("foo bar" + "")) + "")}
    var v6 int64 = -728
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := []int{fi8}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if (!(true && (true && (!((1e10 + 1.5) > (float32(v7) + 0.1)))))) {
                var v12 int32 = 2147483647
                _ = v12
            }
            v13 := GenS2{F0: -128, F1: 0, F2: v6}
            _ = fi11
            _ = v13
        }
        _ = fi10
    }
    _ = v5
    _ = v6
    _ = v7
    return true
}
func genFn1() uint8 {
    var v14 float32 = 3.14159
    v14 += 1.5
    if ((!(false || (-1 <= 42))) && true) {
        if true {
            if ((255 != 2) && true) {
                var v15 uint16 = 511
                fmt.Println(v15, v15)
                _ = v15
            }
            switch "foo bar" {
            case "foo bar":
                v14 = ((v14 - ((3.14159 / v14) * v14)) * v14)
            }
            v14 = (0.0 + v14)
        } else {
            v16 := []int{-1, -488, -95}
            if (!true) {
                v17 := GenS1{F0: 0, F1: ""}
                var v18 float32 = (float32((0.1 / 1.5)) / float32((3.14159 + (1.5 + (3.14159 + -2.25)))))
                var v19 uint64 = 2
                _ = v17
                _ = v18
                _ = v19
            }
            _ = v16
        }
        v20 := []int{32768, -1, 32768, -430}
        v21 := GenS3{F0: -480, F1: 32768}
        _ = v20
        _ = v21
    }
    _ = v14
    return 127
}

func main() {
    var v1 int = 42
    var v2 int8 = 127
    var v3 string = "foo bar"
    v4 := make(map[uint32]string)
    v3 = (v3 + "hello")
    switch -421 {
    case -343:
        var v22 uint32 = 634
        _ = v22
    case 127:
        v23 := []int{v1, (v1 % (((-(v1))) | 1)), (^(v1)), -9223372036854775808}
        _ = v23
    case 746:
        if true {
            if true {
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 string = v3
                    if true {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            v1 = fi24
                            _ = fi26
                        }
                    }
                    v3 = ("foo bar" + v25)
                    _ = fi24
                    _ = v25
                }
                if false {
                    var v27 float64 = float64(1.5)
                    v28 := GenS4{F0: (v2 - (v2 % ((v2) | 1)))}
                    for fi29 := 0; fi29 < 12; fi29++ {
                        for fi30 := 0; fi30 < 7; fi30++ {
                            var v31 int16 = 1000
                            _ = fi30
                            _ = v31
                        }
                        fmt.Println(v3, v3)
                        v32 := []int{28, -31, 526, 127}
                        _ = fi29
                        _ = v32
                    }
                    _ = v27
                    _ = v28
                }
            } else {
                v1 -= v1
            }
            v4[2147483648] = "\n\t"
            switch v3 {
            case "\n\t":
                var v33 bool = (-2147483648 > 1)
                _ = v33
            }
        }
    default:
        v34 := GenS5{F0: v1, F1: 127}
        _ = v34
    }
    v2 = v2
    var v35 int8 = v2
    v36 := v4[2]
    switch (v1 ^ 128) {
    case 127:
        for fi37 := 0; fi37 < 18; fi37++ {
            for fi38 := 0; fi38 < 20; fi38++ {
                if (1 == 1000) {
                    switch fi38 {
                    case 328:
                        v39 := GenS5{F0: fi38, F1: (v2 % ((0) | 1))}
                        _ = v39
                    default:
                        switch ((("foo bar" + (v36 + (v3 + v36))) + "a") + v3) {
                        case "":
                            var v40 int8 = (v35 / ((v35) | 1))
                            _ = v40
                        case "a":
                            var v41 uint32 = 65535
                            _ = v41
                        }
                    }
                    v42 := []int{881, fi38, fi38}
                    if true {
                        v43 := GenS1{F0: -1, F1: v36}
                        switch v1 {
                        case 32768:
                            var v44 uint8 = 24
                            _ = v44
                        case 32767:
                            var v45 uint64 = 544
                            _ = v45
                        default:
                            fmt.Println(v36)
                        }
                        if false {
                            v46 := GenS6{F0: ("a" + v36)}
                            v47 := GenS5{F0: (fi37 ^ (v1 + (fi37 % ((fi37) | 1)))), F1: 100}
                            if true {
                                var v48 int8 = (v2 % ((v2) | 1))
                                v48 = v48
                                _ = v48
                            }
                            _ = v46
                            _ = v47
                        }
                        v49 := GenS7{F0: genFn1(), F1: 933}
                        v3 = v3
                        _ = v43
                        _ = v49
                    }
                    v2 -= (v35 << ((2) & 63))
                    if (true && true) {
                        fmt.Println(fi38)
                        var v50 uint64 = 18446744073709551615
                        var v51 float32 = 0.1
                        _ = v50
                        _ = v51
                    }
                    var v52 int8 = (v2 & 42)
                    _ = v42
                    _ = v52
                } else {
                    v53 := GenS8{F0: 105, F1: -9223372036854775808, F2: v35}
                    _ = v53
                }
                var v54 int32 = -2147483648
                v55 := GenS5{F0: (-(fi37)), F1: 127}
                _ = fi38
                _ = v54
                _ = v55
            }
            v56 := []int{v1, 0, (^(fi37))}
            _ = fi37
            _ = v56
        }
    default:
        var v57 string = v3
        _ = v57
    }
    if (((-2.25 + (3.14159 - float32(((0.0 / 1e10) * 3.14159)))) * 3.14159) != float32(-2.25)) {
        v3 = v3
        var v58 uint64 = 0
        v1 -= v1
        _ = v58
    } else {
        fmt.Println(v36)
        v59 := []int{127, v1, v1}
        v59 = append(v59, (v1 ^ (v1 >> ((579) & 63))))
        _ = v59
    }
    fmt.Println(v2)
    switch ("hello" + v3) {
    case "foo bar":
        fmt.Println(v36, v2)
    case "\n\t":
        v3 = v3
    default:
        var v60 bool = ((128 <= 2) && false)
        _ = v60
    }
    v3 = "a"
    v61 := v4[2]
    v62 := GenS6{F0: v36}
    if false {
        fmt.Println(v2, v61, v36)
    } else {
        switch v36 {
        case "a":
            v63 := GenS9{F0: -1}
            _ = v63
        case "hello":
            if (float32(1.5) <= 0.1) {
                var v64 string = v3
                v65 := GenS10{F0: 65535, F1: 1, F2: 42}
                _ = v64
                _ = v65
            } else {
                switch (v1 & -2147483648) {
                case -300:
                    v66 := v4[65535]
                    _ = v66
                }
                for fi67 := 0; fi67 < 5; fi67++ {
                    if true {
                        if false {
                            for fi68 := 0; fi68 < 10; fi68++ {
                                v3 = ((v61 + ("foo bar" + ((v3 + v3) + v36))) + "\n\t")
                                var v69 int = 32767
                                _ = fi68
                                _ = v69
                            }
                            if false {
                                var v70 bool = true
                                for fi71 := 0; fi71 < 14; fi71++ {
                                    var v72 int = fi71
                                    v72 = 9223372036854775807
                                    _ = fi71
                                    _ = v72
                                }
                                switch (v36 + v36) {
                                case "a":
                                    switch "\n\t" {
                                    case "a":
                                        v2 -= int8(v1)
                                    default:
                                        if (!true) {
                                            fmt.Println(v1, v35)
                                            v73 := GenS7{F0: 42, F1: 2147483647}
                                            v35 -= (v35 + (-(v2)))
                                            _ = v73
                                        }
                                    }
                                case "":
                                    v35 -= (v2 | (^(v35)))
                                case "\n\t":
                                    v74 := GenS11{F0: 2147483647}
                                    _ = v74
                                }
                                _ = v70
                            }
                        } else {
                            if (-963 >= -145) {
                                for fi75 := 0; fi75 < 17; fi75++ {
                                    fmt.Println(fi67, v1)
                                    var v76 bool = true
                                    var v77 uint32 = 0
                                    if (1e10 != (((1e10 * 1e10) * 0.0) + 1.5)) {
                                        v78 := GenS9{F0: 710}
                                        var v79 int64 = -1
                                        var v80 int16 = 1000
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                    }
                                    _ = fi75
                                    _ = v76
                                    _ = v77
                                }
                            }
                            fmt.Println(v2, v61, v2)
                        }
                        v81 := []int{v1, -809}
                        _ = len(v4)
                        _ = v81
                    }
                    var v82 string = "\n\t"
                    _ = fi67
                    _ = v82
                }
                var v83 float32 = 1.5
                _ = v83
            }
        case "\n\t":
            for fi84 := 0; fi84 < 5; fi84++ {
                switch v3 {
                case "a":
                    var v85 uint = 24
                    _ = v85
                case "":
                    var v86 float64 = ((3.14159 / 1e10) - 1e10)
                    _ = v86
                case "foo bar":
                    if false {
                        v87 := []int{v1, int(v35)}
                        var v88 bool = ((3.14159 * 0.0) >= (1.5 / 1.5))
                        _ = v87
                        _ = v88
                    }
                }
                switch 1 {
                case -140:
                    v1 += v1
                case 127:
                    for fi89 := 0; fi89 < 17; fi89++ {
                        var v90 string = ("" + "hello")
                        var v91 int32 = -1
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                case 4294967295:
                    fmt.Println(v35, v36)
                }
                v92 := []int{fi84}
                _ = fi84
                _ = v92
            }
        }
        switch v1 {
        case -2147483648:
            var v93 int64 = -9223372036854775808
            _ = v93
        default:
            for fi94 := 0; fi94 < 5; fi94++ {
                v2 = (-(v2))
                _ = fi94
            }
        }
    }
    if (((!false) || (!false)) && ((1 != 42) && false)) {
        v35 += v2
    } else {
        v2 += 63
        var v95 string = ((v3 + "") + v3)
        _ = v95
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v61
    _ = v62
}
