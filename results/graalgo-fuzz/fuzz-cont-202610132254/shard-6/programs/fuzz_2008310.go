package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 uint; F2 int16 }
type GenS4 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 string; F1 bool; F2 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8 }

func main() {
    v1 := []int{103}
    v2 := []uint32{890, 42, 1}
    var v3 int64 = 66
    var v4 int = 128
    var v5 string = ("" + "\n\t")
    fmt.Println(v4)
    v6 := GenS1{F0: -71, F1: 920, F2: 486}
    if len(v1) > 0 {
        v1[0] = v4
    }
    var v7 string = v5
    var v8 int16 = 1
    switch v4 {
    case 32768:
        v9 := GenS1{F0: -128, F1: (-(v8)), F2: 568}
        _ = v9
    case -9223372036854775808:
        var v10 uint = 0
        _ = v10
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            switch (v4 - fi11) {
            case 1:
                if true {
                    v2 = append(v2, 2147483647)
                    v5 = v5
                }
            default:
                v12 := GenS1{F0: -128, F1: 32767, F2: 145}
                _ = v12
            }
            _ = fi11
        }
    }
    if (18446744073709551615 != 65535) {
        v4 += (-(v4))
    } else {
        switch "a" {
        case "\n\t":
            for fi13 := 0; fi13 < 20; fi13++ {
                _ = len(v1)
                var v14 bool = ((-285 < -430) && ((^(fi13)) <= 27))
                _ = fi13
                _ = v14
            }
        default:
            var v15 bool = (!((0 < 0) && (42 == 111)))
            _ = v15
        }
    }
    v16 := GenS2{F0: 418, F1: 49, F2: -344}
    var v17 string = "foo bar"
    v4 = v4
    fmt.Println(v5, v3)
    var v18 uint = 0
    if false {
        v19 := GenS2{F0: v18, F1: -91, F2: (v3 % (((v3 % (((v3 | 399)) | 1))) | 1))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v19
    } else {
        fmt.Println(v4)
    }
    v7 = ("a" + (v17 + "a"))
    if (!((false || (!true)) || (false || (!(2147483648 != 338))))) {
        v7 = "a"
        for fi20 := 0; fi20 < 13; fi20++ {
            if (v18 < (v18 * 0)) {
                v17 = v7
                v21 := GenS3{F0: (v18 * v18), F1: 42, F2: 1000}
                _ = v21
            } else {
                if ((2147483647 == 2147483648) || true) {
                    v3 = v3
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v4 -= (v4 >> (((v18 << (((v18 * v18)) & 63))) & 63))
                        _ = fi22
                    }
                    v4 = fi20
                } else {
                    if ((2 > 349) && true) {
                        v8 = (v8 - v8)
                        if false {
                            var v23 int32 = 2147483647
                            var v24 uint16 = 534
                            fmt.Println(v23, v4, v4)
                            _ = v23
                            _ = v24
                        } else {
                            if (!(2147483647 < 1)) {
                                if true {
                                    v25 := GenS1{F0: -87, F1: v8, F2: 2}
                                    _ = v25
                                } else {
                                    var v26 uint16 = 42
                                    var v27 float64 = 1.5
                                    v2 = append(v2, 42)
                                    _ = v26
                                    _ = v27
                                }
                                for fi28 := 0; fi28 < 15; fi28++ {
                                    var v29 float32 = (0.1 / 1.0)
                                    v17 = (("" + v17) + (("\n\t" + ("" + v17)) + ((v17 + "a") + (v5 + ("foo bar" + v17)))))
                                    _ = fi28
                                    _ = v29
                                }
                            } else {
                                v1 = append(v1, (v4 - v4))
                            }
                            if (!false) {
                                var v30 int32 = 377
                                _ = v30
                            } else {
                                if (3.14159 > float64(3.14159)) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v8)
                                    var v31 int64 = (-(v3))
                                    _ = v31
                                }
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    if (!(!true)) {
                                        switch (^(fi20)) {
                                        case 32768:
                                            fmt.Println(v5, v4)
                                        }
                                        v33 := GenS2{F0: (v18 - v18), F1: 42, F2: (v3 ^ v3)}
                                        var v34 int64 = (^(v3))
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v35 := GenS4{F0: (^(v3)), F1: 4294967295, F2: 657}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v17 = v5
                                        v36 := GenS5{F0: 2, F1: 18446744073709551615, F2: 108}
                                        _ = v35
                                        _ = v36
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v3 = v3
                                    _ = fi32
                                }
                                switch (v4 & (^(fi20))) {
                                case -51:
                                    for fi37 := 0; fi37 < 8; fi37++ {
                                        v5 = "foo bar"
                                        v5 = "hello"
                                        _ = fi37
                                    }
                                case 32767:
                                    for fi38 := 0; fi38 < 1; fi38++ {
                                        v2 = append(v2, 841)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if false {
                                            var v39 uint64 = 875
                                            v40 := GenS3{F0: v18, F1: (v18 ^ (v18 >> ((255) & 63))), F2: v8}
                                            _ = v39
                                            _ = v40
                                        } else {
                                            if (337 == -107) {
                                                v41 := GenS6{F0: fi20, F1: false}
                                                var v42 uint8 = 42
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                        _ = fi38
                                    }
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                }
                for fi43 := 0; fi43 < 18; fi43++ {
                    var v44 float64 = 0.0
                    fmt.Println(fi43, v18, fi43)
                    var v45 string = "a"
                    v1 = append(v1, fi20)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                var v46 uint8 = 128
                _ = v46
            }
            _ = fi20
        }
    } else {
        v47 := GenS3{F0: v18, F1: v18, F2: (v8 - v8)}
        switch 2147483647 {
        case 32767:
            var v48 uint64 = 0
            _ = v48
        case -2147483648:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 389:
            if (v5 <= v7) {
                switch "hello" {
                case "a":
                    var v49 int16 = (^(v8))
                    _ = v49
                case "":
                    if (!true) {
                        v4 += v4
                        var v50 uint16 = 134
                        fmt.Println(v18, v7)
                        _ = v50
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = v4
                    }
                }
            }
        default:
            if true {
                var v51 uint = v18
                _ = v51
            }
        }
        _ = v47
    }
    switch (v4 - (v4 | v4)) {
    case 626:
        switch (v4 / (((v4 >> ((v18) & 63))) | 1)) {
        case -1:
            if ((1 <= -1) || (100 != 127)) {
                for fi52 := 0; fi52 < 8; fi52++ {
                    v7 = v7
                    fmt.Println(v8, v3)
                    _ = fi52
                }
                var v53 uint8 = 127
                v3 += v3
                _ = v53
            }
        case 38:
            for fi54 := 0; fi54 < 3; fi54++ {
                v1 = append(v1, (-(v4)))
                switch v17 {
                case "hello":
                    fmt.Println(v17, v4)
                case "\n\t":
                    v5 = v7
                default:
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                }
                v55 := GenS5{F0: 520, F1: 42, F2: 127}
                _ = fi54
                _ = v55
            }
        case 1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v4 = v4
        }
    case 32768:
        if (!(!false)) {
            if ((-(v4)) == v4) {
                if true {
                    v8 = (v8 << ((v18) & 63))
                    v4 = (-(v4))
                } else {
                    _ = len(v2)
                    v7 = v5
                }
                _ = len(v2)
                v3 = v3
            }
            if true {
                for fi56 := 0; fi56 < 18; fi56++ {
                    v3 += (v3 << ((42) & 63))
                    if true {
                        for fi57 := 0; fi57 < 8; fi57++ {
                            var v58 uint16 = 0
                            fmt.Println(v17, fi57)
                            fmt.Println(v3, v58, fi56)
                            v7 = ("a" + (v5 + ("a" + (v5 + ("" + v17)))))
                            _ = fi57
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 4; fi59++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch ((v17 + "\n\t") + ("a" + v17)) {
                            case "a":
                                v60 := GenS3{F0: v18, F1: (v18 & 2), F2: v8}
                                _ = v60
                            case "\n\t":
                                var v61 string = v5
                                _ = v61
                            default:
                                var v62 uint16 = 2
                                _ = v62
                            }
                            v63 := GenS7{F0: v7, F1: false, F2: 42}
                            _ = fi59
                            _ = v63
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 = "a"
                    }
                    var v64 uint = (v18 >> (((v18 ^ (v18 & (v18 << ((v18) & 63))))) & 63))
                    _ = fi56
                    _ = v64
                }
                fmt.Println(v18, v18, v17)
            } else {
                v3 = 0
                for fi65 := 0; fi65 < 10; fi65++ {
                    if (((-2.25 + 1e10) <= 3.14159) && true) {
                        if ((3.14159 >= 3.14159) && (49 <= 1)) {
                            v4 -= v4
                            if len(v1) > 0 {
                                v1[0] = fi65
                            }
                            fmt.Println(v8, v8)
                        } else {
                            var v66 int16 = (v8 & (v8 ^ (v8 << ((v18) & 63))))
                            for fi67 := 0; fi67 < 10; fi67++ {
                                switch "foo bar" {
                                case "a":
                                    _ = len(v2)
                                case "":
                                    var v68 uint64 = 42
                                    _ = v68
                                case "hello":
                                    var v69 bool = true
                                    _ = v69
                                }
                                switch (^(v4)) {
                                case -1:
                                    v18 -= v18
                                case 1:
                                    switch ((v17 + (v7 + ("\n\t" + (v17 + v5)))) + v5) {
                                    case "":
                                        switch (v4 & (fi65 | 2147483647)) {
                                        case 42:
                                            var v70 uint32 = 4294967295
                                            _ = v70
                                        default:
                                            var v71 int8 = 0
                                            _ = v71
                                        }
                                        v66 = v8
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v18, fi67, fi67)
                                    case "\n\t":
                                        if true {
                                            v18 -= 1
                                        }
                                    case "foo bar":
                                        v18 -= (v18 >> (((v18 << ((v18) & 63))) & 63))
                                    default:
                                        v8 = (v66 << ((553) & 63))
                                    }
                                default:
                                    v17 = (v17 + ((("" + (v17 + (((v5 + ("foo bar" + "hello")) + v5) + (v7 + (v7 + ("a" + "a")))))) + ("hello" + v7)) + v17))
                                }
                                v72 := GenS8{F0: fi67}
                                _ = fi67
                                _ = v72
                            }
                            _ = v66
                        }
                        for fi73 := 0; fi73 < 7; fi73++ {
                            v18 = v18
                            var v74 int64 = -1
                            _ = fi73
                            _ = v74
                        }
                        var v75 string = v17
                        _ = v75
                    }
                    for fi76 := 0; fi76 < 8; fi76++ {
                        switch 4294967295 {
                        case 2147483647:
                            v77 := GenS9{F0: 127}
                            _ = v77
                        }
                        v18 += v18
                        var v78 uint8 = 127
                        _ = fi76
                        _ = v78
                    }
                    _ = fi65
                }
                if len(v1) > 0 {
                    v1[0] = v4
                }
            }
            for fi79 := 0; fi79 < 13; fi79++ {
                if (false || false) {
                    _ = len(v2)
                    for fi80 := 0; fi80 < 8; fi80++ {
                        var v81 string = "a"
                        var v82 uint8 = 0
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                }
                _ = fi79
            }
        } else {
            v83 := GenS5{F0: 32767, F1: 1, F2: 178}
            _ = v83
        }
    case 1:
        v7 = v5
    default:
        if len(v1) > 0 {
            v1[0] = 59
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
