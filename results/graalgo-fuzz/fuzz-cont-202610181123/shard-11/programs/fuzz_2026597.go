package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 string; F1 string; F2 uint64 }
type GenS7 struct { F0 bool }
func genFn0() uint8 {
    if false {
        var v4 uint16 = 156
        _ = v4
    }
    v5 := []int{-2147483648, 4294967295}
    _ = v5
    return 1
}

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 2
    var v3 uint = 4294967295
    if (v2 <= (^(v2))) {
        var v6 float64 = v1
        var v7 int32 = 0
        var v8 uint8 = 88
        _ = v6
        _ = v7
        _ = v8
    }
    v2 = v2
    if true {
        v9 := []int{}
        switch ("" + ("a" + "foo bar")) {
        case "\n\t":
            var v10 int32 = -2147483648
            _ = v10
        case "hello":
            v3 += v3
        }
        _ = v9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int16 = 285
        switch fi11 {
        case 4294967295:
            v13 := []int{32767, (fi11 + (fi11 >> ((v3) & 63))), fi11, 897}
            _ = v13
        default:
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 int32 = 1
                v3 = v3
                _ = fi14
                _ = v15
            }
        }
        _ = fi11
        _ = v12
    }
    v16 := GenS1{F0: 2147483647, F1: 1, F2: 599}
    if (!false) {
        v17 := GenS2{F0: "\n\t", F1: 42}
        fmt.Println(v1)
        switch "foo bar" {
        case "\n\t":
            if true {
                for fi18 := 0; fi18 < 5; fi18++ {
                    v1 += (((v1 + v1) - 0.0) - 0.0)
                    var v19 int8 = 127
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v3 = 0
                        v21 := []int{fi18, (fi20 % ((fi20) | 1)), 42, -9223372036854775808}
                        _ = fi20
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
                switch 127 {
                case -180:
                    for fi22 := 0; fi22 < 5; fi22++ {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            var v24 uint = 42
                            switch "\n\t" {
                            case "\n\t":
                                v1 -= 1.5
                            case "hello":
                                var v25 float32 = 0.0
                                _ = v25
                            case "":
                                var v26 float32 = (((1.5 - 0.0) * 0.0) / 1.0)
                                _ = v26
                            default:
                                var v27 uint8 = genFn0()
                                _ = v27
                            }
                            var v28 int8 = -1
                            _ = fi23
                            _ = v24
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 12; fi29++ {
                            v2 += (v2 << (((v3 - 276)) & 63))
                            v3 -= (v3 + (v3 >> (((^(v3))) & 63)))
                            v30 := []int{(fi22 * 580), (fi22 * 2147483647)}
                            _ = fi29
                            _ = v30
                        }
                        fmt.Println(v1, v1)
                        _ = fi22
                    }
                case 32768:
                    for fi31 := 0; fi31 < 15; fi31++ {
                        if (false || false) {
                            var v32 uint8 = 2
                            _ = v32
                        } else {
                            fmt.Println(v1)
                        }
                        _ = fi31
                    }
                }
            }
        default:
            var v33 uint16 = 65535
            _ = v33
        }
        _ = v17
    } else {
        switch "hello" {
        case "foo bar":
            switch 128 {
            case 4294967295:
                for fi34 := 0; fi34 < 20; fi34++ {
                    switch "a" {
                    case "":
                        if true {
                            fmt.Println(v1, v2)
                            v3 -= 247
                            switch "a" {
                            case "a":
                                v3 -= (v3 * 2)
                            case "\n\t":
                                v1 = (v1 - 0.0)
                            default:
                                v3 = 1
                            }
                        } else {
                            v35 := []int{127}
                            _ = v35
                        }
                    }
                    _ = fi34
                }
            case 127:
                v1 -= v1
            default:
                v36 := GenS3{F0: 4294967295}
                _ = v36
            }
        }
    }
    v37 := GenS4{F0: (v3 << ((18446744073709551615) & 63))}
    v1 = (v1 - 1.5)
    for fi38 := 0; fi38 < 16; fi38++ {
        v39 := GenS5{F0: 107, F1: "a"}
        var v40 uint8 = 1
        if true {
            for fi41 := 0; fi41 < 14; fi41++ {
                fmt.Println(fi41)
                v42 := []int{int(v3)}
                _ = fi41
                _ = v42
            }
        } else {
            v2 = (v2 % ((1) | 1))
            var v43 uint = (^(v3))
            _ = v43
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    var v44 uint32 = 1
    for fi45 := 0; fi45 < 2; fi45++ {
        v46 := GenS4{F0: 1}
        var v47 uint = v3
        _ = fi45
        _ = v46
        _ = v47
    }
    var v48 uint8 = 80
    switch "" {
    case "\n\t":
        v49 := GenS4{F0: v3}
        _ = v49
    case "foo bar":
        fmt.Println(v44, v48, v44)
    case "a":
        v2 = (v2 - 2)
    default:
        fmt.Println(v44)
    }
    switch "" {
    case "":
        v50 := GenS1{F0: 0, F1: -244, F2: uint(v44)}
        _ = v50
    }
    for fi51 := 0; fi51 < 9; fi51++ {
        v48 = (v48 + (v48 + (v48 % (((v48 >> ((329) & 63))) | 1))))
        if true {
            v44 = (^(v44))
        } else {
            var v52 int16 = -998
            if (1.5 >= (v1 + (v1 - v1))) {
                v1 = v1
            } else {
                switch "hello" {
                case "\n\t":
                    v53 := []int{fi51}
                    _ = v53
                case "foo bar":
                    v52 += (v52 + int16(v2))
                default:
                    if (4294967295 >= (v2 * 2)) {
                        for fi54 := 0; fi54 < 17; fi54++ {
                            if (((v48 | (^(v48))) == (v48 % ((v48) | 1))) || true) {
                                v3 += 0
                                v55 := GenS6{F0: ("a" + "foo bar"), F1: "", F2: 2}
                                fmt.Println(fi51, v48, fi54)
                                _ = v55
                            } else {
                                if (!(v1 >= (((v1 * (((v1 / 1.0) - float64(3.14159)) * v1)) / 1e10) + v1))) {
                                    v56 := GenS6{F0: (("foo bar" + "\n\t") + (("hello" + "hello") + "hello")), F1: "foo bar", F2: v2}
                                    v57 := GenS3{F0: 32767}
                                    _ = v56
                                    _ = v57
                                } else {
                                    v3 -= (v3 / (((v3 ^ (v3 - (v3 - (v3 << ((0) & 63)))))) | 1))
                                    var v58 int16 = (^(v52))
                                    var v59 int32 = 927
                                    _ = v58
                                    _ = v59
                                }
                                v44 -= v44
                                var v60 float32 = 0.1
                                _ = v60
                            }
                            fmt.Println(v48, v3, v3)
                            _ = fi54
                        }
                    } else {
                        v61 := GenS6{F0: ("\n\t" + "\n\t"), F1: ("a" + "a"), F2: v2}
                        v48 += v48
                        fmt.Println(v1, fi51, v3)
                        _ = v61
                    }
                }
                switch 1 {
                case 32768:
                    var v62 uint = (v3 * 65535)
                    _ = v62
                default:
                    fmt.Println(v48, v3, v1)
                }
            }
            var v63 uint = (v3 << ((312) & 63))
            _ = v52
            _ = v63
        }
        v64 := GenS7{F0: false}
        _ = fi51
        _ = v64
    }
    var v65 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v44
    _ = v48
    _ = v65
}
