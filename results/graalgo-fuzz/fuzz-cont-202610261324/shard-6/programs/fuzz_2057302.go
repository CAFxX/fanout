package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS10 struct { F0 uint16 }

func main() {
    var v1 uint = 2
    var v2 string = ("hello" + "a")
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 string = ("a" + v2)
    for fi6 := 0; fi6 < 6; fi6++ {
        switch fi6 {
        case 1:
            v2 = "a"
        case 42:
            if true {
                v5 = v2
            } else {
                v7 := GenS1{F0: 677, F1: 0, F2: 2147483647}
                _ = v7
            }
        default:
            v2 = v5
        }
        switch 2147483647 {
        case -1:
            v8 := GenS2{F0: 111, F1: (fi6 | fi6)}
            _ = v8
        case 32767:
            if true {
                var v9 string = v5
                _ = v9
            }
        default:
            v10 := []int{32768, (fi6 % ((fi6) | 1))}
            _ = v10
        }
        _ = fi6
    }
    if (-733 > -950) {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 65535
            _ = fi11
            _ = v12
        }
        fmt.Println(v1, v5)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS2{F0: -77, F1: fi13}
        fmt.Println(v1, v5, v5)
        if (false && true) {
            v15 := []int{fi13, fi13, 2147483647}
            if true {
                var v16 float64 = 1e10
                _ = v16
            }
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v17 uint = (v1 | (^(v1)))
    if (true || (1e10 == 1.5)) {
        fmt.Println(v17, v2)
        var v18 uint32 = 1
        _ = v18
    }
    var v19 int64 = 42
    var v20 uint32 = 2147483648
    fmt.Println(v1)
    v21 := []int{-53, -670}
    v22 := GenS2{F0: 0, F1: -787}
    if true {
        v23 := GenS3{F0: 0, F1: 0}
        var v24 uint16 = 846
        v24 = v24
        _ = v23
        _ = v24
    } else {
        v21 = append(v21, -9223372036854775808)
        v25 := GenS3{F0: 57, F1: 100}
        _ = v25
    }
    _ = len(v21)
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 uint64 = 18446744073709551615
        _ = fi26
        _ = v27
    }
    v19 = v19
    for fi28 := 0; fi28 < 11; fi28++ {
        if true {
            for fi29 := 0; fi29 < 11; fi29++ {
                if true {
                    v30 := GenS4{F0: 0, F1: 128, F2: (v1 % ((736) | 1))}
                    var v31 int = -1
                    if false {
                        v32 := GenS5{F0: 1, F1: -813}
                        _ = v32
                    } else {
                        if false {
                            var v33 int32 = 2147483647
                            _ = v33
                        }
                        v2 = (v2 + ((v2 + v5) + v2))
                        v34 := GenS5{F0: v20, F1: -1}
                        if true {
                            if true {
                                if (1.5 == 0.0) {
                                    if (0 != 527) {
                                        var v35 bool = true
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                        v5 = "hello"
                                        _ = v35
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v36 := v4[(^(fi28))]
                                        _ = v36
                                    case "\n\t":
                                        v20 -= 65535
                                    case "a":
                                        v2 = v5
                                    default:
                                        if (!false) {
                                            v37 := GenS3{F0: 255, F1: -56}
                                            _ = v37
                                        }
                                    }
                                    switch (fi29 * fi29) {
                                    case 42:
                                        if false {
                                            var v38 uint8 = 117
                                            v39 := GenS6{F0: (0 < 42)}
                                            v2 = v5
                                            _ = v38
                                            _ = v39
                                        } else {
                                            if (true || false) {
                                                v1 = v17
                                                var v40 uint32 = (v20 - v20)
                                                var v41 uint16 = 2
                                                _ = v40
                                                _ = v41
                                            }
                                            fmt.Println(v19)
                                            if true {
                                                delete(v3, 128)
                                                var v42 int = fi29
                                                _ = v42
                                            }
                                            v4[(fi28 & (fi28 % ((v31) | 1)))] = (v20 - v20)
                                            var v43 uint = 2
                                            var v44 int16 = 1
                                            v31 += -1
                                            switch v5 {
                                            case "":
                                                v17 -= v17
                                            case "\n\t":
                                                var v45 string = "foo bar"
                                                _ = v45
                                            case "foo bar":
                                                _ = len(v4)
                                            default:
                                                v46 := GenS7{F0: -128, F1: 0, F2: 42}
                                                switch v2 {
                                                case "hello":
                                                    var v47 int = fi28
                                                    _ = v47
                                                default:
                                                    switch (fi29 & fi28) {
                                                    case -39:
                                                        if len(v21) > 0 {
                                                            v21[0] = -9223372036854775808
                                                        }
                                                    case 2147483647:
                                                        v48 := GenS4{F0: -9223372036854775808, F1: 2, F2: v1}
                                                        _ = v48
                                                    case 127:
                                                        var v49 int16 = v44
                                                        _ = v49
                                                    default:
                                                        fmt.Println(v5)
                                                    }
                                                }
                                                _ = v46
                                            }
                                            _ = v43
                                            _ = v44
                                        }
                                    default:
                                        switch (v31 * v31) {
                                        case 127:
                                            v50 := GenS6{F0: (32767 != 570)}
                                            _ = v50
                                        }
                                    }
                                    if (!(-128 == -1)) {
                                        v2 = v2
                                        if true {
                                            var v51 int8 = -1
                                            _ = v51
                                        }
                                        if false {
                                            if len(v21) > 0 {
                                                v21[0] = fi29
                                            }
                                            v52 := GenS7{F0: -1, F1: (v1 ^ v1), F2: 18446744073709551615}
                                            v17 = 1
                                            v5 = v5
                                            _ = len(v4)
                                            var v53 uint32 = (^(v20))
                                            _ = v52
                                            _ = v53
                                        } else {
                                            if (1 >= v19) {
                                                var v54 int8 = 127
                                                var v55 int32 = 0
                                                if false {
                                                    _ = len(v21)
                                                    fmt.Println(fi28)
                                                }
                                                var v56 float64 = 3.14159
                                                if false {
                                                    v17 += v1
                                                    switch (^(v31)) {
                                                    case 128:
                                                        v57 := GenS8{F0: 1, F1: v5}
                                                        _ = v57
                                                    case 4294967295:
                                                        var v58 int = 0
                                                        _ = v58
                                                    default:
                                                        v56 = v56
                                                    }
                                                } else {
                                                    var v59 string = v5
                                                    v20 += v20
                                                    if len(v21) > 0 {
                                                        v21[0] = (fi28 & -566)
                                                    }
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                    _ = v59
                                                }
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            }
                                            if (-9223372036854775808 >= (v19 + v19)) {
                                                var v60 int64 = (v19 - -128)
                                                if (65535 == v1) {
                                                    switch (v2 + v2) {
                                                    case "\n\t":
                                                        var v61 bool = false
                                                        _ = v61
                                                    case "":
                                                        var v62 uint8 = 60
                                                        _ = v62
                                                    case "foo bar":
                                                        fmt.Println(v31, v31, fi29)
                                                    }
                                                    var v63 string = v5
                                                    _ = v63
                                                }
                                                var v64 string = v2
                                                if (2147483647 <= 313) {
                                                    fmt.Println(v5, fi29)
                                                } else {
                                                    fmt.Println(v5)
                                                    if (-2.25 < 0.1) {
                                                        delete(v4, (fi28 | (fi29 << ((v1) & 63))))
                                                        v20 = v20
                                                    } else {
                                                        switch (fi29 << (((v1 - 491)) & 63)) {
                                                        case 4294967295:
                                                            var v65 float32 = (-2.25 / (0.1 / 1.5))
                                                            _ = v65
                                                        case 127:
                                                            var v66 int = 4294967295
                                                            _ = v66
                                                        }
                                                        if len(v21) > 0 {
                                                            v21[0] = -298
                                                        }
                                                    }
                                                }
                                                v67 := GenS6{F0: (false && true)}
                                                _ = v60
                                                _ = v64
                                                _ = v67
                                            }
                                            v68 := GenS1{F0: 18446744073709551615, F1: 127, F2: fi29}
                                            _ = v68
                                        }
                                    }
                                } else {
                                    fmt.Println(v1, fi29, v20)
                                    if false {
                                        var v69 int32 = 566
                                        _ = v69
                                    }
                                    fmt.Println(v2, v1)
                                }
                                if len(v21) > 0 {
                                    v21[0] = 326
                                }
                            } else {
                                fmt.Println(v19, v17, v20)
                            }
                            var v70 int64 = (-(v19))
                            v71 := v4[(fi29 ^ (fi28 << ((v1) & 63)))]
                            _ = v70
                            _ = v71
                        }
                        var v72 int16 = -690
                        _ = v34
                        _ = v72
                    }
                    _ = v30
                    _ = v31
                } else {
                    v17 = v1
                    fmt.Println(v1)
                    var v73 uint32 = v20
                    _ = v73
                }
                _ = fi29
            }
            v1 = (v1 << (((v1 % ((v17) | 1))) & 63))
        } else {
            for fi74 := 0; fi74 < 14; fi74++ {
                var v75 float32 = 0.1
                _ = fi74
                _ = v75
            }
            switch (^(fi28)) {
            case 758:
                var v76 float64 = ((1e10 + (float64(0.1) + ((1.5 * ((float64(((float32(-2.25) - -2.25) * 3.14159)) * 0.0) * 0.0)) - 3.14159))) + 0.1)
                _ = v76
            case 937:
                v2 = ((("hello" + v2) + v2) + v5)
            case 42:
                if (3.14159 > (0.1 - (1.5 + 0.1))) {
                    for fi77 := 0; fi77 < 17; fi77++ {
                        var v78 uint8 = 127
                        fmt.Println(fi28, fi77, v1)
                        _ = fi77
                        _ = v78
                    }
                } else {
                    fmt.Println(v2, fi28)
                }
            }
            for fi79 := 0; fi79 < 17; fi79++ {
                v80 := GenS9{F0: v20, F1: 992, F2: 32767}
                var v81 uint8 = 2
                _ = fi79
                _ = v80
                _ = v81
            }
        }
        v20 = v20
        var v82 uint8 = 127
        _ = fi28
        _ = v82
    }
    var v83 int8 = -128
    if (0.1 == 0.1) {
        switch ((("\n\t" + "foo bar") + ("hello" + ("a" + v2))) + v2) {
        case "\n\t":
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        case "hello":
            if (v83 == v83) {
                var v84 int32 = -263
                _ = v84
            } else {
                v85 := GenS10{F0: 908}
                var v86 float32 = ((1.5 * 3.14159) - 1e10)
                _ = v85
                _ = v86
            }
        case "foo bar":
            for fi87 := 0; fi87 < 14; fi87++ {
                v83 += v83
                var v88 uint = (v17 << ((v1) & 63))
                v88 = (v1 % ((v88) | 1))
                _ = fi87
                _ = v88
            }
        }
    }
    var v89 uint32 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v89
}
