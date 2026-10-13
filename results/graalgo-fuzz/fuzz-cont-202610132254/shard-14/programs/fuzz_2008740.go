package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int8 }
func genFn0() string {
    v4 := []int{-2147483648, 1, 4294967295}
    if len(v4) > 0 {
        v4[0] = 1
    }
    if false {
        if (true || ((!(!(false && (true || false)))) || (656 > 2147483647))) {
            if true {
                for fi5 := 0; fi5 < 8; fi5++ {
                    _ = len(v4)
                    var v6 int = fi5
                    _ = fi5
                    _ = v6
                }
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := GenS1{F0: true, F1: fi7, F2: -578}
                    var v9 int8 = 1
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                if true {
                    for fi10 := 0; fi10 < 12; fi10++ {
                        if (53 == 12) {
                            switch (fi10 ^ 42) {
                            case 32767:
                                var v11 uint16 = 1
                                _ = v11
                            }
                        }
                        _ = fi10
                    }
                    switch "a" {
                    case "\n\t":
                        if (!true) {
                            v12 := GenS1{F0: false, F1: 2147483647, F2: 42}
                            switch "\n\t" {
                            case "hello":
                                var v13 uint16 = 65535
                                _ = v13
                            default:
                                v14 := GenS1{F0: false, F1: 0, F2: -32768}
                                _ = v14
                            }
                            if len(v4) > 0 {
                                v4[0] = 128
                            }
                            _ = v12
                        }
                    case "a":
                        for fi15 := 0; fi15 < 5; fi15++ {
                            var v16 float64 = -2.25
                            if ((!true) && true) {
                                v16 += (-2.25 * v16)
                            } else {
                                if (((false && false) && false) || false) {
                                    var v17 float64 = float64(-2.25)
                                    var v18 float64 = ((v16 + 3.14159) * v16)
                                    _ = v17
                                    _ = v18
                                }
                                v19 := GenS2{F0: false, F1: 2147483647}
                                var v20 uint32 = 2147483647
                                _ = v19
                                _ = v20
                            }
                            _ = fi15
                            _ = v16
                        }
                    default:
                        v21 := GenS1{F0: true, F1: 4294967295, F2: 1000}
                        _ = v21
                    }
                    var v22 int16 = -32768
                    _ = v22
                } else {
                    var v23 uint8 = 53
                    switch 32767 {
                    case 2147483647:
                        v23 = 255
                    case 42:
                        var v24 string = "\n\t"
                        _ = v24
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v23
                }
            }
            if false {
                var v25 string = ("\n\t" + "hello")
                _ = v25
            } else {
                v26 := GenS1{F0: (true && (false && false)), F1: -9223372036854775808, F2: 1}
                _ = v26
            }
            var v27 string = "\n\t"
            _ = v27
        }
        var v28 float64 = 0.0
        _ = v28
    }
    _ = v4
    return ""
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 1
    var v3 int64 = -83
    fmt.Println(v3, v2)
    v3 = v3
    v2 -= 1
    var v29 int64 = -1
    v30 := GenS2{F0: (!true), F1: -1}
    for fi31 := 0; fi31 < 14; fi31++ {
        v2 = (v2 | 94)
        _ = fi31
    }
    fmt.Println(v2)
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 string = ""
        var v34 bool = ((v2 % ((102) | 1)) == v2)
        fmt.Println(v34)
        _ = fi32
        _ = v33
        _ = v34
    }
    v2 = v2
    v29 += (v29 << ((430) & 63))
    v35 := GenS1{F0: false, F1: 0, F2: -32768}
    switch 509 {
    case -821:
        var v36 bool = (false && true)
        _ = v36
    default:
        fmt.Println(v3, v29, v3)
    }
    v37 := GenS1{F0: (false && false), F1: 9223372036854775807, F2: -1}
    switch ("hello" + ("hello" + ("\n\t" + (("foo bar" + "\n\t") + "\n\t")))) {
    case "a":
        var v38 int32 = 2147483647
        _ = v38
    case "foo bar":
        var v39 bool = true
        _ = v39
    default:
        if false {
            var v40 uint16 = 2
            if true {
                for fi41 := 0; fi41 < 8; fi41++ {
                    if false {
                        for fi42 := 0; fi42 < 1; fi42++ {
                            if false {
                                var v43 int16 = 32767
                                var v44 int64 = -9223372036854775808
                                var v45 uint16 = 65535
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                var v46 bool = true
                                v47 := GenS3{F0: ("foo bar" + "foo bar")}
                                _ = v46
                                _ = v47
                            }
                            _ = fi42
                        }
                        v48 := GenS4{F0: 2, F1: 42, F2: v2}
                        for fi49 := 0; fi49 < 16; fi49++ {
                            delete(v1, 623)
                            if true {
                                v50 := []int{(-(fi49))}
                                _ = v50
                            } else {
                                v40 = (v40 / (((v40 * uint16(v2))) | 1))
                            }
                            var v51 int8 = -128
                            v51 = -128
                            _ = fi49
                            _ = v51
                        }
                        _ = v48
                    }
                    for fi52 := 0; fi52 < 5; fi52++ {
                        var v53 uint64 = uint64(v2)
                        v53 = 885
                        _ = fi52
                        _ = v53
                    }
                    v3 = v3
                    _ = fi41
                }
                for fi54 := 0; fi54 < 6; fi54++ {
                    var v55 float32 = ((float32(((-2.25 * -2.25) - (-2.25 * (-2.25 - 0.0)))) * 0.1) - float32(0.0))
                    _ = fi54
                    _ = v55
                }
                v2 += v2
            }
            _ = v40
        } else {
            switch "\n\t" {
            case "hello":
                v29 += (v29 ^ (v29 / ((v29) | 1)))
            case "foo bar":
                v56 := GenS1{F0: (2147483647 == 998), F1: 2147483647, F2: 582}
                _ = v56
            }
            for fi57 := 0; fi57 < 11; fi57++ {
                for fi58 := 0; fi58 < 18; fi58++ {
                    var v59 uint = 438
                    delete(v1, 65535)
                    var v60 string = "\n\t"
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                _ = fi57
            }
            if (3.14159 >= (3.14159 * 0.1)) {
                for fi61 := 0; fi61 < 2; fi61++ {
                    for fi62 := 0; fi62 < 11; fi62++ {
                        if false {
                            if false {
                                var v63 int8 = 74
                                var v64 uint8 = v2
                                _ = v63
                                _ = v64
                            }
                            v65 := []int{fi62, 868}
                            _ = v65
                        }
                        v66 := GenS4{F0: 274, F1: 0, F2: v2}
                        _ = fi62
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 20; fi67++ {
                        v3 = v29
                        _ = fi67
                    }
                    v29 = (v3 | (v3 << ((255) & 63)))
                    _ = fi61
                }
                fmt.Println(v2, v2, v29)
                v2 -= v2
            }
        }
    }
    for fi68 := 0; fi68 < 19; fi68++ {
        for fi69 := 0; fi69 < 16; fi69++ {
            v70 := []int{(fi69 * 1), fi69}
            var v71 float32 = ((0.1 - 1e10) * (-2.25 / float32(1.5)))
            _ = fi69
            _ = v70
            _ = v71
        }
        var v72 bool = false
        _ = fi68
        _ = v72
    }
    for fi73 := 0; fi73 < 2; fi73++ {
        for fi74 := 0; fi74 < 8; fi74++ {
            if ("a" == "hello") {
                var v75 int16 = 920
                var v76 uint64 = 18446744073709551615
                _ = v75
                _ = v76
            }
            v77 := GenS5{F0: 47}
            fmt.Println(fi74)
            _ = fi74
            _ = v77
        }
        fmt.Println(v2)
        _ = fi73
    }
    v3 = (v29 * v3)
    if ((!false) && ((true || (false && true)) || false)) {
        var v78 int32 = -902
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v37
}
