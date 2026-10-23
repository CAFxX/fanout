package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int16; F2 int32 }
type GenS7 struct { F0 int32; F1 bool; F2 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0() uint64 {
    var v4 uint32 = 1
    v5 := GenS1{F0: 0}
    fmt.Println(v4)
    v4 -= v4
    v4 += (v4 & (^(v4)))
    _ = v4
    _ = v5
    return 0
}
func genFn1() int8 {
    var v6 int32 = 370
    v7 := GenS1{F0: (v6 ^ (^(v6)))}
    switch ("\n\t" + "a") {
    case "a":
        v6 -= (-(v6))
    default:
        if ((-1 > -824) && (!true)) {
            if (!(true || true)) {
                switch 875 {
                case 128:
                    v6 = (v6 << ((463) & 63))
                }
            } else {
                v6 = 2147483647
                if false {
                    if (("a" + "hello") != ("" + "a")) {
                        v8 := GenS2{F0: 965}
                        v9 := GenS3{F0: 255, F1: -972, F2: 114}
                        var v10 int = 128
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                } else {
                    var v11 uint8 = 127
                    v12 := []int{0, 2147483647}
                    var v13 uint = 1
                    _ = v11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := []int{fi14, fi14, fi14, (fi14 + fi14)}
                    var v16 int64 = 134
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
        } else {
            v17 := []int{38, -2147483648, -566, -886}
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS2{F0: 4294967295}
                _ = fi18
                _ = v19
            }
            v20 := GenS4{F0: 283}
            _ = v17
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS2{F0: 42}
        var v23 int64 = 42
        v24 := []int{}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int64 = 1
    v2 := []string{"a", (("" + ("a" + "foo bar")) + "\n\t")}
    var v3 float32 = 1.5
    switch 128 {
    case -9223372036854775808:
        fmt.Println(v1, v1)
    default:
        var v25 int = 127
        _ = v25
    }
    fmt.Println(v3)
    v1 = 1
    _ = len(v2)
    switch "" {
    case "hello":
        for fi26 := 0; fi26 < 6; fi26++ {
            if true {
                if true {
                    var v27 bool = true
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = v27
                } else {
                    var v28 float32 = (1.5 * v3)
                    var v29 int32 = 2147483647
                    _ = len(v2)
                    _ = v28
                    _ = v29
                }
            }
            fmt.Println(v1, v1)
            v30 := GenS5{F0: 1}
            _ = fi26
            _ = v30
        }
    case "foo bar":
        if false {
            v31 := GenS3{F0: 32767, F1: 0, F2: 39}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 10; fi32++ {
                v1 += 897
                if (false || false) {
                    var v33 int8 = 42
                    _ = v33
                } else {
                    if false {
                        var v34 int16 = -1
                        switch "\n\t" {
                        case "":
                            var v35 float64 = (0.1 / 3.14159)
                            _ = v35
                        case "foo bar":
                            switch ("foo bar" + "foo bar") {
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "\n\t":
                                v36 := GenS6{F0: 923, F1: v34, F2: 14}
                                _ = v36
                            }
                        }
                        _ = v34
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                        case "a":
                            v37 := GenS2{F0: 18446744073709551615}
                            _ = v37
                        default:
                            switch ("\n\t" + ("\n\t" + "a")) {
                            case "\n\t":
                                var v38 uint32 = 1
                                _ = v38
                            case "foo bar":
                                var v39 uint = 2
                                _ = v39
                            }
                        }
                        v40 := GenS2{F0: 2}
                        _ = v40
                    }
                    if (2 <= 738) {
                        var v41 uint = 4294967295
                        for fi42 := 0; fi42 < 9; fi42++ {
                            var v43 int16 = -201
                            v44 := GenS7{F0: -1, F1: (!false), F2: (^(v43))}
                            v3 = ((v3 + v3) + v3)
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        if (!false) {
                            if true {
                                var v45 bool = false
                                _ = v45
                            }
                            v46 := GenS8{F0: 104}
                            _ = v46
                        }
                        _ = v41
                    }
                    var v47 bool = false
                    _ = v47
                }
                var v48 int16 = -32768
                _ = fi32
                _ = v48
            }
            for fi49 := 0; fi49 < 16; fi49++ {
                for fi50 := 0; fi50 < 5; fi50++ {
                    v1 = 9223372036854775807
                    v51 := GenS9{F0: 255, F1: 392}
                    _ = fi50
                    _ = v51
                }
                _ = fi49
            }
            switch (("hello" + "hello") + "hello") {
            case "a":
                var v52 uint = 0
                _ = v52
            }
        }
    }
    v1 += v1
    v53 := GenS1{F0: 0}
    if false {
        for fi54 := 0; fi54 < 8; fi54++ {
            v1 += v1
            _ = fi54
        }
        switch "" {
        case "foo bar":
            var v55 uint16 = 65535
            _ = v55
        }
    }
    v56 := GenS5{F0: 128}
    for fi57 := 0; fi57 < 18; fi57++ {
        var v58 uint8 = 128
        var v59 uint8 = v58
        _ = fi57
        _ = v58
        _ = v59
    }
    if false {
        if ((^(v1)) > -9223372036854775808) {
            v60 := GenS9{F0: 1, F1: 768}
            var v61 uint16 = 65535
            switch "" {
            case "hello":
                switch "foo bar" {
                case "a":
                    fmt.Println(v61, v61)
                case "\n\t":
                    var v62 int8 = -128
                    _ = v62
                default:
                    if (false && (!(577 > 682))) {
                        v3 = (((((((v3 + v3) * (v3 / -2.25)) / 1.5) / 1e10) + -2.25) * ((1.5 * (v3 / v3)) / 1.0)) - 0.1)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v63 int32 = 42
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        var v64 uint32 = 364
                        _ = v63
                        _ = v64
                    }
                }
            default:
                if len(v2) > 0 {
                    v2[0] = "a"
                }
            }
            _ = v60
            _ = v61
        } else {
            var v65 uint8 = 2
            var v66 uint64 = 1
            v65 -= v65
            _ = v65
            _ = v66
        }
        for fi67 := 0; fi67 < 13; fi67++ {
            if (2 < 2) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v3 -= v3
                fmt.Println(v1, fi67, fi67)
            }
            fmt.Println(fi67)
            _ = fi67
        }
        v2 = append(v2, "a")
    } else {
        if (!true) {
            if len(v2) > 0 {
                v2[0] = ""
            }
        } else {
            v1 -= v1
            v1 = v1
            v1 = (v1 | (v1 * v1))
        }
        v1 = (v1 % (((v1 ^ -1)) | 1))
    }
    for fi68 := 0; fi68 < 9; fi68++ {
        var v69 uint = 883
        switch (fi68 / ((fi68) | 1)) {
        case 2147483647:
            v70 := GenS4{F0: v1}
            _ = v70
        case 128:
            v71 := GenS10{F0: false, F1: 42, F2: 2147483648}
            _ = v71
        case 42:
            var v72 uint16 = 65535
            _ = v72
        }
        _ = fi68
        _ = v69
    }
    switch -2147483648 {
    case -65:
        v1 = (v1 | (v1 * -628))
    case 4294967295:
        var v73 uint8 = 0
        _ = v73
    case 128:
        v1 -= (^(v1))
    default:
        var v74 float32 = v3
        _ = v74
    }
    if false {
        v75 := GenS10{F0: false, F1: 127, F2: 42}
        _ = v75
    } else {
        v1 = (v1 % ((-1) | 1))
        v3 -= (v3 * float32(0.1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v56
}
