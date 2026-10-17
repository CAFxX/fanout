package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint; F1 int; F2 int32 }
type GenS6 struct { F0 string; F1 string; F2 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8; F1 uint64; F2 string }

func main() {
    v1 := []uint32{65535, 861, 77}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 1
    v4 -= (v4 ^ (v4 / (((v4 >> ((590) & 63))) | 1)))
    var v5 uint32 = 2147483647
    switch 128 {
    case 524:
        var v6 uint16 = 65535
        _ = v6
    case 444:
        if true {
            for fi7 := 0; fi7 < 16; fi7++ {
                switch (("hello" + "\n\t") + "\n\t") {
                case "\n\t":
                    v8 := GenS1{F0: uint64(v5), F1: 0}
                    _ = v8
                case "hello":
                    var v9 uint16 = uint16(v4)
                    _ = v9
                case "":
                    v10 := GenS1{F0: 18446744073709551615, F1: 42}
                    _ = v10
                default:
                    v2[641] = ""
                }
                var v11 uint64 = 97
                _ = fi7
                _ = v11
            }
            v1 = append(v1, 2147483648)
            _ = len(v3)
        } else {
            if false {
                switch ("hello" + "") {
                case "a":
                    switch 32767 {
                    case 2147483647:
                        for fi12 := 0; fi12 < 14; fi12++ {
                            v4 = 127
                            var v13 float32 = float32(0.0)
                            for fi14 := 0; fi14 < 20; fi14++ {
                                var v15 float32 = float32(0.1)
                                _ = fi14
                                _ = v15
                            }
                            _ = fi12
                            _ = v13
                        }
                    case 32768:
                        v16 := GenS2{F0: 100}
                        _ = v16
                    case -9223372036854775808:
                        if (118 > -128) {
                            for fi17 := 0; fi17 < 20; fi17++ {
                                v18 := GenS3{F0: 134, F1: v4}
                                _ = fi17
                                _ = v18
                            }
                        }
                    }
                case "foo bar":
                    if (-42 > -1) {
                        v19 := GenS3{F0: v4, F1: v4}
                        switch 788 {
                        case 9223372036854775807:
                            v3[v5] = (4294967295 != 18446744073709551615)
                        default:
                            v20 := GenS2{F0: 127}
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 uint64 = 887
                            _ = fi21
                            _ = v22
                        }
                        _ = v19
                    } else {
                        fmt.Println(v5, v5)
                    }
                }
            }
        }
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint32 = 4294967295
        v4 = v4
        v4 = v4
        _ = fi23
        _ = v24
    }
    v4 = 255
    switch 37 {
    case 42:
        var v25 uint = 4294967295
        _ = v25
    case -1:
        var v26 int64 = -693
        _ = v26
    }
    v27 := GenS4{F0: true, F1: (v5 < (v5 ^ (v5 >> ((1) & 63))))}
    v5 = (v5 | v5)
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 string = "hello"
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v4 = (v4 / ((v4) | 1))
    if (1 != 6) {
        v31 := GenS5{F0: 1, F1: 32768, F2: -609}
        v32 := GenS2{F0: 42}
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 bool = true
            _ = fi33
            _ = v34
        }
        _ = v31
        _ = v32
    } else {
        switch 539 {
        case 2147483647:
            v2[v5] = "\n\t"
        case 566:
            for fi35 := 0; fi35 < 20; fi35++ {
                for fi36 := 0; fi36 < 16; fi36++ {
                    if true {
                        var v37 float32 = 1.5
                        _ = v37
                    } else {
                        var v38 uint16 = 0
                        switch 128 {
                        case 32767:
                            if false {
                                if true {
                                    switch (fi35 >> ((255) & 63)) {
                                    case -839:
                                        var v39 uint32 = (^(v5))
                                        _ = v39
                                    case 973:
                                        var v40 uint16 = v38
                                        _ = v40
                                    case -636:
                                        switch ("foo bar" + ("" + "hello")) {
                                        case "foo bar":
                                            var v41 float32 = 1e10
                                            _ = v41
                                        case "\n\t":
                                            v42 := GenS6{F0: "hello", F1: "", F2: (v38 * (v38 - (v38 - v38)))}
                                            _ = v42
                                        }
                                    default:
                                        switch fi36 {
                                        case 0:
                                            if false {
                                                var v43 int32 = -2147483648
                                                v5 += (v5 % ((v5) | 1))
                                                _ = v43
                                            } else {
                                                v4 = v4
                                                v44 := GenS7{F0: 746}
                                                switch (-(fi36)) {
                                                case -9223372036854775808:
                                                    delete(v3, v5)
                                                case 4294967295:
                                                    v38 = v38
                                                default:
                                                    v38 = (v38 << ((942) & 63))
                                                }
                                                v5 = (v5 - v5)
                                                var v45 int16 = -1
                                                var v46 int32 = 42
                                                v4 -= (v4 << ((0) & 63))
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            }
                                        case -2147483648:
                                            var v47 uint64 = 4294967295
                                            _ = v47
                                        }
                                    }
                                    if (0 > 255) {
                                        v48 := GenS8{F0: 42, F1: 4294967295, F2: "a"}
                                        if true {
                                            _ = len(v2)
                                            switch -1 {
                                            case -1:
                                                v5 -= (v5 | (v5 + v5))
                                            case 32768:
                                                v5 = (v5 % ((0) | 1))
                                            default:
                                                v4 = (v4 ^ v4)
                                            }
                                        } else {
                                            if true {
                                                v4 = v4
                                            }
                                            v49 := GenS5{F0: 42, F1: (^(fi36)), F2: 0}
                                            v5 = (^(v5))
                                            var v50 int8 = 100
                                            v51 := v3[(^(v5))]
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                        if ((true || false) || false) {
                                            v52 := GenS8{F0: (v4 ^ v4), F1: 2, F2: ""}
                                            _ = v52
                                        }
                                        _ = v48
                                    } else {
                                        if true {
                                            if false {
                                                if (!true) {
                                                    var v53 uint16 = 0
                                                    _ = v53
                                                }
                                                var v54 uint64 = 788
                                                var v55 int64 = 9223372036854775807
                                                _ = v54
                                                _ = v55
                                            } else {
                                                v56 := GenS4{F0: true, F1: true}
                                                var v57 uint64 = 360
                                                var v58 uint64 = 42
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                            var v59 float32 = 1.5
                                            _ = v59
                                        }
                                        if false {
                                            v60 := v3[19]
                                            v61 := GenS8{F0: v4, F1: 42, F2: "\n\t"}
                                            v38 = 42
                                            _ = v60
                                            _ = v61
                                        }
                                        _ = len(v2)
                                        v38 = (v38 + v38)
                                        if len(v1) > 0 {
                                            v1[0] = (v5 | (v5 / (((v5 ^ (v5 >> ((4294967295) & 63)))) | 1)))
                                        }
                                    }
                                    if (!false) {
                                        _ = len(v1)
                                    }
                                } else {
                                    var v62 string = ("a" + "hello")
                                    var v63 string = v62
                                    _ = v62
                                    _ = v63
                                }
                            } else {
                                fmt.Println(fi36, fi36)
                            }
                        }
                        if (!("a" <= (("a" + ("foo bar" + ("\n\t" + ("a" + "hello")))) + "\n\t"))) {
                            v5 += (^(v5))
                        }
                        var v64 uint16 = v38
                        _ = v38
                        _ = v64
                    }
                    var v65 uint8 = 154
                    _ = fi36
                    _ = v65
                }
                if len(v1) > 0 {
                    v1[0] = (^(v5))
                }
                _ = fi35
            }
        }
    }
    switch "hello" {
    case "\n\t":
        v4 += (v4 * (v4 - (v4 << ((255) & 63))))
    default:
        for fi66 := 0; fi66 < 3; fi66++ {
            v5 -= v5
            var v67 uint32 = (^(v5))
            _ = fi66
            _ = v67
        }
    }
    v68 := GenS2{F0: 27}
    v5 -= v5
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v68
}
