package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 bool; F1 int64; F2 int32 }
type GenS8 struct { F0 uint; F1 int8; F2 int16 }
type GenS9 struct { F0 int8; F1 uint }

func main() {
    var v1 int16 = 1
    var v2 uint = 0
    var v3 uint64 = 18446744073709551615
    var v4 int32 = -1
    var v5 float32 = float32(0.0)
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v5, fi7, v2)
        fmt.Println(v2, v4)
        _ = fi7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{-953, fi8, (fi8 ^ (fi8 >> (((v2 + v2)) & 63))), 0}
        _ = fi8
        _ = v9
    }
    v10 := []int{134, 127, 2147483647, 128}
    var v11 uint8 = 3
    var v12 int16 = v1
    v3 = v3
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 string = (("foo bar" + "\n\t") + (("" + "foo bar") + "hello"))
            v11 = 1
            fmt.Println(v5, fi14)
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v12 = (-(v12))
    if (v6 <= float64(float32(v6))) {
        v12 += v12
    } else {
        if true {
            v16 := GenS1{F0: 127, F1: 817, F2: 0}
            v11 = v11
            v17 := GenS1{F0: 1, F1: 162, F2: 1}
            _ = v16
            _ = v17
        }
        v18 := GenS2{F0: (v11 - v11)}
        v6 = 3.14159
        _ = v18
    }
    var v19 float32 = v5
    var v20 int8 = -114
    v21 := GenS1{F0: (v20 & (v20 ^ (v20 ^ (v20 + (v20 | 100))))), F1: 2, F2: 2}
    v6 -= float64((v19 / v19))
    fmt.Println(v6, v2, v1)
    if false {
        var v22 uint16 = 178
        v23 := GenS3{F0: (v4 - int32(v11)), F1: (("hello" + "hello") + (("" + ("\n\t" + "hello")) + "hello"))}
        _ = v22
        _ = v23
    } else {
        var v24 int = -9223372036854775808
        if true {
            switch (v24 & (v24 >> ((v2) & 63))) {
            case 2147483647:
                if (true || ((false && false) || false)) {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        switch (v24 % (((fi25 / ((v24) | 1))) | 1)) {
                        case 680:
                            for fi26 := 0; fi26 < 1; fi26++ {
                                var v27 uint64 = (v3 + 0)
                                var v28 int16 = v1
                                if len(v10) > 0 {
                                    v10[0] = (fi26 % ((v24) | 1))
                                }
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        case 9223372036854775807:
                            var v29 int32 = (v4 % (((v4 + 1)) | 1))
                            _ = v29
                        case 127:
                            var v30 float64 = v6
                            _ = v30
                        }
                        if false {
                            if (true && true) {
                                if (false || true) {
                                    fmt.Println(v5, v19, v12)
                                    var v31 float32 = float32(v6)
                                    if true {
                                        for fi32 := 0; fi32 < 8; fi32++ {
                                            var v33 bool = true
                                            _ = fi32
                                            _ = v33
                                        }
                                        var v34 int64 = -9223372036854775808
                                        v35 := GenS4{F0: -1}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v3 -= (v3 + v3)
                                        var v36 uint64 = (^(v3))
                                        for fi37 := 0; fi37 < 12; fi37++ {
                                            v38 := GenS5{F0: 1, F1: (^(v11))}
                                            v2 = v2
                                            fmt.Println(fi37)
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            v12 = (v1 << (((^(v2))) & 63))
                                            _ = fi37
                                            _ = v38
                                        }
                                        _ = v36
                                    }
                                    _ = v31
                                }
                                fmt.Println(v3, v2, v1)
                                var v39 float32 = ((v5 / v5) * 3.14159)
                                _ = v39
                            } else {
                                for fi40 := 0; fi40 < 8; fi40++ {
                                    v41 := GenS4{F0: 42}
                                    if (-2.25 > (float64(v5) * 1.5)) {
                                        var v42 uint16 = 706
                                        var v43 int32 = (^(v4))
                                        v44 := GenS6{F0: (-(fi40)), F1: (fi40 / ((fi40) | 1))}
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = fi40
                                    _ = v41
                                }
                            }
                            v2 += v2
                            if true {
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    fmt.Println(fi25, v1)
                                    if (((fi45 + (-(v24))) <= (-(fi45))) && (4294967295 != v24)) {
                                        var v46 uint8 = (v11 * v11)
                                        if true {
                                            var v47 uint = (v2 / (((^(v2))) | 1))
                                            _ = v47
                                        }
                                        switch "foo bar" {
                                        case "":
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                        case "hello":
                                            var v48 int64 = -1
                                            _ = v48
                                        case "a":
                                            v49 := GenS7{F0: ((v2 < (^(v2))) && ((v2 * v2) >= 1)), F1: -892, F2: (^(v4))}
                                            _ = v49
                                        }
                                        _ = v46
                                    } else {
                                        fmt.Println(v5, v3)
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                    _ = fi45
                                }
                                fmt.Println(fi25)
                                v2 += v2
                            } else {
                                if true {
                                    var v50 int8 = v20
                                    v51 := GenS8{F0: v2, F1: (v50 << ((2) & 63)), F2: 32767}
                                    var v52 uint = (v2 ^ 0)
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 bool = (("\n\t" + "") > "a")
                                    _ = v53
                                }
                                var v54 uint = v2
                                var v55 uint32 = 42
                                _ = v54
                                _ = v55
                            }
                        } else {
                            v24 += (-(v24))
                            if (0 == 706) {
                                fmt.Println(v11, fi25)
                                if (!(true && false)) {
                                    for fi56 := 0; fi56 < 15; fi56++ {
                                        switch "" {
                                        case "a":
                                            v1 = v12
                                        case "foo bar":
                                            fmt.Println(v6, v19)
                                        case "":
                                            var v57 int = (fi56 - fi25)
                                            _ = v57
                                        }
                                        if (((861 <= 669) && true) || true) {
                                            var v58 int8 = v20
                                            v10 = append(v10, fi25)
                                            v4 = (-(v4))
                                            v12 = (-(v12))
                                            _ = v58
                                        }
                                        _ = fi56
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v6)
                                    case "hello":
                                        for fi59 := 0; fi59 < 18; fi59++ {
                                            v60 := GenS9{F0: 42, F1: 159}
                                            var v61 uint16 = 191
                                            v4 = (v4 - int32(v2))
                                            if ((-1 == int8(v2)) && false) {
                                                var v62 string = "hello"
                                                v63 := GenS8{F0: 949, F1: 0, F2: (v1 | v12)}
                                                var v64 int32 = v4
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            }
                                            var v65 uint32 = 2147483647
                                            _ = fi59
                                            _ = v60
                                            _ = v61
                                            _ = v65
                                        }
                                    case "a":
                                        for fi66 := 0; fi66 < 8; fi66++ {
                                            _ = len(v10)
                                            _ = fi66
                                        }
                                    default:
                                        var v67 int16 = -236
                                        _ = v67
                                    }
                                    v5 = 1e10
                                } else {
                                    var v68 int8 = (v20 + (v20 * (v20 + (-(v20)))))
                                    fmt.Println(v11, v68)
                                    fmt.Println(v24, v5)
                                    _ = v68
                                }
                            }
                        }
                        _ = fi25
                    }
                    var v69 uint32 = 136
                    _ = v69
                } else {
                    fmt.Println(v12, v1)
                    var v70 int16 = (v1 % ((v1) | 1))
                    v71 := GenS8{F0: uint(v11), F1: (^(v20)), F2: (v1 + 1000)}
                    _ = v70
                    _ = v71
                }
            }
        }
        v5 = 3.14159
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
