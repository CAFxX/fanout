package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 string; F1 int32; F2 int32 }
type GenS5 struct { F0 int; F1 int8; F2 int }
type GenS6 struct { F0 uint; F1 int; F2 uint64 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 uint; F1 uint; F2 uint16 }
func genFn0(p0 int64, p1 bool, p2 int32) int32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        switch (fi4 >> ((65535) & 63)) {
        case 0:
            v5 := []int{}
            _ = v5
        case 32767:
            p0 = (p0 + (p0 + (p0 - p0)))
        case -651:
            var v6 int16 = 981
            _ = v6
        default:
            switch "foo bar" {
            case "a":
                v7 := GenS1{F0: int8(p2), F1: ("" + "hello"), F2: 131}
                _ = v7
            case "\n\t":
                p1 = (!p1)
            case "foo bar":
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 float32 = ((float32((1e10 + 1.5)) * 0.0) + 3.14159)
                    var v10 uint32 = 762
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            default:
                v11 := GenS1{F0: 0, F1: ((("foo bar" + ("hello" + ("foo bar" + "foo bar"))) + "hello") + (("\n\t" + "\n\t") + "hello")), F2: 0}
                _ = v11
            }
        }
        _ = fi4
    }
    v12 := GenS1{F0: 111, F1: "", F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}
func genFn1(p0 uint, p1 uint16) int64 {
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            if true {
                switch ("a" + "\n\t") {
                case "hello":
                    p0 += p0
                }
                p1 = uint16(p0)
            }
            var v15 uint64 = 18446744073709551615
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS1{F0: -45, F1: ("" + "foo bar"), F2: 69}
    for fi17 := 0; fi17 < 6; fi17++ {
        p0 -= (^(p0))
        p0 = (^(p0))
        v18 := []int{}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        switch fi19 {
        case -9223372036854775808:
            if (128 > 244) {
                if (2 != 2) {
                    var v20 uint64 = 4294967295
                    _ = v20
                }
                v21 := GenS2{F0: 0, F1: (uint8(p0) >= 255), F2: 0}
                _ = v21
            }
        case 757:
            p1 += p1
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = v16
    return 461
}

func main() {
    var v1 float32 = (-2.25 / (3.14159 * 1.5))
    var v2 int16 = 42
    var v3 int = 519
    var v22 int64 = 0
    var v23 int32 = genFn0(-1, true, -1)
    switch ("\n\t" + "") {
    case "a":
        var v24 uint = 42
        _ = v24
    case "hello":
        v25 := []int{}
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := []int{(v3 | (fi26 / ((fi26) | 1)))}
        if false {
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 int32 = (v23 | (v23 ^ 0))
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                if false {
                    v23 = (v23 ^ (v23 % (((^(v23))) | 1)))
                    v22 -= v22
                    var v31 uint16 = 255
                    _ = v31
                } else {
                    v3 = fi26
                    v27 = append(v27, (v3 + (-(fi30))))
                }
                v3 -= fi30
                var v32 int8 = 127
                switch fi30 {
                case 9223372036854775807:
                    switch 2147483647 {
                    case 128:
                        var v33 uint8 = 127
                        _ = v33
                    }
                case 1:
                    if (true || true) {
                        var v34 int32 = v23
                        if true {
                            v22 = v22
                            v35 := GenS1{F0: (v32 % ((-99) | 1)), F1: ("foo bar" + "a"), F2: 176}
                            var v36 uint = 107
                            if true {
                                switch -2147483648 {
                                case -2147483648:
                                    fmt.Println(v3)
                                case 0:
                                    switch ("a" + "a") {
                                    case "a":
                                        switch 128 {
                                        case -518:
                                            fmt.Println(v36, v32)
                                        case 9223372036854775807:
                                            v37 := GenS1{F0: (v32 / (((v32 + (v32 - (v32 - 1)))) | 1)), F1: ("a" + "hello"), F2: 1}
                                            _ = v37
                                        }
                                    }
                                default:
                                    _ = len(v27)
                                }
                                switch (v3 * (v3 ^ (fi30 * (v3 * fi26)))) {
                                case 326:
                                    v32 += 127
                                    v38 := GenS2{F0: 0, F1: true, F2: (v34 * v34)}
                                    _ = v38
                                case 127:
                                    if (false && false) {
                                        var v39 int64 = (v22 / ((9223372036854775807) | 1))
                                        v36 -= 988
                                        _ = v39
                                    } else {
                                        var v40 uint64 = 4294967295
                                        var v41 int32 = (v34 >> (((v36 ^ 42)) & 63))
                                        var v42 string = "hello"
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                default:
                                    v2 -= -32768
                                }
                                v23 -= (v34 ^ (v23 + (v34 % (((^(v23))) | 1))))
                            } else {
                                v43 := GenS3{F0: v32, F1: (v34 - (-(v34)))}
                                _ = v43
                            }
                            if true {
                                var v44 int64 = genFn1(v36, 412)
                                var v45 uint64 = 0
                                if (((v36 ^ uint(v45)) <= (v36 >> (((v36 << ((v36) & 63))) & 63))) && true) {
                                    if ((-(v34)) != genFn0(v22, false, (^(v23)))) {
                                        var v46 int = 32768
                                        v2 = v2
                                        _ = v46
                                    }
                                }
                                _ = v44
                                _ = v45
                            }
                            switch (((("foo bar" + "hello") + "foo bar") + "\n\t") + (("hello" + ("\n\t" + "a")) + "\n\t")) {
                            case "":
                                var v47 uint8 = 127
                                _ = v47
                            default:
                                var v48 uint8 = 42
                                _ = v48
                            }
                            _ = v35
                            _ = v36
                        } else {
                            switch 32768 {
                            case 1:
                                if (!(!((v32 < int8(v2)) || ((int64(v2) == (v22 << ((42) & 63))) || (1 == 18446744073709551615))))) {
                                    v27 = append(v27, 127)
                                    switch -1 {
                                    case -695:
                                        v32 += (v32 & (-(v32)))
                                    case 128:
                                        v1 = float32(0.1)
                                    }
                                } else {
                                    var v49 uint16 = 32767
                                    _ = v49
                                }
                            }
                            if (!(!true)) {
                                var v50 float32 = -2.25
                                fmt.Println(v22, v34, v3)
                                var v51 uint8 = 127
                                _ = v50
                                _ = v51
                            }
                        }
                        var v52 uint16 = 255
                        v53 := GenS4{F0: "\n\t", F1: (v34 ^ 42), F2: (v34 << ((0) & 63))}
                        fmt.Println(v52, fi26)
                        v54 := GenS5{F0: int(v22), F1: (^(v32)), F2: (^(v3))}
                        _ = v34
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                case -832:
                    v3 = -2147483648
                default:
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                }
                switch "hello" {
                case "foo bar":
                    fmt.Println(v22, fi26, v1)
                    switch "\n\t" {
                    case "foo bar":
                        v3 = int(v23)
                    case "\n\t":
                        v27 = append(v27, fi26)
                    default:
                        _ = len(v27)
                    }
                    switch -845 {
                    case 4294967295:
                        if (127 > 1) {
                            v2 += v2
                            if true {
                                fmt.Println(v22)
                            }
                            if (!(false && (!(!(152 < 1))))) {
                                fmt.Println(v32, fi30, fi26)
                            } else {
                                if true {
                                    v32 += -128
                                }
                                v55 := GenS4{F0: "foo bar", F1: v23, F2: (-(v23))}
                                var v56 bool = false
                                _ = v55
                                _ = v56
                            }
                        }
                    default:
                        v2 = (v2 % ((-32768) | 1))
                    }
                case "":
                    v57 := GenS6{F0: 18446744073709551615, F1: fi26, F2: 2}
                    _ = v57
                }
                var v58 int64 = genFn1(255, 922)
                _ = fi30
                _ = v32
                _ = v58
            }
        } else {
            fmt.Println(fi26, v23, fi26)
            for fi59 := 0; fi59 < 4; fi59++ {
                v60 := GenS7{F0: v2, F1: 65535}
                v27 = append(v27, v3)
                fmt.Println(v1, v3)
                _ = fi59
                _ = v60
            }
        }
        v61 := GenS8{F0: 862, F1: 560, F2: 0}
        _ = fi26
        _ = v27
        _ = v61
    }
    v3 += (v3 * v3)
    v22 = (v22 | 0)
    if (!(((!false) && false) || (false || (1 >= -1)))) {
        var v62 uint16 = 0
        v3 -= 1
        v3 += (-(v3))
        _ = v62
    }
    v63 := []int{v3, v3, (v3 >> ((1) & 63)), v3}
    var v64 uint64 = 897
    v63 = append(v63, (^(v3)))
    var v65 int = v3
    for fi66 := 0; fi66 < 14; fi66++ {
        var v67 uint64 = (v64 | v64)
        if true {
            var v68 bool = (32767 < 65535)
            if ((135 >= 127) && v68) {
                switch "" {
                case "\n\t":
                    for fi69 := 0; fi69 < 7; fi69++ {
                        switch (("a" + "a") + "") {
                        case "a":
                            var v70 int64 = v22
                            _ = v70
                        case "hello":
                            fmt.Println(v67)
                        }
                        v71 := GenS6{F0: 0, F1: fi66, F2: v67}
                        _ = fi69
                        _ = v71
                    }
                }
            }
            v63 = append(v63, (v3 - -9223372036854775808))
            _ = v68
        }
        _ = fi66
        _ = v67
    }
    var v72 uint8 = 128
    v65 -= (v65 + (-(v65)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v63
    _ = v64
    _ = v65
    _ = v72
}
