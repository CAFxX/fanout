package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int64; F1 uint32; F2 int32 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 int; F2 uint16 }
func genFn0(p0 int, p1 uint64, p2 uint) int8 {
    var v5 int8 = 52
    switch 42 {
    case 0:
        p2 = 255
    default:
        v6 := []int{(p0 * 127)}
        _ = v6
    }
    v7 := GenS1{F0: 0}
    v8 := GenS1{F0: 0}
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    _ = v8
    _ = v9
    return (v5 & 42)
}
func genFn1() int8 {
    var v10 float32 = (1.5 + float32(3.14159))
    fmt.Println(v10)
    v10 -= ((v10 / v10) - v10)
    v10 = ((-2.25 * v10) - 0.0)
    _ = v10
    return -128
}

func main() {
    var v1 uint32 = 640
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int64 = -9223372036854775808
        _ = fi11
        _ = v12
    }
    v13 := []int{9223372036854775807, 42}
    if (-2147483648 <= 32768) {
        var v14 int = -9223372036854775808
        _ = v14
    }
    if false {
        v4["a"] = 1
        _ = len(v2)
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint8 = 167
            v1 = v1
            _ = fi15
            _ = v16
        }
        v13 = append(v13, 212)
    }
    var v17 int16 = 32767
    fmt.Println(v1, v1)
    var v18 uint32 = (^(v1))
    for fi19 := 0; fi19 < 5; fi19++ {
        v1 = 451
        _ = fi19
    }
    v20 := GenS2{F0: 42, F1: "\n\t"}
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS1{F0: 127}
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 int64 = -1
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    v25 := GenS1{F0: 41}
    if true {
        v26 := v3[127]
        var v27 int8 = 33
        v27 = (-(v27))
        _ = v26
        _ = v27
    }
    var v28 uint = 255
    v18 = (v1 % ((v18) | 1))
    switch ("a" + "hello") {
    case "":
        var v29 int8 = 1
        _ = v29
    case "a":
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint32 = v18
            _ = fi30
            _ = v31
        }
    case "hello":
        if true {
            if false {
                for fi32 := 0; fi32 < 4; fi32++ {
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v17 -= v17
                        switch (^(fi33)) {
                        case 32768:
                            var v34 uint32 = 2147483648
                            _ = v34
                        case 127:
                            v28 += v28
                        case 42:
                            switch "\n\t" {
                            case "\n\t":
                                var v35 uint32 = v18
                                _ = v35
                            case "foo bar":
                                v36 := GenS3{F0: 1}
                                _ = v36
                            default:
                                v1 = v18
                            }
                        }
                        _ = fi33
                    }
                    v37 := GenS4{F0: 18446744073709551615, F1: -27}
                    _ = fi32
                    _ = v37
                }
                switch ("a" + "a") {
                case "foo bar":
                    var v38 uint = (v28 | v28)
                    _ = v38
                case "":
                    v17 = 411
                }
                var v39 int64 = 991
                _ = v39
            }
            v1 -= (v1 & v1)
            for fi40 := 0; fi40 < 5; fi40++ {
                v2["foo bar"] = 65535
                if true {
                    var v41 float32 = (-2.25 + (0.0 * (0.0 + float32(float64(float32((((float64(-2.25) + 0.0) + 3.14159) * float64(3.14159))))))))
                    _ = v41
                } else {
                    if true {
                        var v42 string = "hello"
                        if (((((^(v1)) <= v18) && (1 >= 18446744073709551615)) || true) || (genFn0(fi40, uint64(v28), 65535) < genFn0(fi40, 42, 4294967295))) {
                            v43 := GenS5{F0: 121, F1: 255, F2: v1}
                            v17 += (v17 | 0)
                            v44 := GenS6{F0: -1, F1: (v18 << (((v28 - v28)) & 63)), F2: 0}
                            _ = v43
                            _ = v44
                        } else {
                            var v45 float64 = float64(1e10)
                            var v46 int64 = -9223372036854775808
                            for fi47 := 0; fi47 < 14; fi47++ {
                                var v48 int32 = 1
                                _ = fi47
                                _ = v48
                            }
                            _ = v45
                            _ = v46
                        }
                        for fi49 := 0; fi49 < 19; fi49++ {
                            var v50 int32 = 1
                            _ = fi49
                            _ = v50
                        }
                        _ = v42
                    }
                    for fi51 := 0; fi51 < 2; fi51++ {
                        fmt.Println(v17, v28)
                        _ = fi51
                    }
                }
                for fi52 := 0; fi52 < 19; fi52++ {
                    fmt.Println(v18, v1)
                    v53 := GenS4{F0: 0, F1: int8(fi40)}
                    var v54 int64 = 9223372036854775807
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = fi40
            }
        } else {
            v1 += v1
            switch "a" {
            case "hello":
                switch -411 {
                case -2147483648:
                    v55 := GenS7{F0: (v18 << ((v28) & 63)), F1: -128}
                    _ = v55
                case 32767:
                    for fi56 := 0; fi56 < 15; fi56++ {
                        if (!(!false)) {
                            if len(v13) > 0 {
                                v13[0] = fi56
                            }
                            if false {
                                if (!(!(true && true))) {
                                    if false {
                                        switch ("foo bar" + "") {
                                        case "foo bar":
                                            switch ("" + "\n\t") {
                                            case "hello":
                                                v18 += v1
                                            case "foo bar":
                                                v18 += v1
                                            }
                                        }
                                        v57 := v3[(fi56 * fi56)]
                                        if true {
                                            switch "foo bar" {
                                            case "a":
                                                v58 := GenS8{F0: (^(v17)), F1: v17}
                                                _ = v58
                                            case "\n\t":
                                                fmt.Println(fi56)
                                            case "foo bar":
                                                switch v57 {
                                                case 128:
                                                    v17 += v17
                                                default:
                                                    v17 = 1000
                                                }
                                            default:
                                                for fi59 := 0; fi59 < 13; fi59++ {
                                                    if true {
                                                        v1 -= 583
                                                        var v60 uint32 = v1
                                                        v18 = 42
                                                        var v61 int64 = 1
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    var v62 string = "\n\t"
                                                    v28 -= (v28 | (^(v28)))
                                                    _ = fi59
                                                    _ = v62
                                                }
                                            }
                                            fmt.Println(v28, v18)
                                        } else {
                                            var v63 uint8 = 114
                                            _ = v63
                                        }
                                        _ = v57
                                    }
                                    if false {
                                        if ((((1e10 + 3.14159) - -2.25) * (1.5 * (1e10 * 0.1))) <= 1e10) {
                                            fmt.Println(v1)
                                        } else {
                                            v1 = 564
                                        }
                                        fmt.Println(v17, fi56)
                                        v64 := GenS9{F0: 71}
                                        _ = v64
                                    }
                                } else {
                                    if true {
                                        for fi65 := 0; fi65 < 11; fi65++ {
                                            v13 = append(v13, 474)
                                            v17 -= (v17 % (((^(v17))) | 1))
                                            _ = fi65
                                        }
                                    } else {
                                        v1 += v1
                                    }
                                }
                            } else {
                                v4["a"] = (-(fi56))
                                for fi66 := 0; fi66 < 4; fi66++ {
                                    var v67 uint8 = 255
                                    if ((false && false) || ((^(v67)) <= 255)) {
                                        var v68 int8 = 90
                                        fmt.Println(v17, v67)
                                        if (0 != int32(fi66)) {
                                            v69 := GenS10{F0: (^(v17)), F1: fi56, F2: 65535}
                                            var v70 int64 = -1
                                            _ = v69
                                            _ = v70
                                        } else {
                                            v3[(fi66 & fi56)] = (fi66 >> (((v28 << ((v28) & 63))) & 63))
                                            v13 = append(v13, fi66)
                                            v71 := GenS10{F0: v17, F1: fi56, F2: 42}
                                            v28 -= 0
                                            _ = v71
                                        }
                                        if (-2.25 < 1.5) {
                                            v17 = (v17 + 42)
                                        }
                                        _ = v68
                                    }
                                    fmt.Println(fi56, v17, v67)
                                    var v72 int = (fi66 % ((320) | 1))
                                    if false {
                                        var v73 int64 = 1
                                        _ = v73
                                    }
                                    _ = fi66
                                    _ = v67
                                    _ = v72
                                }
                            }
                        } else {
                            v28 += (v28 % ((1) | 1))
                        }
                        switch (fi56 >> (((v28 >> ((42) & 63))) & 63)) {
                        case -1:
                            v17 = 42
                        case 127:
                            var v74 int64 = 0
                            _ = v74
                        case 674:
                            v1 = v1
                        default:
                            if false {
                                var v75 uint32 = (v18 << ((2) & 63))
                                var v76 int = fi56
                                fmt.Println(v1, v76, v76)
                                _ = v75
                                _ = v76
                            }
                        }
                        _ = fi56
                    }
                case 128:
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
            }
        }
    default:
        v1 = (v18 * (^(v18)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
