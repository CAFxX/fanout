package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 string; F1 string; F2 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 int32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 string, p1 uint, p2 int64) int {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{}
        v6 := GenS1{F0: 19, F1: 52, F2: 32767}
        v7 := GenS1{F0: 1, F1: 127, F2: 65535}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        p0 = ("" + "a")
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int32, p1 uint64, p2 uint) uint16 {
    v9 := GenS2{F0: (p2 - (p2 | (p2 / ((p2) | 1)))), F1: 0}
    var v10 uint8 = 1
    var v11 bool = false
    for fi12 := 0; fi12 < 19; fi12++ {
        p1 = (^(p1))
        if v11 {
            switch (fi12 >> ((p2) & 63)) {
            case 97:
                for fi13 := 0; fi13 < 4; fi13++ {
                    p1 += (p1 + 998)
                    v11 = (268 != 0)
                    _ = fi13
                }
            }
        }
        _ = fi12
    }
    p2 = 464
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    var v1 bool = true
    var v2 uint64 = 4294967295
    var v3 int8 = 1
    switch 127 {
    case 868:
        if v1 {
            switch "a" {
            case "foo bar":
                if true {
                    v1 = v1
                } else {
                    if v1 {
                        var v14 int = genFn0(("a" + "foo bar"), 862, 1)
                        _ = v14
                    }
                    var v15 int = -2147483648
                    v2 = v2
                    _ = v15
                }
            case "a":
                var v16 uint64 = v2
                _ = v16
            default:
                for fi17 := 0; fi17 < 19; fi17++ {
                    v2 += v2
                    _ = fi17
                }
            }
        } else {
            v2 = (v2 - v2)
        }
    }
    var v18 int32 = 2147483647
    if (500 == -9223372036854775808) {
        v19 := []int{128}
        _ = v19
    } else {
        switch 32768 {
        case -219:
            switch "a" {
            case "a":
                fmt.Println(v3)
            case "\n\t":
                fmt.Println(v18)
            }
        }
    }
    switch -9223372036854775808 {
    case 1:
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int64 = 325
            if v1 {
                switch "" {
                case "foo bar":
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if ((fi20 & (fi22 - -997)) < (fi22 + (fi22 << ((4294967295) & 63)))) {
                            var v23 uint64 = v2
                            if v1 {
                                v24 := GenS3{F0: uint16(v2), F1: 2, F2: 42}
                                _ = v24
                            } else {
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    v3 = (v3 | int8(v23))
                                }
                                var v25 int16 = 1
                                _ = v25
                            }
                            _ = v23
                        } else {
                            fmt.Println(fi20, v1, v21)
                            var v26 int8 = v3
                            var v27 int64 = (v21 % (((v21 * 378)) | 1))
                            _ = v26
                            _ = v27
                        }
                        v28 := []int{(fi20 % ((128) | 1)), (fi20 ^ (fi20 + (-(fi20)))), fi20}
                        _ = fi22
                        _ = v28
                    }
                case "":
                    if (2 >= (^(v2))) {
                        v21 += 9223372036854775807
                        if (!v1) {
                            fmt.Println(v1)
                            var v29 float32 = float32(-2.25)
                            v2 = (v2 - v2)
                            _ = v29
                        } else {
                            var v30 uint32 = 4294967295
                            var v31 uint = 0
                            _ = v30
                            _ = v31
                        }
                    } else {
                        if (!((856 < 0) && (1 <= 42))) {
                            for fi32 := 0; fi32 < 12; fi32++ {
                                if true {
                                    var v33 string = "foo bar"
                                    var v34 float32 = float32(0.1)
                                    _ = v33
                                    _ = v34
                                } else {
                                    if v1 {
                                        if (42 >= 146) {
                                            v3 -= v3
                                            if (v2 >= 2) {
                                                if ((fi20 * fi32) <= int(v21)) {
                                                    switch "" {
                                                    case "a":
                                                        var v35 uint32 = 1
                                                        _ = v35
                                                    case "foo bar":
                                                        var v36 uint64 = 1
                                                        _ = v36
                                                    case "":
                                                        v1 = (v1 || v1)
                                                    default:
                                                        var v37 string = "\n\t"
                                                        _ = v37
                                                    }
                                                } else {
                                                    var v38 int16 = -32768
                                                    v18 += (v18 >> ((961) & 63))
                                                    v38 += 403
                                                    _ = v38
                                                }
                                                v1 = false
                                                v39 := GenS2{F0: 2, F1: 2}
                                                _ = v39
                                            }
                                        } else {
                                            switch (fi32 + fi20) {
                                            case 773:
                                                var v40 bool = v1
                                                _ = v40
                                            case 1:
                                                var v41 uint32 = 65535
                                                _ = v41
                                            case 393:
                                                v42 := GenS1{F0: 2, F1: 128, F2: 2}
                                                _ = v42
                                            default:
                                                v3 += (v3 - -9)
                                            }
                                            var v43 int32 = v18
                                            v21 = v21
                                            _ = v43
                                        }
                                    }
                                    v44 := []int{genFn0("hello", 921, v21), fi32, genFn0("hello", 122, v21)}
                                    fmt.Println(fi20, v21)
                                    fmt.Println(fi20)
                                    fmt.Println(v1, v18, v3)
                                    _ = v44
                                }
                                _ = fi32
                            }
                            v45 := []int{fi20, (fi20 << ((18446744073709551615) & 63)), (fi20 * (fi20 / ((4294967295) | 1))), fi20}
                            var v46 int8 = 127
                            _ = v45
                            _ = v46
                        }
                        v1 = (1 <= 2)
                        v47 := GenS4{F0: "hello", F1: "hello", F2: 42}
                        _ = v47
                    }
                default:
                    v48 := GenS5{F0: (v1 && true), F1: -477, F2: v18}
                    _ = v48
                }
                v2 = (v2 + (v2 & v2))
            }
            v21 -= v21
            _ = fi20
            _ = v21
        }
    case -978:
        v49 := GenS6{F0: v3}
        _ = v49
    default:
        v50 := []int{32767, 32768, 856, 494}
        _ = v50
    }
    v51 := []int{-939, genFn0(("a" + ("hello" + ("foo bar" + "a"))), 0, 1), 0, -2147483648}
    if (255 != 253) {
        v52 := GenS2{F0: 42, F1: 439}
        v18 = (v18 & (v18 | v18))
        var v53 uint8 = 255
        _ = v52
        _ = v53
    } else {
        if len(v51) > 0 {
            v51[0] = 2147483647
        }
        if ((v18 <= v18) || (!v1)) {
            v54 := GenS3{F0: 255, F1: 135, F2: 127}
            for fi55 := 0; fi55 < 14; fi55++ {
                var v56 int32 = (-(v18))
                _ = fi55
                _ = v56
            }
            if len(v51) > 0 {
                v51[0] = -650
            }
            _ = v54
        }
        for fi57 := 0; fi57 < 10; fi57++ {
            var v58 string = ("a" + "\n\t")
            if (0.1 != 3.14159) {
                if len(v51) >= 2 {
                    v51 = v51[1:len(v51)]
                }
                var v59 int32 = v18
                _ = v59
            } else {
                v2 -= (v2 - 4294967295)
                for fi60 := 0; fi60 < 4; fi60++ {
                    _ = len(v51)
                    _ = fi60
                }
            }
            _ = fi57
            _ = v58
        }
    }
    v1 = v1
    var v61 float32 = (0.1 + (1.5 * (1.5 + -2.25)))
    var v62 uint64 = (v2 * (v2 + v2))
    for fi63 := 0; fi63 < 15; fi63++ {
        v64 := GenS4{F0: "a", F1: "a", F2: 42}
        v62 = v62
        _ = fi63
        _ = v64
    }
    v2 -= v62
    var v65 float64 = -2.25
    for fi66 := 0; fi66 < 10; fi66++ {
        v51 = append(v51, (fi66 | fi66))
        fmt.Println(v2, v62)
        switch (fi66 & -1) {
        case 4294967295:
            switch "foo bar" {
            case "\n\t":
                v61 += 1e10
            }
        default:
            var v67 int = (fi66 * (-(fi66)))
            _ = v67
        }
        _ = fi66
    }
    if (0.1 == ((((v65 - (0.1 - (v65 + float64((v61 * float32(v65)))))) - v65) + 0.0) - 1e10)) {
        for fi68 := 0; fi68 < 9; fi68++ {
            var v69 uint64 = 2
            switch ("\n\t" + "foo bar") {
            case "a":
                if len(v51) > 0 {
                    v51[0] = (fi68 | fi68)
                }
            case "\n\t":
                for fi70 := 0; fi70 < 9; fi70++ {
                    var v71 bool = v1
                    _ = fi70
                    _ = v71
                }
            }
            _ = fi68
            _ = v69
        }
        for fi72 := 0; fi72 < 9; fi72++ {
            v1 = ((v3 << ((1) & 63)) != 127)
            v18 -= v18
            var v73 uint32 = 2
            _ = fi72
            _ = v73
        }
    }
    v61 = v61
    if false {
        for fi74 := 0; fi74 < 9; fi74++ {
            var v75 string = "foo bar"
            v3 = -128
            _ = len(v51)
            _ = fi74
            _ = v75
        }
        for fi76 := 0; fi76 < 15; fi76++ {
            for fi77 := 0; fi77 < 17; fi77++ {
                fmt.Println(fi76, v3, v3)
                _ = fi77
            }
            _ = fi76
        }
    } else {
        fmt.Println(v65)
    }
    switch "" {
    case "foo bar":
        v78 := GenS1{F0: 191, F1: 186, F2: 649}
        _ = v78
    case "\n\t":
        var v79 string = "\n\t"
        _ = v79
    case "a":
        var v80 uint16 = 819
        _ = v80
    default:
        for fi81 := 0; fi81 < 11; fi81++ {
            if v1 {
                v51 = append(v51, (fi81 ^ (fi81 & fi81)))
                var v82 int = fi81
                _ = v82
            } else {
                if len(v51) > 0 {
                    v51[0] = fi81
                }
            }
            var v83 int32 = (v18 >> ((255) & 63))
            var v84 uint64 = 18446744073709551615
            _ = fi81
            _ = v83
            _ = v84
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v51
    _ = v61
    _ = v62
    _ = v65
}
