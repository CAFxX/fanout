package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
type GenS3 struct { F0 int8; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint32; F1 string; F2 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int8; F2 int8 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS9 struct { F0 int16 }
func genFn0() int16 {
    var v4 uint8 = 1
    if true {
        v5 := []int{0}
        switch 32767 {
        case -819:
            for fi6 := 0; fi6 < 11; fi6++ {
                if len(v5) > 0 {
                    v5[0] = (fi6 | (fi6 ^ -641))
                }
                _ = fi6
            }
        case -832:
            var v7 uint16 = 32767
            _ = v7
        case -1:
            if false {
                switch "foo bar" {
                case "":
                    if len(v5) > 0 {
                        v5[0] = 128
                    }
                case "hello":
                    v8 := GenS1{F0: -1}
                    _ = v8
                case "\n\t":
                    fmt.Println(v4)
                }
                switch "" {
                case "\n\t":
                    var v9 uint64 = 315
                    _ = v9
                case "a":
                    fmt.Println(v4)
                case "hello":
                    var v10 float64 = 1e10
                    _ = v10
                }
                if (!true) {
                    var v11 int32 = int32(v4)
                    if false {
                        var v12 uint32 = 2
                        _ = v12
                    }
                    if (!true) {
                        v4 = v4
                    } else {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            v14 := GenS2{F0: (fi13 + 128), F1: -151, F2: 4294967295}
                            _ = fi13
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 18; fi15++ {
                            if true {
                                v11 -= v11
                                v16 := GenS3{F0: 100, F1: 19, F2: 4294967295}
                                var v17 uint64 = 2
                                _ = v16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 13; fi18++ {
                                var v19 uint64 = 2
                                if true {
                                    var v20 float32 = 1.5
                                    _ = len(v5)
                                    _ = v20
                                }
                                v5 = append(v5, (fi15 << ((42) & 63)))
                                _ = fi18
                                _ = v19
                            }
                            _ = fi15
                        }
                        if ((("\n\t" + "hello") + "hello") != ("hello" + "foo bar")) {
                            v21 := GenS4{F0: 65535, F1: 342}
                            _ = v21
                        } else {
                            fmt.Println(v4, v4)
                            var v22 int8 = -123
                            v23 := GenS5{F0: 1, F1: "hello", F2: 1}
                            _ = v22
                            _ = v23
                        }
                    }
                    _ = v11
                } else {
                    switch 0 {
                    case -1:
                        v4 += v4
                    case 127:
                        var v24 uint64 = 117
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v26 := GenS4{F0: 948, F1: 42}
                        fmt.Println(v4)
                        fmt.Println(fi25)
                        _ = fi25
                        _ = v26
                    }
                    switch 2147483647 {
                    case 9223372036854775807:
                        fmt.Println(v4)
                    case 0:
                        switch ((("foo bar" + "hello") + "hello") + "foo bar") {
                        case "foo bar":
                            fmt.Println(v4)
                        case "hello":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            v4 = (v4 / ((v4) | 1))
                        }
                    case 299:
                        v4 = (v4 | 127)
                    }
                }
            }
        default:
            var v27 bool = false
            _ = v27
        }
        v4 -= v4
        _ = v5
    } else {
        v28 := GenS6{F0: 718}
        if ((42 > 0) || true) {
            switch "hello" {
            case "\n\t":
                v4 += v4
            }
            switch 9223372036854775807 {
            case -810:
                if true {
                    var v29 uint = 238
                    v30 := []int{}
                    _ = v29
                    _ = v30
                }
            case -469:
                v31 := GenS1{F0: -1}
                _ = v31
            }
        }
        fmt.Println(v4)
        _ = v28
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 int64 = 852
        _ = fi32
        _ = v33
    }
    _ = v4
    return -673
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = ((-2.25 / (1e10 + 1.5)) / -2.25)
    var v3 uint32 = 2147483648
    v34 := GenS5{F0: (^(v3)), F1: "hello", F2: 332}
    var v35 float32 = (float32(1e10) + v2)
    v35 = (0.0 / v2)
    if ((((((v35 / 3.14159) + 1.5) / v35) - v35) + v35) > (float32((0.1 * 1e10)) + v35)) {
        for fi36 := 0; fi36 < 13; fi36++ {
            v3 += v3
            _ = fi36
        }
        v37 := GenS4{F0: 1, F1: -889}
        if true {
            var v38 int8 = 100
            _ = v38
        } else {
            v39 := GenS5{F0: (v3 << ((1) & 63)), F1: "", F2: 42}
            _ = v39
        }
        _ = v37
    }
    v2 = -2.25
    if (-80 >= 42) {
        fmt.Println(v1, v35)
        v40 := []int{-1, 2147483647, 9223372036854775807, 1}
        _ = v40
    }
    fmt.Println(v3)
    fmt.Println(v2)
    for fi41 := 0; fi41 < 8; fi41++ {
        for fi42 := 0; fi42 < 4; fi42++ {
            v43 := []int{fi42, (fi41 ^ 464)}
            _ = fi42
            _ = v43
        }
        v44 := GenS7{F0: true, F1: 100, F2: 39}
        for fi45 := 0; fi45 < 17; fi45++ {
            v3 -= 2
            _ = fi45
        }
        _ = fi41
        _ = v44
    }
    switch 1 {
    case -912:
        v1 -= (^(v1))
    case 0:
        v46 := []int{66}
        _ = v46
    case 127:
        for fi47 := 0; fi47 < 4; fi47++ {
            for fi48 := 0; fi48 < 5; fi48++ {
                var v49 bool = false
                switch (fi48 & (fi48 | 1)) {
                case -9223372036854775808:
                    var v50 float64 = 1.5
                    _ = v50
                default:
                    if (false && v49) {
                        v51 := GenS1{F0: 1}
                        v49 = true
                        v3 -= 42
                        v52 := []int{128, fi48, (fi47 / ((fi48) | 1))}
                        fmt.Println(fi47)
                        v53 := GenS6{F0: 65535}
                        var v54 uint64 = 42
                        _ = v51
                        _ = v52
                        _ = v53
                        _ = v54
                    } else {
                        var v55 int8 = 0
                        var v56 int = (^(fi47))
                        _ = v55
                        _ = v56
                    }
                }
                v3 -= v3
                _ = fi48
                _ = v49
            }
            v57 := GenS6{F0: uint(v1)}
            v58 := GenS3{F0: -93, F1: -12, F2: v1}
            _ = fi47
            _ = v57
            _ = v58
        }
    default:
        v2 -= 1e10
    }
    v59 := GenS5{F0: v3, F1: ("foo bar" + "a"), F2: 2147483647}
    v60 := GenS8{F0: 261, F1: 42, F2: v3}
    v1 = v1
    for fi61 := 0; fi61 < 13; fi61++ {
        if (true || true) {
            switch "hello" {
            case "foo bar":
                v62 := []int{fi61, (fi61 ^ (fi61 - (fi61 << ((477) & 63))))}
                _ = v62
            case "":
                v2 = (-2.25 * v2)
            case "a":
                v63 := GenS9{F0: genFn0()}
                _ = v63
            default:
                fmt.Println(fi61, v3)
            }
            var v64 uint = 846
            var v65 float32 = 0.0
            _ = v64
            _ = v65
        } else {
            v1 = v1
        }
        _ = fi61
    }
    v66 := GenS6{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v59
    _ = v60
    _ = v66
}
