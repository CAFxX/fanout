package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1e10
    v3 := []int{-9223372036854775808, 32767, 127, 1}
    v4 := make(map[int]string)
    var v5 int64 = -1
    fmt.Println(v5)
    v6 := GenS1{F0: "\n\t"}
    var v7 int64 = (v5 / ((-1) | 1))
    var v8 int64 = v5
    switch "foo bar" {
    case "hello":
        switch -387 {
        case -659:
            v3 = append(v3, 128)
        case 2147483647:
            var v9 uint16 = 42
            _ = v9
        case -9223372036854775808:
            v8 = -347
        default:
            v2 += float32(-2.25)
        }
    default:
        _ = len(v3)
    }
    var v10 uint32 = 242
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(v7, v8, v10)
            v10 = v10
            _ = fi11
        }
    } else {
        var v12 bool = (758 != 1)
        _ = v12
    }
    v7 -= 9223372036854775807
    delete(v4, 32767)
    if (true && (!(42 <= (v10 * (v10 | 65535))))) {
        var v13 int16 = -32768
        if ((^(v8)) <= 0) {
            var v14 int64 = v8
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 bool = (fi15 >= 2147483647)
                _ = fi15
                _ = v16
            }
            v17 := v1[78]
            _ = v14
            _ = v17
        } else {
            for fi18 := 0; fi18 < 11; fi18++ {
                if (false && true) {
                    var v19 int = 32767
                    var v20 int64 = -416
                    _ = v19
                    _ = v20
                } else {
                    if (!true) {
                        if ((!(fi18 < (-(fi18)))) && true) {
                            if (-32768 != v13) {
                                var v21 int32 = 775
                                _ = len(v3)
                                _ = v21
                            } else {
                                v22 := GenS1{F0: "a"}
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 1; fi23++ {
                                if (!((true && (!true)) || (false && false))) {
                                    switch "" {
                                    case "hello":
                                        var v24 uint16 = 2
                                        _ = v24
                                    case "a":
                                        v10 = 2147483648
                                    case "\n\t":
                                        var v25 uint32 = (^(v10))
                                        _ = v25
                                    default:
                                        fmt.Println(fi18)
                                    }
                                    v26 := GenS2{F0: ((v10 << ((1) & 63)) > 0), F1: v10}
                                    _ = v26
                                } else {
                                    v27 := GenS3{F0: ""}
                                    v28 := GenS3{F0: "a"}
                                    if false {
                                        fmt.Println(v5, v13, v5)
                                        v29 := GenS4{F0: 2, F1: 229}
                                        _ = v29
                                    } else {
                                        switch (fi18 * fi18) {
                                        case 128:
                                            v30 := GenS5{F0: 209, F1: (^(v13)), F2: 2147483647}
                                            _ = v30
                                        case -9223372036854775808:
                                            var v31 uint = 212
                                            _ = v31
                                        }
                                        v10 = (v10 << ((4294967295) & 63))
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                _ = fi23
                            }
                        } else {
                            if true {
                                var v32 uint16 = 42
                                v8 = 436
                                _ = v32
                            } else {
                                v7 = v7
                            }
                            v4[-1] = "hello"
                        }
                        var v33 bool = (742 < 2)
                        _ = v33
                    } else {
                        v13 = v13
                        v34 := GenS6{F0: 0, F1: v5}
                        _ = v34
                    }
                }
                v2 -= (v2 + float32(float64(v2)))
                _ = fi18
            }
            for fi35 := 0; fi35 < 8; fi35++ {
                for fi36 := 0; fi36 < 15; fi36++ {
                    var v37 int64 = (v5 & (v5 % ((v7) | 1)))
                    v2 = v2
                    var v38 uint8 = 255
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v13 = (v13 - (v13 ^ 1000))
                _ = fi35
            }
            if (float64(0.1) <= 3.14159) {
                v7 = (v5 << ((0) & 63))
            } else {
                v39 := v4[32767]
                fmt.Println(v10, v7)
                var v40 int32 = 2147483647
                _ = v39
                _ = v40
            }
        }
        _ = v13
    } else {
        switch 435 {
        case -39:
            var v41 uint32 = (v10 + v10)
            _ = v41
        case -410:
            switch 299 {
            case -9223372036854775808:
                for fi42 := 0; fi42 < 7; fi42++ {
                    switch (fi42 & (fi42 << ((4294967295) & 63))) {
                    case 0:
                        var v43 int32 = 0
                        _ = v43
                    default:
                        for fi44 := 0; fi44 < 17; fi44++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (((v2 - v2) + 0.1) == v2) {
                                var v45 int64 = (v8 ^ v5)
                                _ = v45
                            } else {
                                v46 := GenS7{F0: 128}
                                v7 = -1
                                v47 := GenS5{F0: 42, F1: 1000, F2: (v10 >> ((2) & 63))}
                                _ = v46
                                _ = v47
                            }
                            _ = fi44
                        }
                    }
                    v5 = (v7 - -1)
                    _ = fi42
                }
            }
        case -878:
            if false {
                _ = len(v1)
                for fi48 := 0; fi48 < 10; fi48++ {
                    v49 := v1[2147483648]
                    _ = len(v4)
                    fmt.Println(v5)
                    _ = fi48
                    _ = v49
                }
            }
        default:
            var v50 bool = true
            _ = v50
        }
    }
    var v51 uint8 = 2
    switch (((("a" + ("\n\t" + "\n\t")) + "") + "a") + "\n\t") {
    case "a":
        for fi52 := 0; fi52 < 6; fi52++ {
            v51 = 62
            var v53 uint32 = (v10 % (((v10 % (((v10 % ((v10) | 1))) | 1))) | 1))
            if (true && (-24 < -104)) {
                v54 := GenS6{F0: 184, F1: v7}
                _ = v54
            }
            _ = fi52
            _ = v53
        }
    case "":
        switch 421 {
        case -567:
            if true {
                var v55 float32 = (-2.25 / (((v2 * v2) - (v2 + (((0.0 * v2) - 1.5) / v2))) - ((v2 - 3.14159) + v2)))
                if (42 < 255) {
                    switch "foo bar" {
                    case "a":
                        v7 += (^(v8))
                    case "\n\t":
                        if (false || false) {
                            fmt.Println(v7)
                        } else {
                            if (true || (false || false)) {
                                if (true && (true && true)) {
                                    fmt.Println(v51)
                                    var v56 int32 = 0
                                    switch "hello" {
                                    case "hello":
                                        v57 := GenS2{F0: ((32768 != int(v5)) || true), F1: (v10 << ((13) & 63))}
                                        _ = v57
                                    }
                                    _ = v56
                                } else {
                                    if false {
                                        fmt.Println(v2)
                                    } else {
                                        v58 := GenS2{F0: (!((65535 < 42) || (false && true))), F1: v10}
                                        var v59 uint8 = v51
                                        v2 -= 1e10
                                        _ = v58
                                        _ = v59
                                    }
                                }
                            } else {
                                var v60 uint32 = (^(v10))
                                if (127 >= 19) {
                                    v60 -= v10
                                    var v61 float32 = v55
                                    _ = v61
                                }
                                var v62 float64 = (1e10 - 1.5)
                                _ = v60
                                _ = v62
                            }
                        }
                    }
                    if false {
                        if true {
                            var v63 int32 = -2147483648
                            for fi64 := 0; fi64 < 16; fi64++ {
                                var v65 int32 = (v63 + (v63 | v63))
                                _ = fi64
                                _ = v65
                            }
                            _ = v63
                        } else {
                            var v66 int16 = 42
                            var v67 uint64 = 18446744073709551615
                            var v68 int32 = 2147483647
                            _ = v66
                            _ = v67
                            _ = v68
                        }
                        if true {
                            fmt.Println(v51, v55)
                        }
                    }
                }
                _ = v55
            }
        default:
            var v69 bool = true
            _ = v69
        }
    case "foo bar":
        fmt.Println(v2)
    default:
        if false {
            fmt.Println(v5, v7, v8)
            var v70 float32 = (float32(1e10) / 3.14159)
            _ = v70
        }
    }
    v2 = v2
    for fi71 := 0; fi71 < 11; fi71++ {
        switch 32768 {
        case -1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        switch (-(fi71)) {
        case 553:
            var v72 int8 = 127
            _ = v72
        default:
            var v73 uint64 = 2
            _ = v73
        }
        var v74 uint64 = 42
        _ = fi71
        _ = v74
    }
    delete(v1, (^(v10)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
    _ = v8
}
