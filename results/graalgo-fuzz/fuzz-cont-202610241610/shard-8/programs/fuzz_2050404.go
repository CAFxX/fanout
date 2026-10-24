package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16 }
func genFn0() int32 {
    v4 := []int{2147483647, 9223372036854775807}
    var v5 uint64 = 1
    v5 = (v5 * 18446744073709551615)
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 bool = (("\n\t" == ("" + ("hello" + "\n\t"))) || (127 < 127))
    v2 := make(map[uint32]int)
    var v3 uint32 = 0
    var v6 int8 = 100
    if ((0.0 + 1e10) < (((-2.25 + -2.25) * (1e10 * 0.0)) * 1e10)) {
        fmt.Println(v3, v1, v6)
        v7 := GenS1{F0: -429, F1: (v3 & 42)}
        _ = v7
    } else {
        var v8 string = (("hello" + "a") + "hello")
        if v1 {
            var v9 uint = 2
            v10 := GenS2{F0: -403, F1: 268, F2: (v6 << (((v9 + (v9 % ((v9) | 1)))) & 63))}
            _ = v10
            _ = v9
        } else {
            v11 := GenS3{F0: 900, F1: (^(v3)), F2: uint(v3)}
            v12 := []int{127, 1, 642, 127}
            _ = v11
            _ = v12
        }
        v13 := []int{32767}
        _ = v13
        _ = v8
    }
    var v14 int = 128
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int = fi15
        switch "a" {
        case "a":
            if false {
                v3 = v3
                var v17 uint64 = 657
                _ = v17
            }
        case "":
            switch "\n\t" {
            case "foo bar":
                v6 += 98
            case "":
                v16 += (v16 << ((42) & 63))
            }
        case "foo bar":
            if v1 {
                if ("\n\t" <= "hello") {
                    v6 = -1
                }
                var v18 string = "\n\t"
                _ = v18
            }
        default:
            if (v16 >= -9223372036854775808) {
                if v1 {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        var v20 int64 = -9223372036854775808
                        switch "foo bar" {
                        case "":
                            var v21 uint64 = 18446744073709551615
                            _ = v21
                        case "hello":
                            v22 := v2[v3]
                            _ = v22
                        default:
                            v23 := GenS2{F0: (fi19 * -1), F1: (v20 / (((v20 % ((int64(v3)) | 1))) | 1)), F2: v6}
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                    }
                    v24 := GenS3{F0: 9223372036854775807, F1: 893, F2: 42}
                    _ = v24
                }
                switch "\n\t" {
                case "hello":
                    v25 := GenS4{F0: 2, F1: (v6 - (v6 ^ (v6 - 121))), F2: 560}
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v27 := []int{}
                        if v1 {
                            if v1 {
                                v28 := GenS3{F0: 42, F1: v3, F2: 255}
                                switch "foo bar" {
                                case "foo bar":
                                    var v29 uint16 = 1
                                    _ = v29
                                case "":
                                    if false {
                                        v30 := GenS4{F0: 2, F1: v6, F2: 765}
                                        v31 := GenS4{F0: 127, F1: (v6 * v6), F2: 910}
                                        v27 = append(v27, fi15)
                                        _ = v30
                                        _ = v31
                                    }
                                case "hello":
                                    if (fi15 == v14) {
                                        var v32 uint64 = 1
                                        if v1 {
                                            var v33 uint = 65535
                                            _ = v33
                                        }
                                        var v34 int8 = v6
                                        v3 += v3
                                        _ = v32
                                        _ = v34
                                    } else {
                                        var v35 uint8 = 42
                                        switch "" {
                                        case "":
                                            v36 := GenS1{F0: 0, F1: 42}
                                            _ = v36
                                        default:
                                            switch fi15 {
                                            case 369:
                                                v6 = v6
                                            }
                                        }
                                        fmt.Println(v16, fi15, v16)
                                        _ = v35
                                    }
                                }
                                v37 := GenS5{F0: -839}
                                _ = v28
                                _ = v37
                            } else {
                                fmt.Println(fi26)
                            }
                        }
                        delete(v2, v3)
                        v1 = false
                        v14 = fi26
                        _ = fi26
                        _ = v27
                    }
                }
            } else {
                if (-32768 < 1) {
                    var v38 uint32 = v3
                    _ = v38
                }
                var v39 float64 = -2.25
                _ = v39
            }
        }
        var v40 int32 = genFn0()
        _ = fi15
        _ = v16
        _ = v40
    }
    v1 = v1
    switch (("" + "\n\t") + "\n\t") {
    case "hello":
        for fi41 := 0; fi41 < 11; fi41++ {
            if true {
                v42 := GenS6{F0: uint16(v3)}
                switch (("" + "hello") + "\n\t") {
                case "":
                    var v43 int16 = 1000
                    _ = v43
                default:
                    fmt.Println(v1, fi41)
                }
                _ = v42
            }
            _ = fi41
        }
    case "foo bar":
        if true {
            v6 -= 0
            switch (("a" + ("hello" + ("hello" + (("" + ("a" + "")) + ("\n\t" + "hello"))))) + "") {
            case "":
                switch "" {
                case "hello":
                    v44 := GenS4{F0: 2, F1: v6, F2: 353}
                    _ = v44
                case "":
                    fmt.Println(v1)
                case "a":
                    var v45 float64 = (1.5 - (float64(0.0) - 3.14159))
                    _ = v45
                }
            default:
                fmt.Println(v1, v1, v3)
            }
            var v46 uint16 = 1
            _ = v46
        }
    case "":
        var v47 uint = 18446744073709551615
        _ = v47
    }
    var v48 uint32 = (^(v3))
    v49 := []int{(v14 & (v14 & -9223372036854775808)), -9223372036854775808, v14, v14}
    v1 = (v1 || (!((v3 - (^(v48))) <= (v48 * v3))))
    if false {
        v50 := GenS2{F0: v14, F1: 1, F2: v6}
        for fi51 := 0; fi51 < 2; fi51++ {
            var v52 int32 = -418
            _ = fi51
            _ = v52
        }
        fmt.Println(v3, v1, v1)
        _ = v50
    }
    if (genFn0() < -477) {
        var v53 string = (("foo bar" + "foo bar") + "foo bar")
        var v54 uint64 = 399
        _ = v53
        _ = v54
    } else {
        var v55 uint16 = 0
        switch (v14 + v14) {
        case 42:
            v56 := GenS1{F0: -87, F1: (v3 | v48)}
            _ = v56
        default:
            for fi57 := 0; fi57 < 15; fi57++ {
                if false {
                    for fi58 := 0; fi58 < 19; fi58++ {
                        v3 = v48
                        _ = fi58
                    }
                    v1 = true
                } else {
                    var v59 uint16 = v55
                    var v60 int32 = -2147483648
                    v2[v48] = fi57
                    _ = v59
                    _ = v60
                }
                _ = fi57
            }
        }
        _ = v55
    }
    if (721 != 42) {
        if v1 {
            var v61 int16 = 247
            v62 := v2[v3]
            switch v14 {
            case 4294967295:
                switch "\n\t" {
                case "foo bar":
                    var v63 uint8 = 252
                    _ = v63
                }
            default:
                for fi64 := 0; fi64 < 5; fi64++ {
                    for fi65 := 0; fi65 < 16; fi65++ {
                        var v66 int32 = -209
                        var v67 uint8 = 42
                        v3 -= v3
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    var v68 uint8 = 140
                    _ = fi64
                    _ = v68
                }
            }
            _ = v61
            _ = v62
        }
        var v69 uint8 = 128
        _ = v69
    } else {
        v14 = v14
        if (v1 || v1) {
            for fi70 := 0; fi70 < 11; fi70++ {
                v1 = (!(!false))
                v3 = 927
                v6 += (v6 >> ((255) & 63))
                _ = fi70
            }
            if (true && ((^(v3)) != (v48 + v48))) {
                v6 -= v6
            } else {
                fmt.Println(v14)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v6
}
