package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 int16; F2 uint64 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64; F1 int; F2 bool }

func main() {
    var v1 uint = 141
    var v2 uint16 = 468
    var v3 float64 = (((float64((((float32(0.1) + float32(float64((1.5 / 0.1)))) * 1e10) + float32(0.1))) - (1e10 * (0.1 * 1.5))) * 1e10) / 1.0)
    v4 := []int{}
    var v5 int16 = -32768
    switch 32767 {
    case 2147483647:
        _ = len(v4)
    case -2147483648:
        var v6 string = ("hello" + "hello")
        _ = v6
    case 9223372036854775807:
        var v7 int64 = int64(v5)
        _ = v7
    default:
        var v8 int64 = -1
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: 458, F2: v2}
        if (false || true) {
            if true {
                v11 := GenS1{F0: 382, F1: 820, F2: (v2 % ((656) | 1))}
                v5 = v5
                v1 -= v1
                _ = v11
            } else {
                v12 := GenS2{F0: (fi9 >> (((v1 & (v1 ^ uint(v2)))) & 63))}
                _ = v12
            }
            var v13 bool = false
            v1 -= 42
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    switch 999 {
    case 32767:
        v3 = (-2.25 / (0.1 / 1.5))
    case 128:
        var v14 int = int(v2)
        _ = v14
    case 0:
        if (true && true) {
            v5 = (-(v5))
        } else {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint32 = 1
                for fi17 := 0; fi17 < 10; fi17++ {
                    _ = len(v4)
                    switch "a" {
                    case "foo bar":
                        switch "foo bar" {
                        case "hello":
                            v18 := GenS2{F0: (-(fi17))}
                            _ = v18
                        case "foo bar":
                            v19 := GenS1{F0: 4294967295, F1: 459, F2: v2}
                            _ = v19
                        default:
                            var v20 string = "a"
                            _ = v20
                        }
                    default:
                        if ((false || false) && (float64(-2.25) <= float64(3.14159))) {
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = (^(fi17))
                                }
                                v21 := GenS3{F0: 245}
                                var v22 float32 = ((0.1 + (float32(3.14159) / (float32(v3) / ((((0.1 + -2.25) / (((0.0 - (0.0 * float32(((float64(1e10) * (((((v3 / (float64(1e10) / 0.1)) - v3) / 0.1) * -2.25) / (v3 + float64(-2.25)))) * v3)))) / -2.25) - -2.25)) + 3.14159) - 0.1)))) + (1e10 / float32(v3)))
                                _ = v21
                                _ = v22
                            } else {
                                v23 := GenS4{F0: (v1 + (v1 | 255)), F1: 0, F2: 42}
                                fmt.Println(v16, v3)
                                v24 := GenS3{F0: v2}
                                _ = v23
                                _ = v24
                            }
                            v25 := GenS5{F0: "a"}
                            _ = v25
                        } else {
                            v26 := GenS4{F0: (v1 / ((v1) | 1)), F1: 0, F2: 128}
                            v27 := GenS2{F0: (fi15 ^ 42)}
                            _ = v26
                            _ = v27
                        }
                        var v28 int = -1
                        v29 := GenS6{F0: 839, F1: true}
                        _ = v28
                        _ = v29
                    }
                    v3 -= v3
                    var v30 uint8 = 42
                    v1 = v1
                    _ = fi17
                    _ = v30
                }
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 float64 = 0.0
                    v16 = v16
                    fmt.Println(v2)
                    var v33 uint32 = 65535
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = fi15
                _ = v16
            }
        }
    }
    if (false || (true && ((105 == 1) || (87 <= -80)))) {
        for fi34 := 0; fi34 < 7; fi34++ {
            if (((v3 - (v3 - 1e10)) + 1e10) == float64(0.1)) {
                v5 -= 1000
                var v35 uint64 = 4294967295
                _ = v35
            }
            switch fi34 {
            case 369:
                fmt.Println(v5, fi34)
            default:
                for fi36 := 0; fi36 < 6; fi36++ {
                    v5 -= (-(v5))
                    _ = fi36
                }
            }
            var v37 float32 = 3.14159
            _ = fi34
            _ = v37
        }
    }
    var v38 uint64 = 42
    v3 = (3.14159 * v3)
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 int64 = 329
        if (("" + "foo bar") < "hello") {
            switch -828 {
            case 2147483647:
                v1 -= 42
            case 42:
                for fi41 := 0; fi41 < 6; fi41++ {
                    v42 := GenS3{F0: v2}
                    v2 = 32767
                    _ = fi41
                    _ = v42
                }
            case 4294967295:
                var v43 int64 = (v40 | 0)
                _ = v43
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        _ = fi39
        _ = v40
    }
    var v44 int64 = -1
    if (!false) {
        if false {
            if (!(!true)) {
                _ = len(v4)
                if (uint(v2) != (v1 | v1)) {
                    var v45 int16 = v5
                    if len(v4) > 0 {
                        v4[0] = -1
                    }
                    _ = v45
                }
                if (2147483647 < 886) {
                    switch "hello" {
                    case "":
                        fmt.Println(v38, v2, v5)
                    case "a":
                        if (!false) {
                            switch -259 {
                            case 4294967295:
                                if false {
                                    fmt.Println(v3, v44)
                                    for fi46 := 0; fi46 < 3; fi46++ {
                                        v44 = (v44 ^ 9223372036854775807)
                                        _ = len(v4)
                                        _ = fi46
                                    }
                                    v4 = append(v4, 127)
                                }
                            case 0:
                                v2 -= 1
                            case -9223372036854775808:
                                v47 := GenS7{F0: 47}
                                _ = v47
                            }
                            v48 := GenS8{F0: v1, F1: 1, F2: v38}
                            _ = v48
                        }
                    case "hello":
                        if (1000 >= 32767) {
                            switch ("a" + "") {
                            case "\n\t":
                                fmt.Println(v44)
                            case "hello":
                                v38 = (^(v38))
                            case "":
                                v49 := GenS9{F0: ("hello" + ("\n\t" + "foo bar")), F1: (v38 * (^(v38)))}
                                _ = v49
                            }
                        } else {
                            if true {
                                var v50 string = (("" + "\n\t") + "\n\t")
                                _ = v50
                            } else {
                                var v51 uint32 = 4294967295
                                fmt.Println(v51)
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 19; fi52++ {
                                v2 = v2
                                if false {
                                    fmt.Println(v2, v2, v3)
                                }
                                v53 := GenS10{F0: false}
                                _ = fi52
                                _ = v53
                            }
                        }
                    }
                    v54 := GenS11{F0: false}
                    switch 32768 {
                    case 581:
                        v55 := GenS8{F0: (^(v1)), F1: (v5 / ((v5) | 1)), F2: 0}
                        _ = v55
                    case 9223372036854775807:
                        var v56 float64 = (v3 - (float64(0.0) + float64(0.1)))
                        _ = v56
                    default:
                        v44 = (v44 << ((v1) & 63))
                    }
                    _ = v54
                }
            }
            v57 := GenS12{F0: (^(v38)), F1: 230, F2: false}
            if false {
                switch ("hello" + "a") {
                case "foo bar":
                    var v58 uint16 = (v2 >> ((1) & 63))
                    _ = v58
                default:
                    v2 += v2
                }
                if (!true) {
                    if true {
                        var v59 float64 = v3
                        if (28 < 255) {
                            var v60 uint32 = 0
                            _ = v60
                        } else {
                            v1 -= v1
                            v61 := GenS10{F0: false}
                            if false {
                                var v62 string = (("a" + ("hello" + "a")) + ("a" + ("a" + (("\n\t" + ("a" + "a")) + "hello"))))
                                _ = v62
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    v1 += (v1 - v1)
                                case "hello":
                                    v44 -= -9223372036854775808
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                            _ = v61
                        }
                        _ = v59
                    }
                }
            }
            _ = v57
        }
    }
    if len(v4) > 0 {
        v4[0] = 997
    }
    if (!(true || ((127 < 455) && false))) {
        v63 := GenS1{F0: 18446744073709551615, F1: (v38 - v38), F2: 597}
        _ = v63
    }
    v1 += (v1 / ((987) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v5
}
