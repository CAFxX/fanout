package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint; F2 bool }
type GenS7 struct { F0 uint32; F1 uint; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 int64 }
func genFn0(p0 string, p1 uint16) uint32 {
    v5 := GenS1{F0: 0}
    var v6 int = 32767
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint32) int16 {
    var v7 float64 = (0.1 * 0.0)
    v8 := []int{32767, -9223372036854775808, 2147483647, -9223372036854775808}
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS2{F0: (^(p0))}
        var v11 float64 = -2.25
        _ = fi9
        _ = v10
        _ = v11
    }
    if ((!((0 >= 18446744073709551615) || (!true))) && true) {
        v12 := GenS3{F0: 42, F1: 372, F2: p0}
        _ = v12
    } else {
        fmt.Println(p0)
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint64 = 2
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    var v3 float64 = 0.0
    var v4 uint64 = 0
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 string = ("hello" + ("foo bar" + "foo bar"))
        if true {
            var v17 uint64 = v4
            if false {
                v18 := GenS2{F0: 1}
                v3 = 0.1
                var v19 string = v16
                _ = v18
                _ = v19
            }
            _ = v17
        } else {
            if (42 > 42) {
                fmt.Println(v16)
                for fi20 := 0; fi20 < 17; fi20++ {
                    v1 += v1
                    _ = fi20
                }
                if true {
                    for fi21 := 0; fi21 < 8; fi21++ {
                        v3 -= (((float64(1e10) * v3) - v3) + (v3 * 0.0))
                        v22 := []int{(fi21 / ((-815) | 1)), -9223372036854775808, 2147483647}
                        _ = len(v2)
                        _ = fi21
                        _ = v22
                    }
                }
            } else {
                if (2147483647 != 0) {
                    v23 := []int{}
                    var v24 int8 = 1
                    v25 := v2[2]
                    _ = v23
                    _ = v24
                    _ = v25
                }
                fmt.Println(v16, v4)
                if (true || (1 <= 42)) {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        var v27 uint16 = 32767
                        var v28 float32 = 1.5
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    if false {
                        for fi29 := 0; fi29 < 12; fi29++ {
                            var v30 uint8 = v1
                            var v31 uint64 = v4
                            v2[42] = false
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        _ = len(v2)
                    }
                    fmt.Println(v1, v3)
                } else {
                    var v32 bool = (!(!true))
                    _ = v32
                }
            }
            var v33 float32 = (((0.1 - float32(v3)) / 0.1) / 1e10)
            _ = v33
        }
        _ = fi15
        _ = v16
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v34 int8 = 42
        _ = v34
    default:
        switch 836 {
        case -647:
            if false {
                v35 := GenS4{F0: (792 >= 642), F1: (v4 / ((v4) | 1))}
                _ = v35
            }
        default:
            for fi36 := 0; fi36 < 13; fi36++ {
                v37 := GenS4{F0: (!false), F1: 2}
                _ = fi36
                _ = v37
            }
        }
    }
    if ((v1 - v1) >= v1) {
        switch "foo bar" {
        case "\n\t":
            var v38 float64 = 0.1
            _ = v38
        default:
            if true {
                fmt.Println(v4, v3)
                var v39 uint16 = 347
                _ = v39
            }
        }
    } else {
        for fi40 := 0; fi40 < 13; fi40++ {
            for fi41 := 0; fi41 < 4; fi41++ {
                v42 := GenS5{F0: false}
                var v43 string = "hello"
                v44 := GenS6{F0: -1, F1: 65535, F2: (1.5 >= -2.25)}
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
            switch ("\n\t" + ((("hello" + "a") + "a") + "a")) {
            case "foo bar":
                switch "foo bar" {
                case "":
                    v1 += (v1 & v1)
                }
            case "hello":
                for fi45 := 0; fi45 < 18; fi45++ {
                    v4 = 1
                    v2[0] = (!true)
                    if true {
                        if (!false) {
                            v46 := GenS7{F0: 901, F1: 509, F2: 0}
                            v47 := GenS8{F0: 1}
                            _ = v46
                            _ = v47
                        }
                    } else {
                        var v48 uint8 = v1
                        v4 += v4
                        _ = v48
                    }
                    _ = fi45
                }
            }
            _ = fi40
        }
        var v49 uint64 = 18446744073709551615
        var v50 int64 = 9223372036854775807
        _ = v49
        _ = v50
    }
    v3 = v3
    fmt.Println(v4, v1, v1)
    v4 = (^(v4))
    var v51 int16 = 32767
    _ = len(v2)
    v3 = float64(-2.25)
    for fi52 := 0; fi52 < 8; fi52++ {
        var v53 uint64 = (v4 + (^(v4)))
        fmt.Println(v4, v1)
        _ = fi52
        _ = v53
    }
    var v54 uint32 = uint32(v1)
    for fi55 := 0; fi55 < 12; fi55++ {
        fmt.Println(fi55, v1)
        v3 -= 1e10
        _ = fi55
    }
    for fi56 := 0; fi56 < 7; fi56++ {
        var v57 uint8 = v1
        if false {
            switch "a" {
            case "a":
                v58 := GenS5{F0: true}
                _ = v58
            default:
                switch "foo bar" {
                case "\n\t":
                    v51 = (-(v51))
                case "hello":
                    var v59 int16 = (-(v51))
                    _ = v59
                default:
                    fmt.Println(v51)
                }
            }
            v51 += genFn1(42)
        }
        switch "a" {
        case "":
            var v60 int = (fi56 + (fi56 | 0))
            _ = v60
        case "hello":
            for fi61 := 0; fi61 < 11; fi61++ {
                if (573 <= uint(v54)) {
                    v62 := GenS9{F0: 854, F1: 926}
                    _ = v62
                }
                _ = fi61
            }
        }
        _ = fi56
        _ = v57
    }
    for fi63 := 0; fi63 < 1; fi63++ {
        v64 := v2[genFn0("\n\t", 0)]
        _ = fi63
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v54
}
