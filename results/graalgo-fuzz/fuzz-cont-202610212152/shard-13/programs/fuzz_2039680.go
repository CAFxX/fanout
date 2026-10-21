package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int8; F1 int16; F2 uint }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 string; F1 uint16; F2 uint }
func genFn0(p0 uint16) uint {
    v6 := GenS1{F0: 1}
    var v7 string = "hello"
    _ = p0
    _ = v6
    _ = v7
    return 41
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]string)
    var v3 int8 = -128
    v4 := []string{}
    var v5 uint8 = 42
    var v8 uint16 = 32767
    switch 32767 {
    case 324:
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS2{F0: 749}
            _ = fi9
            _ = v10
        }
    case -2147483648:
        switch 1 {
        case 2147483647:
            var v11 float64 = float64(0.1)
            _ = v11
        case 724:
            var v12 string = ""
            _ = v12
        case 886:
            v13 := GenS3{F0: 42, F1: 463}
            _ = v13
        default:
            for fi14 := 0; fi14 < 17; fi14++ {
                if (false && ((v8 * (v8 + v8)) == (v8 & (v8 & 605)))) {
                    v5 = (v5 / ((v5) | 1))
                    var v15 int = fi14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 uint16 = v8
                        switch fi16 {
                        case 2147483647:
                            switch (fi14 + (fi14 >> ((1) & 63))) {
                            case 446:
                                var v18 uint16 = v17
                                _ = v18
                            default:
                                v19 := GenS4{F0: (^(fi16)), F1: 361, F2: (v5 / (((v5 << ((255) & 63))) | 1))}
                                _ = v19
                            }
                            v17 = (v17 & 2)
                        case 32768:
                            var v20 float64 = 0.0
                            _ = v20
                        case 4294967295:
                            v1 -= (v1 ^ v1)
                        default:
                            var v21 int64 = -1
                            var v22 int16 = 1
                            v22 -= (v22 ^ (-(v22)))
                            _ = v21
                            _ = v22
                        }
                        _ = fi16
                        _ = v17
                    }
                    v5 = 127
                    var v23 uint16 = v8
                    _ = v23
                }
                _ = fi14
            }
        }
    case 128:
        var v24 float32 = 1.5
        _ = v24
    }
    fmt.Println(v8, v8, v3)
    var v25 int8 = (-(v3))
    if true {
        v26 := GenS2{F0: (v1 << ((255) & 63))}
        if true {
            if false {
                v5 = v5
            }
            v27 := GenS2{F0: v1}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 2; fi28++ {
                v3 = (v25 & (v25 | (v25 | (v3 ^ 1))))
                _ = fi28
            }
            v4 = append(v4, (("a" + ("foo bar" + "a")) + "a"))
            var v29 int32 = (v1 % ((-1) | 1))
            _ = v29
        }
        v30 := GenS2{F0: (v1 >> ((genFn0((v8 << ((0) & 63)))) & 63))}
        _ = v26
        _ = v30
    }
    v1 += v1
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v8)
    switch 32768 {
    case 380:
        for fi31 := 0; fi31 < 7; fi31++ {
            for fi32 := 0; fi32 < 14; fi32++ {
                var v33 string = "a"
                if (false && false) {
                    var v34 uint16 = (v8 % ((80) | 1))
                    if false {
                        switch (fi31 << ((42) & 63)) {
                        case 127:
                            var v35 uint = 4294967295
                            if false {
                                v1 -= -553
                                v36 := GenS5{F0: (v3 / (((v3 + (v25 - v3))) | 1)), F1: -1, F2: (^(v35))}
                                _ = v36
                            }
                            _ = v35
                        case 128:
                            if true {
                                _ = len(v4)
                                fmt.Println(v33, v8, fi32)
                            } else {
                                _ = len(v4)
                                switch "\n\t" {
                                case "hello":
                                    if ((v8 << ((1) & 63)) != v34) {
                                        var v37 uint64 = 685
                                        v5 -= (v5 % ((v5) | 1))
                                        _ = v37
                                    } else {
                                        switch v33 {
                                        case "hello":
                                            var v38 int64 = 42
                                            _ = v38
                                        case "\n\t":
                                            var v39 uint = 99
                                            var v40 uint32 = 0
                                            v40 = (v40 & v40)
                                            delete(v2, v40)
                                            v2[v40] = v33
                                            _ = v39
                                            _ = v40
                                        }
                                        v41 := GenS6{F0: (v33 + "\n\t"), F1: 105}
                                        switch -1 {
                                        case 128:
                                            switch v33 {
                                            case "foo bar":
                                                v42 := GenS7{F0: v33, F1: 1, F2: 18446744073709551615}
                                                _ = v42
                                            }
                                        case -940:
                                            v3 -= (v25 % (((v3 / ((v25) | 1))) | 1))
                                        case -77:
                                            fmt.Println(v25, v8)
                                        }
                                        _ = v41
                                    }
                                case "\n\t":
                                    v43 := GenS1{F0: 42}
                                    _ = v43
                                case "foo bar":
                                    v44 := GenS3{F0: 1, F1: 9223372036854775807}
                                    _ = v44
                                }
                                var v45 int32 = (v1 ^ v1)
                                v8 -= v34
                                _ = v45
                            }
                        default:
                            v46 := GenS5{F0: (-(v25)), F1: 32767, F2: genFn0((v34 % (((v34 / ((799) | 1))) | 1)))}
                            _ = v46
                        }
                        switch fi31 {
                        case 9223372036854775807:
                            var v47 int8 = (v25 % ((v3) | 1))
                            _ = v47
                        case -834:
                            v33 = v33
                        case 32768:
                            fmt.Println(v34, fi31, v25)
                        default:
                            var v48 uint64 = 969
                            _ = v48
                        }
                    }
                    _ = v34
                } else {
                    fmt.Println(v3, fi31, fi32)
                    fmt.Println(v3)
                    var v49 int = 1
                    fmt.Println(v25, v8)
                    fmt.Println(v5, v49)
                    _ = v49
                }
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
    case 1:
        var v50 uint32 = 1
        _ = v50
    case 127:
        delete(v2, 4294967295)
    default:
        var v51 uint64 = 4294967295
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
