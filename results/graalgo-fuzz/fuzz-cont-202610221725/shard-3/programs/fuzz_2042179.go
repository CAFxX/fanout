package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 string }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 bool; F1 string }
func genFn0(p0 uint16) uint {
    if (!false) {
        var v4 int8 = 100
        _ = v4
    } else {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int = -2147483648
            _ = fi5
            _ = v6
        }
        if (-1 > -32768) {
            for fi7 := 0; fi7 < 8; fi7++ {
                p0 -= p0
                p0 += 0
                _ = fi7
            }
        } else {
            var v8 int = -126
            _ = v8
        }
        fmt.Println(p0)
    }
    v9 := GenS1{F0: 0, F1: -9223372036854775808, F2: 255}
    fmt.Println(p0)
    _ = p0
    _ = v9
    return 42
}

func main() {
    v1 := []int{4294967295, -2147483648, -2147483648}
    v2 := make(map[int]int)
    v3 := []string{"foo bar"}
    var v10 int16 = -849
    if true {
        if (-1 >= -1) {
            switch "a" {
            case "foo bar":
                if false {
                    var v11 uint = 403
                    _ = v11
                }
            case "a":
                var v12 uint32 = 42
                _ = v12
            case "\n\t":
                var v13 int8 = 42
                _ = v13
            }
            if (42 > 0) {
                if (32767 >= 939) {
                    fmt.Println(v10)
                    var v14 int64 = -1
                    _ = v14
                }
                var v15 int16 = v10
                _ = v15
            }
        } else {
            if (false || ((!false) && (!(!false)))) {
                if (true || (379 == 65535)) {
                    if true {
                        v16 := GenS2{F0: 42, F1: 65535, F2: "foo bar"}
                        v10 += (v10 - v10)
                        _ = v16
                    } else {
                        switch -341 {
                        case 9223372036854775807:
                            v17 := GenS2{F0: -2147483648, F1: 65535, F2: ("" + "a")}
                            _ = v17
                        case -2147483648:
                            fmt.Println(v10)
                        }
                        var v18 uint64 = 4294967295
                        _ = v18
                    }
                    var v19 int8 = -128
                    _ = v19
                }
                if (false && true) {
                    fmt.Println(v10)
                } else {
                    if (42 <= -9223372036854775808) {
                        fmt.Println(v10)
                        if (("a" <= "a") || true) {
                            v20 := GenS2{F0: 916, F1: 229, F2: "foo bar"}
                            var v21 int16 = v10
                            fmt.Println(v10, v21)
                            _ = v20
                            _ = v21
                        }
                    } else {
                        var v22 int8 = -128
                        _ = v22
                    }
                    v23 := GenS1{F0: int64(v10), F1: 9223372036854775807, F2: 4294967295}
                    var v24 uint16 = 42
                    _ = v23
                    _ = v24
                }
                v10 -= v10
            }
            delete(v2, -826)
            switch 42 {
            case 42:
                switch "foo bar" {
                case "":
                    if false {
                        fmt.Println(v10)
                        if len(v1) > 0 {
                            v1[0] = 808
                        }
                        for fi25 := 0; fi25 < 9; fi25++ {
                            fmt.Println(fi25)
                            for fi26 := 0; fi26 < 16; fi26++ {
                                if false {
                                    switch fi25 {
                                    case -259:
                                        v27 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: 18446744073709551615}
                                        _ = v27
                                    case -2147483648:
                                        v10 = v10
                                    case 445:
                                        _ = len(v2)
                                    }
                                    v10 += (v10 % ((v10) | 1))
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(fi25, v10)
                                    case "hello":
                                        v10 = v10
                                    case "foo bar":
                                        v10 = (v10 >> ((65535) & 63))
                                    default:
                                        v10 = v10
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if (("" + "foo bar") != "\n\t") {
                                            switch "" {
                                            case "":
                                                fmt.Println(fi25, fi26, fi26)
                                            }
                                            if (!false) {
                                                if false {
                                                    if len(v1) > 0 {
                                                        v1[0] = 42
                                                    }
                                                } else {
                                                    v28 := GenS2{F0: 42, F1: 2, F2: "hello"}
                                                    var v29 int64 = -9223372036854775808
                                                    v10 = (v10 - 42)
                                                    _ = v28
                                                    _ = v29
                                                }
                                                v10 = (v10 | (-(v10)))
                                            }
                                        } else {
                                            if true {
                                                v10 += (v10 % ((v10) | 1))
                                                v30 := GenS1{F0: 0, F1: fi26, F2: 4294967295}
                                                _ = v30
                                            } else {
                                                v31 := GenS1{F0: 833, F1: fi26, F2: 767}
                                                v10 = v10
                                                switch 127 {
                                                case 127:
                                                    var v32 uint16 = 65535
                                                    _ = v32
                                                case 32768:
                                                    var v33 uint8 = 195
                                                    _ = v33
                                                }
                                                v10 = v10
                                                _ = v31
                                            }
                                            switch 935 {
                                            case -2147483648:
                                                v34 := GenS1{F0: 654, F1: (fi26 | fi25), F2: 4294967295}
                                                _ = v34
                                            }
                                            v10 = v10
                                            if true {
                                                var v35 int32 = 0
                                                if (true || false) {
                                                    var v36 uint = 42
                                                    fmt.Println(v10, fi26)
                                                    _ = v36
                                                }
                                                v35 += v35
                                                _ = v35
                                            }
                                        }
                                    case "foo bar":
                                        v37 := GenS3{F0: 1, F1: fi25}
                                        delete(v2, (fi25 & (^(fi26))))
                                        _ = v37
                                    case "a":
                                        v10 = -214
                                    }
                                }
                                _ = fi26
                            }
                            _ = fi25
                        }
                    }
                default:
                    v38 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: 255}
                    _ = v38
                }
            case -9223372036854775808:
                if true {
                    if false {
                        switch (("a" + "foo bar") + "a") {
                        case "hello":
                            if false {
                                var v39 int64 = -9223372036854775808
                                _ = v39
                            } else {
                                v40 := GenS1{F0: 0, F1: 32767, F2: 474}
                                var v41 uint64 = 4294967295
                                v42 := GenS4{F0: -838, F1: 0}
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        default:
                            v10 = (v10 ^ v10)
                        }
                        fmt.Println(v10)
                    }
                    delete(v2, -2147483648)
                    fmt.Println(v10)
                }
            case -915:
                for fi43 := 0; fi43 < 16; fi43++ {
                    for fi44 := 0; fi44 < 6; fi44++ {
                        v45 := GenS5{F0: 128, F1: false}
                        if (!false) {
                            v46 := GenS4{F0: (^(fi44)), F1: 0}
                            _ = v46
                        }
                        _ = fi44
                        _ = v45
                    }
                    _ = fi43
                }
            default:
                if (!(true || true)) {
                    var v47 uint16 = 2
                    _ = v47
                } else {
                    fmt.Println(v10)
                }
            }
        }
        var v48 float32 = -2.25
        _ = v48
    } else {
        for fi49 := 0; fi49 < 11; fi49++ {
            v10 = (v10 % (((-(v10))) | 1))
            var v50 uint64 = 2
            _ = fi49
            _ = v50
        }
        switch 0 {
        case 42:
            for fi51 := 0; fi51 < 15; fi51++ {
                fmt.Println(v10, fi51)
                v10 -= -1
                v10 += v10
                _ = fi51
            }
        case -2147483648:
            v10 += (v10 / ((-32768) | 1))
        default:
            var v52 int = 1
            _ = v52
        }
        v53 := GenS5{F0: 127, F1: (!false)}
        _ = v53
    }
    v10 = (v10 & v10)
    if ("" <= ("hello" + "")) {
        var v54 bool = ((!true) && (true || (100 != 127)))
        if (42 != 0) {
            var v55 uint = 150
            switch 2147483647 {
            case -2147483648:
                for fi56 := 0; fi56 < 2; fi56++ {
                    fmt.Println(v54, v54)
                    fmt.Println(v10, v54)
                    for fi57 := 0; fi57 < 7; fi57++ {
                        v54 = ((348 == 2) || v54)
                        v10 = (v10 << ((18446744073709551615) & 63))
                        _ = fi57
                    }
                    _ = fi56
                }
            case 4294967295:
                var v58 uint32 = 1
                _ = v58
            default:
                var v59 bool = true
                _ = v59
            }
            fmt.Println(v10)
            _ = v55
        } else {
            var v60 float64 = (3.14159 - 0.1)
            v60 = 3.14159
            fmt.Println(v10, v60)
            _ = v60
        }
        v61 := GenS6{F0: 1, F1: 32767, F2: 737}
        _ = v54
        _ = v61
    }
    v62 := GenS7{F0: 583, F1: 0}
    fmt.Println(v10)
    if true {
        _ = len(v1)
        var v63 int = 1
        _ = v63
    } else {
        if true {
            v10 = v10
            if len(v1) > 0 {
                v1[0] = 42
            }
        } else {
            v10 += (v10 / ((0) | 1))
            v64 := GenS2{F0: 2147483647, F1: 255, F2: "foo bar"}
            _ = v64
        }
        var v65 int = -315
        _ = v65
    }
    fmt.Println(v10)
    v66 := GenS8{F0: (false || false), F1: "a"}
    fmt.Println(v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v62
    _ = v66
}
