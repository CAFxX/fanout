package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 int16, p1 string, p2 uint) int64 {
    for fi5 := 0; fi5 < 19; fi5++ {
        p2 = (p2 % (((p2 & (p2 % ((p2) | 1)))) | 1))
        var v6 int16 = (p0 ^ int16(p2))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int = fi7
        p2 = 255
        var v9 int32 = -336
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 0
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 9223372036854775807
}
func genFn1() int32 {
    var v12 uint64 = 716
    if true {
        fmt.Println(v12)
        v12 = 0
        v12 = 18446744073709551615
    }
    var v13 int = 0
    v14 := []int{-2147483648, v13, 42}
    if true {
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 float32 = 0.1
                if (-1 == 1) {
                    if len(v14) > 0 {
                        v14[0] = v13
                    }
                    if (794 == 65535) {
                        _ = len(v14)
                        var v18 int = (-(fi15))
                        if (-2147483648 == 42) {
                            if true {
                                if true {
                                    fmt.Println(fi15, v17)
                                    var v19 int = 143
                                    _ = v19
                                } else {
                                    if false {
                                        v20 := GenS1{F0: -19, F1: -2147483648}
                                        _ = v20
                                    }
                                    v13 += v18
                                    v18 = (v18 >> ((255) & 63))
                                }
                                var v21 uint64 = v12
                                if (!false) {
                                    fmt.Println(v17, v12)
                                }
                                switch (-(fi15)) {
                                case 323:
                                    if (-9223372036854775808 == genFn0(int16(fi16), "hello", 897)) {
                                        v22 := GenS2{F0: (v12 / (((v12 / (((v12 * v21)) | 1))) | 1)), F1: "hello", F2: 4294967295}
                                        fmt.Println(v21, fi15)
                                        var v23 float64 = 1e10
                                        _ = v22
                                        _ = v23
                                    }
                                case 9223372036854775807:
                                    if false {
                                        v18 = fi15
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    } else {
                                        var v24 float64 = 1e10
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        _ = v24
                                    }
                                case -942:
                                    fmt.Println(v13, fi15)
                                }
                                if (593 >= 0) {
                                    v12 = 4294967295
                                    v25 := GenS2{F0: 42, F1: ("a" + "\n\t"), F2: 609}
                                    var v26 bool = (42 <= 0)
                                    switch "foo bar" {
                                    case "":
                                        if false {
                                            v27 := GenS1{F0: 127, F1: -1}
                                            v26 = v26
                                            _ = len(v14)
                                            _ = v27
                                        }
                                    }
                                    v28 := GenS3{F0: 1000}
                                    _ = v25
                                    _ = v26
                                    _ = v28
                                }
                                switch ("a" + "foo bar") {
                                case "foo bar":
                                    v13 += -1
                                default:
                                    v13 -= v18
                                }
                                _ = v21
                            }
                        }
                        _ = v18
                    } else {
                        var v29 int16 = 1
                        switch "" {
                        case "\n\t":
                            switch ("a" + "hello") {
                            case "":
                                var v30 uint64 = v12
                                _ = v30
                            case "foo bar":
                                v29 += (v29 & v29)
                            case "\n\t":
                                v17 = (v17 * v17)
                                fmt.Println(v13, fi16, v13)
                                var v31 int32 = -140
                                _ = v31
                            default:
                                v17 -= (((v17 + 0.1) / (v17 - v17)) / v17)
                            }
                        case "hello":
                            if false {
                                var v32 float32 = (v17 / v17)
                                var v33 uint32 = 2147483647
                                v34 := GenS4{F0: 1}
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        default:
                            var v35 int16 = -32768
                            switch (("hello" + "foo bar") + ((("" + "\n\t") + "a") + "foo bar")) {
                            case "":
                                v12 -= v12
                            case "a":
                                var v36 uint8 = 255
                                _ = v36
                            case "hello":
                                _ = len(v14)
                            }
                            if (int64(v35) != 42) {
                                v29 += v35
                                _ = len(v14)
                                v37 := GenS3{F0: (-(v35))}
                                fmt.Println(v35)
                                _ = v37
                            } else {
                                v38 := GenS1{F0: -47, F1: -2147483648}
                                v39 := GenS2{F0: v12, F1: ("hello" + "\n\t"), F2: 4294967295}
                                if (!(float32(float64((v17 - (v17 + 1.5)))) > ((3.14159 + (1.5 - (float32(0.0) * v17))) + (-2.25 * v17)))) {
                                    if len(v14) > 0 {
                                        v14[0] = (-(v13))
                                    }
                                    var v40 float64 = 1e10
                                    v14 = append(v14, 226)
                                    _ = v40
                                }
                                _ = v38
                                _ = v39
                            }
                            _ = v35
                        }
                        _ = v29
                    }
                }
                v41 := GenS5{F0: 1, F1: 0}
                _ = fi16
                _ = v17
                _ = v41
            }
            switch (v13 % ((1) | 1)) {
            case 9223372036854775807:
                if false {
                    v14 = append(v14, (-(fi15)))
                    if false {
                        v12 = (^(v12))
                        _ = len(v14)
                        if len(v14) > 0 {
                            v14[0] = 32768
                        }
                    } else {
                        if len(v14) > 0 {
                            v14[0] = (fi15 & (fi15 - v13))
                        }
                        v42 := GenS3{F0: -1}
                        _ = v42
                    }
                    v13 -= fi15
                } else {
                    var v43 uint16 = 1
                    _ = v43
                }
            case 851:
                v12 -= v12
            case 780:
                if len(v14) > 0 {
                    v14[0] = (fi15 >> ((42) & 63))
                }
            default:
                v12 = v12
            }
            _ = fi15
        }
        for fi44 := 0; fi44 < 13; fi44++ {
            for fi45 := 0; fi45 < 9; fi45++ {
                v46 := GenS2{F0: v12, F1: "a", F2: 18446744073709551615}
                _ = fi45
                _ = v46
            }
            switch "foo bar" {
            case "\n\t":
                v47 := GenS4{F0: 10}
                _ = v47
            case "foo bar":
                var v48 int16 = 32767
                _ = v48
            case "hello":
                var v49 uint32 = 848
                _ = v49
            }
            v13 += 9223372036854775807
            _ = fi44
        }
    } else {
        if true {
            if (genFn0(32767, ("a" + "foo bar"), 1) <= 42) {
                var v50 int16 = -32768
                _ = len(v14)
                v13 -= (v13 | (v13 ^ v13))
                _ = v50
            }
            v13 += v13
        }
        v14 = append(v14, 127)
        if len(v14) > 0 {
            v14[0] = v13
        }
    }
    _ = v12
    _ = v13
    _ = v14
    return -974
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 42
    var v3 string = "foo bar"
    var v4 int16 = 1
    var v51 uint32 = 1
    v52 := GenS6{F0: 999, F1: 707}
    delete(v1, 2147483647)
    for fi53 := 0; fi53 < 18; fi53++ {
        var v54 string = "hello"
        _ = fi53
        _ = v54
    }
    var v55 string = ((v3 + "foo bar") + (("foo bar" + ("" + v3)) + v3))
    var v56 uint = 4294967295
    var v57 float32 = 0.1
    var v58 uint64 = 219
    v3 = ("foo bar" + (v55 + ""))
    for fi59 := 0; fi59 < 14; fi59++ {
        var v60 int8 = -1
        v2 = (v2 - (v2 * (v2 - v2)))
        v56 -= 241
        _ = fi59
        _ = v60
    }
    v2 = (v2 << (((v56 >> ((685) & 63))) & 63))
    if (!(uint32(v56) < v51)) {
        v61 := GenS7{F0: v4, F1: 9223372036854775807}
        v56 -= (v56 ^ (v56 & v56))
        fmt.Println(v57)
        _ = v61
    } else {
        for fi62 := 0; fi62 < 12; fi62++ {
            var v63 uint16 = 0
            switch "a" {
            case "\n\t":
                for fi64 := 0; fi64 < 20; fi64++ {
                    v65 := GenS8{F0: 1, F1: 0, F2: v56}
                    var v66 int16 = (v4 | v4)
                    _ = fi64
                    _ = v65
                    _ = v66
                }
            case "foo bar":
                v51 += (v51 >> ((v56) & 63))
            case "a":
                _ = len(v1)
            }
            _ = fi62
            _ = v63
        }
    }
    v55 = "a"
    if (9223372036854775807 >= 32767) {
        fmt.Println(v58, v55)
        if true {
            switch v3 {
            case "hello":
                v67 := []int{}
                _ = v67
            case "a":
                switch v55 {
                case "a":
                    v3 = ("" + ("hello" + (v55 + v3)))
                case "\n\t":
                    if true {
                        fmt.Println(v56, v55, v51)
                    }
                default:
                    fmt.Println(v56, v51)
                }
            }
            if true {
                fmt.Println(v55, v58, v57)
            }
            v56 += 511
        }
    } else {
        if true {
            var v68 int64 = (v2 >> (((v56 ^ v56)) & 63))
            _ = v68
        } else {
            var v69 uint16 = 42
            var v70 int32 = 1
            v57 += (v57 + 1.5)
            _ = v69
            _ = v70
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
