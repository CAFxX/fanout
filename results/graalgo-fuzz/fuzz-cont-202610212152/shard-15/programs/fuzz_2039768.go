package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint64, p2 string) uint8 {
    v7 := []int{42, 32768, -131}
    var v8 uint32 = 2
    var v9 int = 32768
    v10 := GenS1{F0: "foo bar", F1: 1, F2: -128}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 196
}
func genFn1() int64 {
    if false {
        v11 := GenS1{F0: "a", F1: 1, F2: 100}
        var v12 float64 = (0.1 - (1.5 + 1.5))
        _ = v11
        _ = v12
    } else {
        var v13 uint16 = 228
        v14 := GenS2{F0: 2, F1: ((("\n\t" + "a") != "a") && (false && ("a" > "foo bar")))}
        switch 127 {
        case 127:
            for fi15 := 0; fi15 < 1; fi15++ {
                if ((3.14159 == ((0.0 * 0.1) - float64(1e10))) && true) {
                    switch fi15 {
                    case 422:
                        var v16 int64 = -9223372036854775808
                        _ = v16
                    }
                    if false {
                        for fi17 := 0; fi17 < 16; fi17++ {
                            var v18 uint64 = 88
                            switch 510 {
                            case 32767:
                                if (65535 != 339) {
                                    fmt.Println(fi17, fi15)
                                    if ((fi17 + fi17) != fi17) {
                                        var v19 int = (fi15 % ((fi15) | 1))
                                        v13 -= (v13 + v13)
                                        _ = v19
                                    } else {
                                        if true {
                                            switch "a" {
                                            case "":
                                                if true {
                                                    switch 32768 {
                                                    case 4294967295:
                                                        var v20 int16 = 42
                                                        _ = v20
                                                    case 32768:
                                                        fmt.Println(fi15, fi17)
                                                    }
                                                    if false {
                                                        v21 := GenS1{F0: "a", F1: 0, F2: 42}
                                                        _ = v21
                                                    }
                                                    if false {
                                                        var v22 uint32 = 2147483647
                                                        v23 := GenS1{F0: ("\n\t" + "foo bar"), F1: int64(v18), F2: -128}
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                }
                                            case "a":
                                                v18 += 0
                                                v18 = v18
                                            case "\n\t":
                                                fmt.Println(fi15)
                                            default:
                                                fmt.Println(v13, v18, v13)
                                            }
                                        } else {
                                            var v24 uint = 76
                                            _ = v24
                                        }
                                    }
                                    switch fi15 {
                                    case 460:
                                        switch "\n\t" {
                                        case "foo bar":
                                            v25 := GenS3{F0: 15}
                                            _ = v25
                                        }
                                    case 127:
                                        var v26 uint = 2
                                        v18 = 611
                                        _ = v26
                                    default:
                                        if (true || false) {
                                            v13 = v13
                                            v13 -= v13
                                        }
                                    }
                                } else {
                                    var v27 uint64 = 558
                                    fmt.Println(v13)
                                    _ = v27
                                }
                            }
                            _ = fi17
                            _ = v18
                        }
                        var v28 int64 = 9223372036854775807
                        v13 = (v13 & (v13 << ((0) & 63)))
                        _ = v28
                    }
                    v29 := []int{fi15, (fi15 & (fi15 ^ 2147483647)), (fi15 - (fi15 * fi15))}
                    _ = v29
                } else {
                    fmt.Println(fi15)
                    v13 += v13
                    v30 := []int{4294967295, fi15}
                    _ = v30
                }
                v31 := GenS2{F0: 4294967295, F1: true}
                _ = fi15
                _ = v31
            }
        case 580:
            for fi32 := 0; fi32 < 16; fi32++ {
                v13 += v13
                _ = fi32
            }
        }
        _ = v13
        _ = v14
    }
    v33 := GenS3{F0: genFn0(true, 623, "foo bar")}
    var v34 int64 = 9223372036854775807
    if (2 <= 4294967295) {
        v35 := []int{-60}
        if false {
            var v36 int8 = -128
            _ = v36
        } else {
            v37 := GenS2{F0: 2, F1: (128 >= -12)}
            v38 := GenS4{F0: ""}
            switch "foo bar" {
            case "a":
                var v39 float64 = (1.5 * 0.1)
                _ = v39
            default:
                v34 += (v34 * 9223372036854775807)
            }
            _ = v37
            _ = v38
        }
        _ = v35
    } else {
        var v40 int = 32768
        _ = v40
    }
    v41 := []int{127, 32767, 42}
    _ = v33
    _ = v34
    _ = v41
    return -515
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"foo bar", "foo bar", ("" + "\n\t"), "a"}
    v3 := make(map[int]uint32)
    var v4 uint16 = 1
    var v5 int16 = 29
    var v6 uint8 = 255
    for fi42 := 0; fi42 < 8; fi42++ {
        v5 -= v5
        _ = fi42
    }
    if (-9223372036854775808 < -260) {
        v4 -= v4
    }
    fmt.Println(v6, v5, v4)
    v43 := GenS4{F0: "\n\t"}
    v5 = v5
    if (!(((3.14159 + -2.25) * 3.14159) >= -2.25)) {
        var v44 uint16 = (v4 << ((42) & 63))
        v44 = (v44 + (^(v4)))
        _ = v44
    }
    v2 = append(v2, ("\n\t" + ""))
    v45 := GenS4{F0: ("hello" + ("hello" + (("a" + "") + ((("\n\t" + "foo bar") + "a") + "foo bar"))))}
    if true {
        var v46 int16 = 1
        if false {
            for fi47 := 0; fi47 < 7; fi47++ {
                v4 += 2
                _ = fi47
            }
        }
        var v48 uint8 = (v6 + 0)
        _ = v46
        _ = v48
    }
    v4 -= (v4 % (((^(v4))) | 1))
    _ = len(v3)
    _ = len(v1)
    if (9223372036854775807 != 541) {
        fmt.Println(v6, v5)
        var v49 int32 = 2147483647
        delete(v1, 2147483647)
        _ = v49
    }
    if (-1 <= 9223372036854775807) {
        var v50 int16 = 257
        v6 -= (v6 & 127)
        _ = v50
    } else {
        if true {
            v2 = append(v2, "a")
        } else {
            for fi51 := 0; fi51 < 13; fi51++ {
                var v52 uint8 = genFn0(false, 18446744073709551615, "")
                var v53 bool = true
                v5 += v5
                _ = fi51
                _ = v52
                _ = v53
            }
            var v54 uint64 = 273
            _ = v54
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch -963 {
    case 572:
        var v55 int16 = v5
        _ = v55
    default:
        for fi56 := 0; fi56 < 19; fi56++ {
            for fi57 := 0; fi57 < 15; fi57++ {
                v6 = (v6 >> ((1) & 63))
                if (4294967295 < 42) {
                    v58 := GenS5{F0: 65535, F1: false}
                    _ = v58
                } else {
                    v59 := GenS6{F0: 42}
                    _ = v59
                }
                _ = fi57
            }
            var v60 bool = (!true)
            for fi61 := 0; fi61 < 14; fi61++ {
                if ((!(-9223372036854775808 > 9223372036854775807)) || (3.14159 < (0.0 / 1.0))) {
                    if (!(!true)) {
                        var v62 uint64 = 628
                        switch (^(fi56)) {
                        case 2147483647:
                            v63 := GenS6{F0: uint8(v4)}
                            _ = v63
                        case 128:
                            var v64 uint64 = v62
                            _ = v64
                        case 42:
                            fmt.Println(v5, v62, v62)
                        }
                        _ = v62
                    }
                }
                v60 = v60
                var v65 int = (fi61 + fi61)
                _ = fi61
                _ = v65
            }
            _ = fi56
            _ = v60
        }
    }
    fmt.Println(v4, v5)
    fmt.Println(v5)
    fmt.Println(v4, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v6
}
