package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int64; F2 int64 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int8; F1 uint; F2 string }
type GenS8 struct { F0 bool; F1 int16; F2 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 65535
    var v3 uint8 = 1
    v1 = v1
    if (18446744073709551615 <= 304) {
        v1 = 9223372036854775807
        v1 = (v1 & v1)
    } else {
        v4 := GenS1{F0: (v1 - v1), F1: 42, F2: (v2 << ((4294967295) & 63))}
        v2 = 65535
        v5 := GenS2{F0: v2}
        _ = v4
        _ = v5
    }
    if false {
        v2 = v2
        v1 -= (v1 | v1)
    }
    if true {
        fmt.Println(v1, v1, v3)
        if true {
            v6 := []int{-2147483648}
            if true {
                switch "hello" {
                case "a":
                    v3 = v3
                case "\n\t":
                    var v7 uint64 = 887
                    _ = v7
                default:
                    if (2 <= 1) {
                        if (1 != 105) {
                            if true {
                                for fi8 := 0; fi8 < 8; fi8++ {
                                    for fi9 := 0; fi9 < 2; fi9++ {
                                        v2 += v2
                                        switch "a" {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "\n\t":
                                                v2 = v2
                                            default:
                                                var v10 uint16 = 421
                                                _ = v10
                                            }
                                        default:
                                            fmt.Println(fi8)
                                        }
                                        v3 = v3
                                        _ = fi9
                                    }
                                    v3 += v3
                                    _ = fi8
                                }
                                if ((v2 & v2) <= (v2 << ((uint(v3)) & 63))) {
                                    _ = len(v6)
                                    var v11 uint = 255
                                    v11 -= (v11 ^ (v11 * (^(v11))))
                                    _ = v11
                                }
                            }
                        } else {
                            var v12 int64 = (v1 & v1)
                            _ = v12
                        }
                        switch "hello" {
                        case "\n\t":
                            var v13 string = ("a" + (("\n\t" + (((("" + "\n\t") + "foo bar") + "foo bar") + "hello")) + ("a" + ("foo bar" + ""))))
                            _ = v13
                        case "a":
                            if len(v6) > 0 {
                                v6[0] = 9223372036854775807
                            }
                        case "":
                            fmt.Println(v1)
                        default:
                            v14 := GenS2{F0: (v2 % ((v2) | 1))}
                            _ = v14
                        }
                        v3 += 128
                    }
                }
                if (-2.25 >= 3.14159) {
                    fmt.Println(v3, v3)
                    if (false && true) {
                        v15 := GenS3{F0: (!((0 >= 42) && ((!false) || true))), F1: (v2 / ((v2) | 1))}
                        _ = v15
                    }
                    v16 := GenS4{F0: 4294967295, F1: (v1 | (^(v1))), F2: -1}
                    _ = v16
                }
                v3 += v3
            } else {
                v17 := GenS5{F0: "", F1: -222}
                var v18 int8 = -59
                _ = v17
                _ = v18
            }
            v19 := GenS6{F0: (v3 >> ((1) & 63)), F1: 65535}
            _ = v19
            _ = v6
        }
        if ("" != "a") {
            var v20 string = "a"
            _ = v20
        } else {
            var v21 int = 1
            v21 = v21
            var v22 string = "hello"
            _ = v21
            _ = v22
        }
    }
    if (int8(v1) > 1) {
        v23 := []int{}
        if (true || true) {
            fmt.Println(v2, v1)
            if (("hello" + (("hello" + "\n\t") + (("\n\t" + "") + "foo bar"))) == "") {
                v23 = append(v23, 0)
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v3)
                    v25 := GenS5{F0: ("\n\t" + "\n\t"), F1: -1}
                    _ = fi24
                    _ = v25
                }
            } else {
                if false {
                    _ = len(v23)
                    _ = len(v23)
                }
                _ = len(v23)
            }
            fmt.Println(v3)
        } else {
            v2 = 32767
            var v26 uint32 = 551
            var v27 uint16 = v2
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int8 = 95
            if false {
                v2 = v2
                fmt.Println(v2)
            } else {
                for fi30 := 0; fi30 < 1; fi30++ {
                    fmt.Println(fi30, fi30, fi30)
                    _ = len(v23)
                    v2 = (v2 << ((889) & 63))
                    v23 = append(v23, (-(fi30)))
                    _ = fi30
                }
            }
            for fi31 := 0; fi31 < 3; fi31++ {
                switch (("" + (("foo bar" + "a") + "\n\t")) + "foo bar") {
                case "a":
                    var v32 float64 = 0.0
                    _ = v32
                }
                if len(v23) > 0 {
                    v23[0] = (fi28 + int(v2))
                }
                v33 := GenS7{F0: (v29 * v29), F1: 18446744073709551615, F2: "hello"}
                _ = fi31
                _ = v33
            }
            _ = fi28
            _ = v29
        }
        _ = v23
    } else {
        v34 := []int{585, 4294967295, 2147483647, 32767}
        _ = v34
    }
    if true {
        fmt.Println(v1, v3)
        switch "a" {
        case "a":
            v35 := []int{650, 9223372036854775807}
            _ = v35
        case "\n\t":
            var v36 uint64 = 301
            _ = v36
        case "":
            for fi37 := 0; fi37 < 14; fi37++ {
                v1 -= (^(v1))
                switch 2147483647 {
                case 9223372036854775807:
                    v38 := GenS3{F0: false, F1: 762}
                    _ = v38
                case 0:
                    v39 := GenS6{F0: (v3 + 0), F1: v2}
                    _ = v39
                }
                _ = fi37
            }
        default:
            switch (("hello" + "\n\t") + ("" + "foo bar")) {
            case "a":
                v40 := GenS2{F0: (v2 << ((504) & 63))}
                _ = v40
            case "hello":
                v3 = v3
            default:
                if (!((-2.25 - (1e10 - (1e10 / 1.0))) != (1.5 / ((float32(3.14159) - 0.0) * float32((1e10 / -2.25)))))) {
                    for fi41 := 0; fi41 < 20; fi41++ {
                        for fi42 := 0; fi42 < 14; fi42++ {
                            v43 := []int{(^(fi41))}
                            v44 := GenS4{F0: 2, F1: v1, F2: v1}
                            var v45 uint16 = 8
                            _ = fi42
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        _ = fi41
                    }
                    for fi46 := 0; fi46 < 12; fi46++ {
                        switch "hello" {
                        case "a":
                            v1 = v1
                        }
                        _ = fi46
                    }
                    if false {
                        v47 := []int{}
                        var v48 float64 = 1.5
                        _ = v47
                        _ = v48
                    }
                } else {
                    v2 = v2
                    fmt.Println(v3, v1, v3)
                }
            }
        }
        switch "hello" {
        case "":
            switch "hello" {
            case "a":
                var v49 int8 = 0
                _ = v49
            case "hello":
                var v50 int = 1
                _ = v50
            }
        case "a":
            v3 = v3
        default:
            switch "\n\t" {
            case "foo bar":
                var v51 int = -108
                _ = v51
            }
        }
    }
    v1 += (v1 >> ((184) & 63))
    switch 2147483647 {
    case -872:
        var v52 uint16 = v2
        _ = v52
    case 1:
        v53 := []int{}
        _ = v53
    }
    v3 += (v3 + 0)
    if true {
        v54 := []int{9223372036854775807, 721, -2147483648}
        _ = v54
    } else {
        fmt.Println(v3)
    }
    var v55 uint32 = 4294967295
    v55 = v55
    if true {
        for fi56 := 0; fi56 < 2; fi56++ {
            v57 := GenS8{F0: (!(v1 < (-(v1)))), F1: -1, F2: -1}
            for fi58 := 0; fi58 < 11; fi58++ {
                v59 := GenS6{F0: 1, F1: (v2 & (v2 ^ 768))}
                fmt.Println(fi58)
                fmt.Println(fi58, fi58)
                _ = fi58
                _ = v59
            }
            _ = fi56
            _ = v57
        }
        var v60 int64 = 9223372036854775807
        var v61 uint = 758
        _ = v60
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
}
