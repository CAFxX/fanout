package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 uint16; F2 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 uint = 458
    var v2 int8 = 42
    v3 := []int{}
    var v4 int16 = 1000
    var v5 uint8 = 220
    v6 := GenS1{F0: 2}
    v1 += (^(v1))
    var v7 float32 = -2.25
    var v8 int32 = 42
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            switch "" {
            case "\n\t":
                if true {
                    if (!false) {
                        v2 = v2
                    } else {
                        if (false && (!("hello" != ("a" + "hello")))) {
                            var v10 uint8 = v5
                            fmt.Println(v8)
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                v12 := GenS1{F0: 42}
                                var v13 float32 = ((v7 * 0.0) * v7)
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                            v14 := GenS2{F0: v8, F1: v1}
                            var v15 uint = 951
                            _ = v14
                            _ = v15
                        }
                        if len(v3) > 0 {
                            v3[0] = fi9
                        }
                        v8 = (v8 | v8)
                    }
                    fmt.Println(v2, v7)
                    if true {
                        var v16 int64 = 0
                        _ = v16
                    } else {
                        var v17 bool = false
                        _ = v17
                    }
                } else {
                    v1 = v1
                    fmt.Println(v8, fi9, v2)
                }
            case "a":
                v5 -= v5
            default:
                fmt.Println(fi9)
            }
            v8 += (v8 % (((v8 % ((v8) | 1))) | 1))
            _ = fi9
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            if (false && false) {
                if ((false || (true && true)) && true) {
                    var v19 int32 = (v8 * -1)
                    v2 -= (v2 << ((v1) & 63))
                    _ = v19
                } else {
                    fmt.Println(v5, v8, v8)
                }
                if false {
                    var v20 float32 = (v7 * v7)
                    fmt.Println(v1, v7)
                    _ = v20
                } else {
                    switch (^(fi18)) {
                    case -9223372036854775808:
                        var v21 string = "foo bar"
                        _ = v21
                    default:
                        v5 = v5
                    }
                    var v22 uint32 = 1
                    _ = v22
                }
            }
            v1 = v1
            _ = fi18
        }
    } else {
        if false {
            for fi23 := 0; fi23 < 7; fi23++ {
                fmt.Println(v7)
                var v24 float32 = v7
                _ = fi23
                _ = v24
            }
        } else {
            v2 -= v2
            if false {
                var v25 uint8 = (v5 % ((v5) | 1))
                for fi26 := 0; fi26 < 12; fi26++ {
                    v25 -= v5
                    var v27 uint8 = 255
                    v28 := GenS1{F0: 525}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            } else {
                v29 := GenS3{F0: 42, F1: 0}
                var v30 uint32 = 2147483648
                _ = v29
                _ = v30
            }
            if false {
                if (((("a" + ("foo bar" + "\n\t")) + "\n\t") + ("\n\t" + "a")) < "hello") {
                    var v31 int = -1
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v2 = v2
                        v5 += 127
                        _ = fi32
                    }
                    var v33 uint = (v1 ^ 355)
                    _ = v31
                    _ = v33
                } else {
                    var v34 int64 = -9223372036854775808
                    _ = v34
                }
                v35 := GenS2{F0: (v8 % ((0) | 1)), F1: (v1 / ((v1) | 1))}
                switch "a" {
                case "\n\t":
                    fmt.Println(v5, v8)
                case "foo bar":
                    if true {
                        v8 += v8
                    }
                case "":
                    v36 := GenS1{F0: 1}
                    _ = v36
                default:
                    if (!true) {
                        v4 = (-(v4))
                    } else {
                        var v37 bool = (42 < 42)
                        _ = v37
                    }
                }
                _ = v35
            } else {
                _ = len(v3)
            }
        }
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    v5 -= v5
    v38 := GenS4{F0: 65535, F1: 9223372036854775807, F2: 855}
    for fi39 := 0; fi39 < 15; fi39++ {
        if (true && true) {
            switch (-(fi39)) {
            case 0:
                v40 := GenS1{F0: 18446744073709551615}
                _ = v40
            case 42:
                if true {
                    var v41 string = "a"
                    v3 = append(v3, fi39)
                    _ = v41
                } else {
                    var v42 int32 = (v8 << ((v1) & 63))
                    v2 = v2
                    v5 = v5
                    _ = v42
                }
            case 32768:
                v5 -= (v5 - (v5 | uint8(v1)))
            default:
                var v43 float32 = 1e10
                _ = v43
            }
        }
        switch ("\n\t" + ((("a" + "\n\t") + ("foo bar" + "a")) + ("foo bar" + "a"))) {
        case "hello":
            v44 := GenS5{F0: ("foo bar" + "\n\t")}
            _ = v44
        case "foo bar":
            fmt.Println(v5)
        case "\n\t":
            if ((1 <= 0) && ((723 >= 4294967295) || (65535 < 255))) {
                var v45 int64 = -1
                var v46 int = (fi39 / ((127) | 1))
                _ = v45
                _ = v46
            }
        }
        v7 = ((v7 - (1e10 * v7)) + v7)
        _ = fi39
    }
    v8 = -1
    for fi47 := 0; fi47 < 18; fi47++ {
        v4 = v4
        v48 := GenS6{F0: v8, F1: 65535, F2: 18446744073709551615}
        var v49 uint8 = (v5 % ((v5) | 1))
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 6; fi50++ {
        fmt.Println(v1, v8)
        v51 := GenS6{F0: v8, F1: 117, F2: 65535}
        _ = fi50
        _ = v51
    }
    for fi52 := 0; fi52 < 14; fi52++ {
        if true {
            v2 = v2
            v8 -= (v8 ^ (v8 << ((708) & 63)))
            switch (fi52 | 4294967295) {
            case -688:
                if true {
                    if ("" == "a") {
                        switch "hello" {
                        case "a":
                            for fi53 := 0; fi53 < 7; fi53++ {
                                if ((2 < 85) && ((v5 * v5) != (v5 + (v5 % (((v5 & (v5 - (^(v5))))) | 1))))) {
                                    v5 -= (^(v5))
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            switch "foo bar" {
                                            case "a":
                                                var v54 int64 = 415
                                                _ = v54
                                            }
                                            if ((false || false) && (int32(v2) >= (v8 << ((v1) & 63)))) {
                                                v8 -= (v8 % ((v8) | 1))
                                            } else {
                                                v3 = append(v3, fi52)
                                                var v55 uint64 = 0
                                                _ = v55
                                            }
                                        } else {
                                            var v56 string = ("\n\t" + ("\n\t" + ""))
                                            v7 = 3.14159
                                            _ = v56
                                        }
                                    case "hello":
                                        fmt.Println(fi53)
                                    }
                                    v57 := GenS6{F0: (v8 % (((v8 - (v8 >> ((18446744073709551615) & 63)))) | 1)), F1: 65535, F2: (^(v1))}
                                    _ = v57
                                }
                                _ = fi53
                            }
                        case "hello":
                            v8 = (v8 - v8)
                        case "\n\t":
                            v1 += (v1 << ((255) & 63))
                        default:
                            v58 := GenS7{F0: 2}
                            _ = v58
                        }
                    }
                }
            case -1:
                var v59 uint = v1
                _ = v59
            case 32768:
                if (fi52 > (fi52 | fi52)) {
                    var v60 uint = (v1 << ((v1) & 63))
                    v60 = 255
                    _ = v60
                }
            }
        }
        _ = fi52
    }
    v61 := GenS8{F0: (^(v5))}
    v3 = append(v3, -314)
    var v62 uint8 = v5
    fmt.Println(v8, v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v7
    _ = v8
}
