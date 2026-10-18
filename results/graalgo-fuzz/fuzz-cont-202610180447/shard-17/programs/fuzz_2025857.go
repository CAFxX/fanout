package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int; F1 uint; F2 int64 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint; F1 uint16; F2 int }
type GenS9 struct { F0 int64; F1 string; F2 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64; F1 uint32 }
func genFn0() uint32 {
    var v7 uint8 = 255
    switch "foo bar" {
    case "foo bar":
        var v8 string = "hello"
        _ = v8
    default:
        var v9 uint8 = (v7 - (v7 & 2))
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: -27, F2: 32768}
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint = 207
        var v13 int = (fi11 % (((^(fi11))) | 1))
        fmt.Println(v13, fi11, v13)
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v7)
    _ = v10
    _ = v7
    return 0
}
func genFn1() uint32 {
    if false {
        var v14 uint64 = 0
        switch "foo bar" {
        case "\n\t":
            v14 = v14
        case "foo bar":
            switch 127 {
            case 32768:
                for fi15 := 0; fi15 < 18; fi15++ {
                    v14 += (v14 - 776)
                    _ = fi15
                }
            case 4294967295:
                v16 := GenS1{F0: 0, F1: 127, F2: 42}
                _ = v16
            case 127:
                fmt.Println(v14)
            }
        }
        if (int(v14) >= -9223372036854775808) {
            v14 = (v14 >> ((597) & 63))
            v17 := []int{35, 1, 9223372036854775807, -9223372036854775808}
            _ = v17
        } else {
            if false {
                switch ("foo bar" + (("hello" + "foo bar") + "")) {
                case "":
                    var v18 float64 = 1.5
                    _ = v18
                case "\n\t":
                    var v19 int64 = -9223372036854775808
                    _ = v19
                }
                v20 := GenS1{F0: -32768, F1: 127, F2: -73}
                _ = v20
            } else {
                var v21 uint32 = 1
                v22 := GenS1{F0: -32768, F1: 42, F2: -2147483648}
                var v23 uint8 = 255
                _ = v21
                _ = v22
                _ = v23
            }
        }
        _ = v14
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 15; fi25++ {
            if (0.1 == float64((0.0 - (1e10 * 1e10)))) {
                if false {
                    v26 := GenS2{F0: 708, F1: 32767}
                    var v27 uint16 = 341
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS3{F0: 0}
                    _ = v28
                }
            }
            var v29 int64 = 397
            _ = fi25
            _ = v29
        }
        _ = fi24
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        if false {
            if false {
                v31 := GenS4{F0: (fi30 % (((fi30 * fi30)) | 1))}
                switch 1 {
                case -383:
                    switch (("hello" + "hello") + "a") {
                    case "\n\t":
                        for fi32 := 0; fi32 < 7; fi32++ {
                            var v33 float32 = (0.0 + 0.0)
                            var v34 float32 = (v33 - v33)
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    case "":
                        if false {
                            var v35 float32 = (3.14159 * (1e10 + 0.0))
                            _ = v35
                        } else {
                            var v36 int32 = 42
                            v37 := GenS5{F0: 265, F1: 2147483647}
                            _ = v36
                            _ = v37
                        }
                    }
                case 514:
                    v38 := []int{(fi30 % (((fi30 << ((0) & 63))) | 1)), fi30}
                    _ = v38
                case -9223372036854775808:
                    var v39 int16 = 0
                    _ = v39
                default:
                    if (3.14159 <= float32(-2.25)) {
                        v40 := []int{(fi30 << ((1) & 63)), fi30}
                        var v41 bool = false
                        for fi42 := 0; fi42 < 15; fi42++ {
                            v41 = (0 >= -504)
                            var v43 int = fi42
                            if (0 == 9223372036854775807) {
                                switch "hello" {
                                case "hello":
                                    switch fi42 {
                                    case 4294967295:
                                        switch ((("hello" + "a") + ("\n\t" + "")) + ("\n\t" + ("foo bar" + "foo bar"))) {
                                        case "a":
                                            var v44 uint = 500
                                            _ = v44
                                        case "\n\t":
                                            v45 := GenS5{F0: 9223372036854775807, F1: 42}
                                            _ = v45
                                        case "hello":
                                            if (711 >= 1) {
                                                var v46 int64 = 42
                                                _ = v46
                                            } else {
                                                var v47 int8 = 42
                                                var v48 uint32 = 1
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                    case -216:
                                        v43 = v43
                                    default:
                                        var v49 uint32 = genFn0()
                                        _ = v49
                                    }
                                case "foo bar":
                                    switch "" {
                                    case "\n\t":
                                        v50 := GenS6{F0: -1, F1: 690, F2: 1}
                                        _ = v50
                                    }
                                case "":
                                    if v41 {
                                        v43 = fi42
                                    } else {
                                        v51 := GenS7{F0: 139, F1: 4294967295}
                                        fmt.Println(v43, fi30)
                                        _ = v51
                                    }
                                }
                            }
                            v41 = false
                            _ = fi42
                            _ = v43
                        }
                        _ = v40
                        _ = v41
                    } else {
                        switch 9223372036854775807 {
                        case 1:
                            for fi52 := 0; fi52 < 13; fi52++ {
                                v53 := []int{fi52, 127, fi52}
                                v54 := GenS3{F0: 1}
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                        case 42:
                            for fi55 := 0; fi55 < 7; fi55++ {
                                fmt.Println(fi30, fi55)
                                var v56 string = "hello"
                                if false {
                                    v57 := GenS7{F0: 2, F1: 18446744073709551615}
                                    _ = v57
                                } else {
                                    v58 := GenS8{F0: 1, F1: 185, F2: 222}
                                    _ = v58
                                }
                                _ = fi55
                                _ = v56
                            }
                        case 32767:
                            var v59 uint64 = 831
                            _ = v59
                        }
                        var v60 uint64 = 274
                        _ = v60
                    }
                }
                _ = v31
            }
            var v61 int = fi30
            _ = v61
        }
        var v62 uint64 = 42
        _ = fi30
        _ = v62
    }
    v63 := []int{1, 32768, 9223372036854775807, 1}
    var v64 uint64 = 756
    _ = v63
    _ = v64
    return genFn0()
}

func main() {
    var v1 uint = 442
    var v2 float32 = 1e10
    v3 := []uint32{2147483648, 0, 65535}
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    var v6 float32 = (v2 + v2)
    fmt.Println(v1, v6, v6)
    v65 := GenS9{F0: 9223372036854775807, F1: ("hello" + "hello"), F2: 2}
    var v66 int = 634
    switch ("hello" + ("foo bar" + "a")) {
    case "a":
        v3 = append(v3, 2147483647)
    case "\n\t":
        v2 -= 3.14159
    }
    if (-1 >= 0) {
        v67 := GenS2{F0: 315, F1: 1}
        _ = v67
    }
    if true {
        fmt.Println(v6)
        var v68 int = v66
        _ = v68
    }
    for fi69 := 0; fi69 < 11; fi69++ {
        v2 = 0.1
        _ = fi69
    }
    var v70 int = -688
    v6 += v2
    v3 = append(v3, 913)
    var v71 uint64 = 2
    for fi72 := 0; fi72 < 6; fi72++ {
        var v73 string = "\n\t"
        if (!(true || false)) {
            if true {
                if false {
                    fmt.Println(v71, v66)
                }
            }
            v71 = (v71 / (((v71 % ((v71) | 1))) | 1))
            var v74 uint = v1
            _ = v74
        }
        var v75 float32 = v6
        _ = fi72
        _ = v73
        _ = v75
    }
    var v76 string = "hello"
    if true {
        for fi77 := 0; fi77 < 3; fi77++ {
            if ((v71 >> ((v1) & 63)) != v71) {
                v76 = "\n\t"
            } else {
                v3 = append(v3, 0)
            }
            v78 := GenS10{F0: 2147483647}
            _ = fi77
            _ = v78
        }
    } else {
        switch v70 {
        case 9223372036854775807:
            v79 := GenS11{F0: v71, F1: 970}
            _ = v79
        case 0:
            v80 := GenS7{F0: 255, F1: v71}
            _ = v80
        case 32767:
            for fi81 := 0; fi81 < 6; fi81++ {
                var v82 uint32 = 966
                _ = fi81
                _ = v82
            }
        default:
            v71 -= (v71 ^ v71)
        }
        v3 = append(v3, 75)
    }
    delete(v5, (v66 * (^(v70))))
    v66 -= (v66 >> (((v1 - (v1 / ((v1) | 1)))) & 63))
    if (3.14159 != (v6 - 1.5)) {
        var v83 uint8 = 128
        _ = v83
    } else {
        fmt.Println(v70, v66)
        v6 += 1.5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v70
    _ = v71
    _ = v76
}
