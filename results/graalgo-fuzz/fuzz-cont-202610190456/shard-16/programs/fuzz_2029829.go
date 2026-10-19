package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int64; F1 bool; F2 int }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int8; F1 int8; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 uint8; F2 uint }
type GenS7 struct { F0 int64 }
func genFn0(p0 int64, p1 bool, p2 uint8) int8 {
    if true {
        if false {
            v5 := GenS1{F0: -36, F1: (("foo bar" + (("" + "\n\t") + ((("a" + ("\n\t" + ((("hello" + "hello") + ("" + "a")) + "hello"))) + ("a" + ("a" + "foo bar"))) + ("" + "hello")))) + "a")}
            _ = v5
        }
        fmt.Println(p2)
    } else {
        p0 = p0
        if false {
            if false {
                var v6 int32 = 0
                var v7 int8 = -55
                _ = v6
                _ = v7
            }
        } else {
            var v8 float32 = 1.5
            for fi9 := 0; fi9 < 8; fi9++ {
                p2 = 113
                var v10 int32 = -1
                var v11 int8 = 100
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
    }
    fmt.Println(p0, p0)
    v12 := GenS2{F0: -1, F1: p1, F2: 42}
    v13 := GenS2{F0: p0, F1: (p1 && false), F2: -2147483648}
    if p1 {
        v14 := []int{127, 732, 128}
        v15 := GenS2{F0: (p0 + -9223372036854775808), F1: p1, F2: 0}
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 19; fi17++ {
                if false {
                    switch (("\n\t" + "hello") + "a") {
                    case "foo bar":
                        v18 := GenS1{F0: (fi17 + (^(fi16))), F1: "foo bar"}
                        _ = v18
                    case "\n\t":
                        var v19 bool = (float32(-2.25) >= 0.0)
                        _ = v19
                    case "a":
                        var v20 uint32 = 724
                        _ = v20
                    default:
                        v21 := GenS1{F0: (fi16 | 337), F1: "hello"}
                        _ = v21
                    }
                }
                _ = fi17
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                p2 -= p2
                v23 := GenS2{F0: p0, F1: true, F2: 9223372036854775807}
                p2 -= p2
                _ = fi22
                _ = v23
            }
            if ((((^(p0)) <= 42) || false) && p1) {
                var v24 uint = 255
                v25 := GenS3{F0: p2, F1: 2147483647}
                v26 := GenS1{F0: fi16, F1: "foo bar"}
                _ = v24
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 7; fi27++ {
                    if len(v14) > 0 {
                        v14[0] = (fi16 << ((2) & 63))
                    }
                    _ = len(v14)
                    p1 = (34 >= -2147483648)
                    _ = fi27
                }
            }
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint64 = 604
    var v2 uint = uint(v1)
    var v3 uint32 = 1
    var v4 int8 = 1
    if true {
        v1 = 1
        v28 := []int{}
        v4 = 100
        _ = v28
    } else {
        v29 := []int{128, 127, -9223372036854775808, 257}
        for fi30 := 0; fi30 < 14; fi30++ {
            switch "" {
            case "":
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
            default:
                fmt.Println(v1, fi30, v1)
            }
            var v31 int64 = 1
            fmt.Println(v31, v31)
            _ = fi30
            _ = v31
        }
        var v32 int8 = (v4 >> (((^(v2))) & 63))
        _ = v29
        _ = v32
    }
    v33 := []int{}
    if (int16(v4) == 42) {
        var v34 int = -640
        _ = v34
    }
    var v35 uint32 = v3
    var v36 int64 = 1
    if (!(0 >= 0)) {
        switch "" {
        case "hello":
            fmt.Println(v3, v4, v4)
        default:
            var v37 uint64 = (v1 >> ((v2) & 63))
            _ = v37
        }
        v38 := GenS4{F0: (v4 << ((v2) & 63)), F1: v4, F2: ("foo bar" + "hello")}
        _ = v38
    } else {
        _ = len(v33)
    }
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 int32 = 2147483647
        if len(v33) > 0 {
            v33[0] = fi39
        }
        switch ("foo bar" + "") {
        case "foo bar":
            var v41 int = fi39
            _ = v41
        case "\n\t":
            for fi42 := 0; fi42 < 10; fi42++ {
                var v43 uint = (v2 * v2)
                v44 := GenS5{F0: -32768}
                if true {
                    v45 := GenS6{F0: 9223372036854775807, F1: 255, F2: uint(v1)}
                    _ = v45
                } else {
                    v36 = (^(v36))
                    v46 := GenS7{F0: v36}
                    if true {
                        fmt.Println(fi42, v1)
                        switch "\n\t" {
                        case "a":
                            v40 = -527
                        case "":
                            v33 = append(v33, fi39)
                        case "\n\t":
                            _ = len(v33)
                        default:
                            v47 := GenS1{F0: (fi39 << (((v2 + (^(v43)))) & 63)), F1: "\n\t"}
                            _ = v47
                        }
                        v1 = 1
                        fmt.Println(v35)
                    }
                    _ = v46
                }
                _ = fi42
                _ = v43
                _ = v44
            }
        case "":
            for fi48 := 0; fi48 < 4; fi48++ {
                if ((v1 | v1) >= v1) {
                    v1 = (v1 * 854)
                }
                var v49 uint8 = 10
                _ = fi48
                _ = v49
            }
        default:
            v1 -= v1
        }
        _ = fi39
        _ = v40
    }
    for fi50 := 0; fi50 < 14; fi50++ {
        if (true && true) {
            v33 = append(v33, int(v4))
        }
        _ = fi50
    }
    switch 32767 {
    case 9223372036854775807:
        v51 := GenS5{F0: -32768}
        _ = v51
    case -9223372036854775808:
        switch "hello" {
        case "a":
            _ = len(v33)
        case "foo bar":
            v3 = (v3 + v3)
        case "":
            _ = len(v33)
        default:
            for fi52 := 0; fi52 < 14; fi52++ {
                v53 := GenS2{F0: v36, F1: true, F2: (fi52 ^ (fi52 >> (((v2 | (v2 % (((v2 >> ((2) & 63))) | 1)))) & 63)))}
                v54 := GenS1{F0: (fi52 & int(v1)), F1: "\n\t"}
                v33 = append(v33, fi52)
                _ = fi52
                _ = v53
                _ = v54
            }
        }
    case -2147483648:
        var v55 float64 = (1e10 + 3.14159)
        _ = v55
    default:
        var v56 uint = v2
        _ = v56
    }
    switch "hello" {
    case "foo bar":
        var v57 uint16 = 243
        _ = v57
    }
    _ = len(v33)
    switch "\n\t" {
    case "foo bar":
        if true {
            v36 -= 0
        }
    case "hello":
        var v58 float32 = (-2.25 / 1.5)
        _ = v58
    default:
        fmt.Println(v4)
    }
    v4 -= 100
    for fi59 := 0; fi59 < 18; fi59++ {
        if false {
            var v60 uint16 = 389
            _ = v60
        } else {
            fmt.Println(v4)
            v2 = v2
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
}
