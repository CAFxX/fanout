package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 uint8; F2 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32; F1 bool; F2 int32 }
func genFn0(p0 uint, p1 string) int32 {
    if false {
        p0 += p0
        if true {
            fmt.Println(p0, p1)
        }
        v5 := []int{9223372036854775807}
        _ = v5
    }
    v6 := GenS1{F0: 2, F1: -1}
    fmt.Println(p0)
    switch 32767 {
    case -1:
        p0 += (p0 * 18446744073709551615)
    case 42:
        if false {
            v7 := []int{-625, -9223372036854775808, 4294967295}
            if false {
                v8 := GenS2{F0: 9223372036854775807, F1: -1}
                v9 := GenS2{F0: -1, F1: 2147483647}
                _ = v8
                _ = v9
            } else {
                v7 = append(v7, 32767)
                switch 9223372036854775807 {
                case 0:
                    var v10 bool = true
                    _ = v10
                default:
                    v11 := GenS2{F0: 0, F1: 42}
                    _ = v11
                }
                p0 = p0
            }
            _ = v7
        } else {
            fmt.Println(p1, p1)
            if false {
                v12 := GenS3{F0: 1, F1: p0}
                p0 = (p0 * p0)
                v13 := []int{-483, 127, 42, -9223372036854775808}
                _ = v12
                _ = v13
            }
            fmt.Println(p0)
        }
    default:
        switch ("\n\t" + p1) {
        case "hello":
            v14 := GenS4{F0: 965}
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 11; fi15++ {
                if ((^(p0)) < p0) {
                    v16 := GenS5{F0: 42, F1: 414}
                    fmt.Println(p1)
                    _ = v16
                }
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 float64 = -2.25
                    v19 := []int{728, (fi15 & 42)}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                switch fi15 {
                case 127:
                    switch ((p1 + (p1 + p1)) + p1) {
                    case "hello":
                        var v20 int64 = -15
                        _ = v20
                    case "":
                        switch (("hello" + p1) + p1) {
                        case "":
                            fmt.Println(fi15, p0, p1)
                        case "\n\t":
                            for fi21 := 0; fi21 < 4; fi21++ {
                                v22 := []int{fi21, int(p0)}
                                _ = fi21
                                _ = v22
                            }
                        }
                    }
                default:
                    var v23 bool = true
                    _ = v23
                }
                _ = fi15
            }
        default:
            p0 = (p0 << ((p0) & 63))
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 38
}
func genFn1(p0 int16, p1 int8, p2 uint16) uint {
    var v24 string = "hello"
    switch -283 {
    case 184:
        var v25 uint8 = 2
        _ = v25
    case 4294967295:
        v26 := GenS6{F0: ((true && ((193 != 2) || true)) || (((3.14159 * (1.5 - -2.25)) / (1.5 * ((0.1 + (0.1 - (float64((0.0 / 1.0)) * 1e10))) / (-2.25 - 0.0)))) != 1e10)), F1: 18446744073709551615}
        _ = v26
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := GenS3{F0: 201, F1: 18446744073709551615}
            _ = fi27
            _ = v28
        }
    }
    switch 2147483647 {
    case 42:
        v29 := GenS1{F0: 42, F1: -1}
        _ = v29
    case 2147483647:
        var v30 uint = 1
        _ = v30
    default:
        for fi31 := 0; fi31 < 18; fi31++ {
            v24 = (("hello" + v24) + ((v24 + v24) + ((((v24 + v24) + (v24 + (((v24 + ((v24 + "foo bar") + "")) + ("" + v24)) + (("a" + ("\n\t" + v24)) + (v24 + ((v24 + "\n\t") + v24)))))) + "\n\t") + "foo bar")))
            var v32 float32 = ((-2.25 * 0.1) + 1.5)
            _ = fi31
            _ = v32
        }
    }
    v33 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v33
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 uint64 = 791
    var v4 float64 = 1.5
    for fi34 := 0; fi34 < 2; fi34++ {
        v3 += (v3 << ((65535) & 63))
        _ = fi34
    }
    var v35 uint32 = 42
    v35 -= v35
    if true {
        switch 42 {
        case 42:
            var v36 uint = 560
            _ = v36
        default:
            delete(v1, v35)
        }
    } else {
        var v37 int16 = -900
        v4 += v4
        _ = v37
    }
    v4 += v4
    v35 += v35
    var v38 uint = genFn1(1000, 1, 980)
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 uint = 18446744073709551615
        _ = fi39
        _ = v40
    }
    switch 42 {
    case 127:
        fmt.Println(v3)
    case 2147483647:
        fmt.Println(v4, v4)
    default:
        v3 -= (v3 & 42)
    }
    v41 := []int{127, 9223372036854775807, 32768, 2147483647}
    _ = len(v41)
    if (!true) {
        if (-128 > -56) {
            fmt.Println(v38, v35, v4)
            switch "foo bar" {
            case "foo bar":
                if (false && ("a" != ("a" + (("foo bar" + (("foo bar" + (("hello" + "") + "")) + "foo bar")) + "hello")))) {
                    v42 := GenS7{F0: -1}
                    var v43 uint16 = 0
                    var v44 uint32 = uint32(v3)
                    _ = v42
                    _ = v43
                    _ = v44
                }
            }
        }
        v41 = append(v41, -846)
        switch 572 {
        case 42:
            if ((^(v38)) <= 742) {
                v3 = 0
                v45 := GenS6{F0: true, F1: 42}
                _ = v45
            } else {
                for fi46 := 0; fi46 < 10; fi46++ {
                    v3 = 18446744073709551615
                    _ = fi46
                }
            }
        case -9223372036854775808:
            v4 = 0.1
        default:
            v35 = v35
        }
    }
    if (((v3 / ((v3) | 1)) > (v3 ^ (v3 + v3))) || ((v38 << ((v38) & 63)) <= (^(v38)))) {
        if (!(((1e10 - float64((float32(v4) + (0.0 - 1.5)))) == v4) || (!(0 != -9223372036854775808)))) {
            v47 := GenS3{F0: 42, F1: (v38 ^ (^(v38)))}
            v48 := GenS3{F0: 42, F1: genFn1(718, 100, 32767)}
            _ = v47
            _ = v48
        } else {
            var v49 float64 = float64((-2.25 + 0.0))
            for fi50 := 0; fi50 < 12; fi50++ {
                for fi51 := 0; fi51 < 3; fi51++ {
                    switch 612 {
                    case 506:
                        delete(v1, v35)
                        if (int32(v3) != -2147483648) {
                            v4 -= 1e10
                            var v52 int = fi51
                            var v53 uint8 = 239
                            _ = v52
                            _ = v53
                        } else {
                            if len(v41) > 0 {
                                v41[0] = (fi51 - 295)
                            }
                            v49 = v4
                        }
                    case -183:
                        if ("\n\t" <= (((("hello" + ("hello" + "\n\t")) + "foo bar") + "hello") + "\n\t")) {
                            fmt.Println(fi50, v49, fi51)
                            fmt.Println(v49, v38)
                        }
                    }
                    switch "" {
                    case "":
                        if true {
                            v49 = (v4 + float64(3.14159))
                        } else {
                            if (v3 < v3) {
                                fmt.Println(v3, v35, v35)
                            } else {
                                fmt.Println(v35, fi50)
                                var v54 uint64 = (v3 & 1)
                                v38 = 1
                                _ = v54
                            }
                            if (-9223372036854775808 != 1) {
                                var v55 string = "foo bar"
                                _ = v55
                            }
                        }
                    case "\n\t":
                        if false {
                            v56 := GenS8{F0: (genFn0((v38 / ((uint(v35)) | 1)), "hello") < -2147483648), F1: 127, F2: fi50}
                            v57 := GenS7{F0: -1}
                            _ = v56
                            _ = v57
                        } else {
                            var v58 int32 = 42
                            if true {
                                v59 := GenS2{F0: (fi50 >> ((v38) & 63)), F1: genFn0(389, "\n\t")}
                                _ = v59
                            }
                            _ = v58
                        }
                    case "foo bar":
                        var v60 int64 = -9223372036854775808
                        _ = v60
                    }
                    _ = fi51
                }
                if true {
                    if true {
                        if len(v41) > 0 {
                            v41[0] = fi50
                        }
                        switch ("hello" + "") {
                        case "hello":
                            v38 += genFn1(-32768, 42, 153)
                        }
                        var v61 int8 = 93
                        _ = v61
                    }
                    if ("\n\t" >= ("" + ("a" + "foo bar"))) {
                        v62 := v2[762]
                        _ = v62
                    } else {
                        var v63 float32 = (0.0 + (float32(float64(3.14159)) + (0.1 / 1e10)))
                        _ = v63
                    }
                    switch (fi50 >> (((v38 * (v38 >> ((255) & 63)))) & 63)) {
                    case 852:
                        switch (fi50 / (((fi50 / (((-(fi50))) | 1))) | 1)) {
                        case -9223372036854775808:
                            fmt.Println(v38)
                        default:
                            if (234 >= 128) {
                                v64 := GenS9{F0: 1}
                                _ = v64
                            }
                        }
                    default:
                        fmt.Println(v35, v4)
                    }
                } else {
                    v41 = append(v41, (fi50 & 128))
                    v65 := GenS10{F0: (v35 + v35), F1: (42 != 32767), F2: genFn0(v38, ("foo bar" + ("hello" + ((("hello" + "hello") + ("a" + (((("hello" + ("hello" + "a")) + (("\n\t" + "a") + (("" + "") + ("\n\t" + ("" + ("hello" + ("hello" + ("hello" + "")))))))) + (("" + "a") + "hello")) + "foo bar"))) + "hello"))))}
                    _ = v65
                }
                _ = fi50
            }
            var v66 uint64 = 167
            _ = v49
            _ = v66
        }
        var v67 uint16 = 1
        _ = v67
    } else {
        if (true && false) {
            v68 := GenS3{F0: 1, F1: (v38 - uint(v3))}
            var v69 string = "hello"
            _ = v68
            _ = v69
        } else {
            fmt.Println(v3, v35)
        }
    }
    v3 = v3
    if (423 <= 127) {
        var v70 uint = v38
        var v71 int8 = -1
        var v72 uint = 18446744073709551615
        _ = v70
        _ = v71
        _ = v72
    }
    var v73 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v41
    _ = v73
}
