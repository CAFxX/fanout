package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 bool; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 uint64; F2 string }
type GenS11 struct { F0 int64; F1 bool; F2 int8 }
type GenS12 struct { F0 int64; F1 bool }
func genFn0(p0 bool, p1 int32, p2 uint8) bool {
    switch "" {
    case "hello":
        var v7 uint = 42
        _ = v7
    case "\n\t":
        var v8 float32 = 1.5
        _ = v8
    default:
        var v9 uint8 = p2
        _ = v9
    }
    p1 -= -2147483648
    p0 = (!true)
    p0 = true
    _ = p0
    _ = p1
    _ = p2
    return (0 == 602)
}
func genFn1(p0 uint, p1 int, p2 int) int8 {
    for fi10 := 0; fi10 < 8; fi10++ {
        if true {
            v11 := GenS1{F0: 42, F1: 32767, F2: 2147483647}
            v12 := GenS1{F0: -1, F1: -790, F2: 1}
            _ = v11
            _ = v12
        } else {
            v13 := []int{p1}
            _ = v13
        }
        fmt.Println(fi10, fi10)
        v14 := []int{(-(p1)), (p2 % (((-(p1))) | 1)), (p2 << (((p0 * p0)) & 63)), fi10}
        _ = fi10
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        p2 = fi15
        p2 = 4294967295
        fmt.Println(p2, p0)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    return 84
}

func main() {
    var v1 float32 = (((1e10 + 3.14159) / 1e10) / float32(-2.25))
    var v2 int16 = -1
    v3 := []string{("a" + "hello"), (("\n\t" + "foo bar") + "a")}
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    if (!false) {
        var v16 uint = 776
        _ = v16
    }
    v2 = (v2 % ((135) | 1))
    switch ("hello" + "foo bar") {
    case "hello":
        v17 := GenS2{F0: (109 == 0)}
        _ = v17
    case "\n\t":
        delete(v5, ("hello" + "\n\t"))
    default:
        if (255 != 1) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi18 := 0; fi18 < 10; fi18++ {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                var v19 uint16 = 91
                _ = fi18
                _ = v19
            }
            switch "hello" {
            case "hello":
                if (0 != 65535) {
                    var v20 uint32 = 584
                    var v21 uint = 18446744073709551615
                    v22 := GenS3{F0: (9223372036854775807 >= -1)}
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    if false {
                        var v23 uint64 = 0
                        v1 = (v1 + v1)
                        _ = v23
                    }
                    if true {
                        switch ("" + "foo bar") {
                        case "foo bar":
                            v1 = v1
                        case "":
                            if false {
                                fmt.Println(v2, v1)
                                var v24 uint = 4294967295
                                v25 := GenS3{F0: (false || true)}
                                _ = v24
                                _ = v25
                            }
                        case "a":
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v27 := GenS4{F0: 571}
                                _ = fi26
                                _ = v27
                            }
                        }
                        if (true || false) {
                            var v28 bool = (1 >= 0)
                            _ = v28
                        } else {
                            switch 32768 {
                            case 2147483647:
                                delete(v4, 4294967295)
                            case 32768:
                                for fi29 := 0; fi29 < 8; fi29++ {
                                    var v30 int8 = -43
                                    _ = fi29
                                    _ = v30
                                }
                            case -2147483648:
                                var v31 int32 = 0
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 2; fi32++ {
                                var v33 int32 = 2147483647
                                _ = fi32
                                _ = v33
                            }
                            if (false && (!(830 <= 0))) {
                                if true {
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v1, v1)
                                    default:
                                        switch 2147483647 {
                                        case 128:
                                            switch -212 {
                                            case -883:
                                                fmt.Println(v1, v2)
                                            case -9223372036854775808:
                                                switch ("hello" + ("" + (("" + ("foo bar" + "")) + "\n\t"))) {
                                                case "foo bar":
                                                    if (!false) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if false {
                                                            var v34 int16 = v2
                                                            switch "hello" {
                                                            case "hello":
                                                                v35 := GenS2{F0: (!true)}
                                                                _ = v35
                                                            case "":
                                                                var v36 uint = 1
                                                                _ = v36
                                                            default:
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    switch 2147483647 {
                                                                    case 32768:
                                                                        v1 = ((0.0 + v1) * (v1 + v1))
                                                                    case 32767:
                                                                        switch -156 {
                                                                        case 1:
                                                                            switch 2147483647 {
                                                                            case 0:
                                                                                for fi37 := 0; fi37 < 2; fi37++ {
                                                                                    if false {
                                                                                        var v38 int16 = v34
                                                                                        var v39 uint8 = 2
                                                                                        fmt.Println(v38, v1, v34)
                                                                                        _ = v38
                                                                                        _ = v39
                                                                                    }
                                                                                    _ = fi37
                                                                                }
                                                                            case 1:
                                                                                for fi40 := 0; fi40 < 12; fi40++ {
                                                                                    var v41 float32 = 3.14159
                                                                                    for fi42 := 0; fi42 < 13; fi42++ {
                                                                                        var v43 int8 = -99
                                                                                        v44 := GenS5{F0: 388, F1: "foo bar"}
                                                                                        _ = fi42
                                                                                        _ = v43
                                                                                        _ = v44
                                                                                    }
                                                                                    _ = fi40
                                                                                    _ = v41
                                                                                }
                                                                            }
                                                                        default:
                                                                            v34 = v2
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v34
                                                        } else {
                                                            var v45 uint64 = 179
                                                            switch "a" {
                                                            case "hello":
                                                                fmt.Println(v45, v45)
                                                            case "foo bar":
                                                                v1 = float32(-2.25)
                                                            case "\n\t":
                                                                v3 = append(v3, ((("\n\t" + "") + ("hello" + ("" + "hello"))) + ("foo bar" + ("\n\t" + "\n\t"))))
                                                            default:
                                                                v45 = 198
                                                            }
                                                            for fi46 := 0; fi46 < 2; fi46++ {
                                                                switch (fi46 - fi46) {
                                                                case -1:
                                                                    _ = len(v6)
                                                                }
                                                                _ = fi46
                                                            }
                                                            _ = v45
                                                        }
                                                    }
                                                default:
                                                    v47 := GenS6{F0: (v2 | (v2 ^ v2)), F1: 42}
                                                    _ = v47
                                                }
                                            }
                                        case 2147483647:
                                            v1 = -2.25
                                        case 9223372036854775807:
                                            v48 := GenS3{F0: (4294967295 <= 199)}
                                            _ = v48
                                        default:
                                            var v49 int16 = (v2 - v2)
                                            _ = v49
                                        }
                                    }
                                }
                            }
                        }
                        v1 += float32(-2.25)
                    } else {
                        v1 += ((1.5 - v1) - 3.14159)
                        v1 -= v1
                    }
                    v3 = append(v3, ("a" + "\n\t"))
                }
            }
        } else {
            _ = len(v3)
            var v50 int16 = (-(v2))
            _ = v50
        }
    }
    var v51 float32 = (((v1 + ((1.5 / (((v1 / (((v1 * v1) * float32((3.14159 / (float64(-2.25) + -2.25)))) / float32((0.1 + 1.5)))) - v1) + v1)) * 1.5)) - (((v1 + (-2.25 / (float32(-2.25) / 3.14159))) - v1) * v1)) + (v1 / v1))
    delete(v6, 2147483647)
    for fi52 := 0; fi52 < 8; fi52++ {
        v53 := GenS7{F0: 2}
        _ = fi52
        _ = v53
    }
    var v54 string = ""
    v55 := GenS8{F0: 1, F1: (1 >= -128), F2: v54}
    if (1 != -242) {
        var v56 int16 = -950
        var v57 int16 = (v2 + 1000)
        _ = v56
        _ = v57
    }
    if len(v3) > 0 {
        v3[0] = (v54 + v54)
    }
    v58 := GenS9{F0: -1}
    v59 := GenS10{F0: 127, F1: 61, F2: ((v54 + "foo bar") + "a")}
    v60 := GenS11{F0: 1, F1: ((v2 ^ 1000) > (v2 << ((897) & 63))), F2: -60}
    v61 := GenS12{F0: 0, F1: false}
    v1 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
