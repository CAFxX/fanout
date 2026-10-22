package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 int64 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int, p1 uint8) int64 {
    fmt.Println(p1, p1)
    v7 := []int{(p0 >> ((370) & 63))}
    v7 = append(v7, -804)
    var v8 uint8 = p1
    if (!false) {
        var v9 int8 = -1
        _ = v9
    } else {
        v10 := GenS1{F0: 2}
        var v11 int32 = 42
        var v12 string = "\n\t"
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint64) uint16 {
    if false {
        var v13 uint64 = (p0 ^ p0)
        _ = v13
    }
    v14 := GenS1{F0: 108}
    _ = p0
    _ = v14
    return 2
}

func main() {
    v1 := []string{"\n\t", "hello", (("\n\t" + ("" + ("\n\t" + "hello"))) + "hello"), "hello"}
    v2 := make(map[string]uint32)
    v3 := []string{("a" + "foo bar"), "hello", "foo bar", "\n\t"}
    var v4 int8 = -1
    var v5 uint32 = 601
    v6 := make(map[int]bool)
    var v15 int16 = 32767
    fmt.Println(v4, v4, v5)
    var v16 bool = (1 != 2)
    var v17 uint64 = 532
    fmt.Println(v16, v4)
    fmt.Println(v4, v4, v15)
    var v18 float64 = 0.0
    v15 -= v15
    v19 := GenS2{F0: int(v5), F1: 255}
    v20 := GenS2{F0: -2147483648, F1: 998}
    if v16 {
        v18 = v18
        fmt.Println(v4, v15)
        v18 = (v18 + v18)
    }
    var v21 uint8 = 255
    switch "foo bar" {
    case "a":
        if v16 {
            v15 -= (v15 % ((v15) | 1))
            var v22 uint8 = v21
            switch "" {
            case "hello":
                v23 := GenS1{F0: 1}
                _ = v23
            case "":
                switch 1 {
                case 4294967295:
                    for fi24 := 0; fi24 < 1; fi24++ {
                        if (v17 < uint64(v5)) {
                            var v25 string = ""
                            var v26 uint32 = v5
                            if true {
                                v5 = (v26 >> ((1) & 63))
                                switch "foo bar" {
                                case "":
                                    v22 = (v22 * (v22 << ((710) & 63)))
                                case "\n\t":
                                    var v27 float64 = 1e10
                                    _ = v27
                                default:
                                    v6[(fi24 ^ fi24)] = (32767 < 307)
                                }
                            }
                            _ = v25
                            _ = v26
                        } else {
                            if (((0.1 * ((3.14159 / v18) / v18)) * float64(((0.1 - 3.14159) - ((1.5 + -2.25) / 0.1)))) >= float64(((-2.25 / 0.1) / -2.25))) {
                                v28 := v6[4294967295]
                                _ = v28
                            } else {
                                fmt.Println(v5, v15)
                            }
                            v6[127] = (4294967295 != v17)
                        }
                        for fi29 := 0; fi29 < 4; fi29++ {
                            var v30 string = ("" + "")
                            _ = fi29
                            _ = v30
                        }
                        _ = fi24
                    }
                case -718:
                    v4 += (v4 ^ v4)
                default:
                    switch (("\n\t" + "hello") + "") {
                    case "":
                        if (!((2 > 18446744073709551615) || v16)) {
                            var v31 int8 = v4
                            _ = v31
                        }
                    case "foo bar":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if (v16 || v16) {
                            v15 -= v15
                            v32 := GenS3{F0: uint(v22), F1: true, F2: (^(v21))}
                            _ = v32
                        } else {
                            for fi33 := 0; fi33 < 2; fi33++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi33
                            }
                            var v34 uint = 536
                            _ = v34
                        }
                    }
                }
            case "\n\t":
                for fi35 := 0; fi35 < 5; fi35++ {
                    if v16 {
                        if v16 {
                            fmt.Println(v21)
                        }
                        v36 := GenS4{F0: 4294967295, F1: (v15 & 42)}
                        fmt.Println(v5)
                        _ = v36
                    }
                    _ = fi35
                }
            default:
                var v37 float64 = v18
                _ = v37
            }
            _ = v22
        }
    }
    delete(v6, -813)
    v38 := GenS2{F0: 32767, F1: 0}
    v2["hello"] = v5
    switch "\n\t" {
    case "foo bar":
        if (genFn1(v17) <= 0) {
            switch 0 {
            case -857:
                fmt.Println(v4, v17, v17)
            case 997:
                v21 = 127
            case 0:
                v18 -= (v18 * (v18 / float64(((1.5 * (0.1 * ((-2.25 / 0.1) * 3.14159))) * (0.1 - float32((float64(1.5) / -2.25)))))))
            }
            v39 := GenS5{F0: -2147483648, F1: (v17 * 824), F2: 42}
            if (!true) {
                v15 += v15
                if ((false || (!((v15 < v15) && (94 >= 42)))) || ((v21 % ((1) | 1)) != v21)) {
                    switch -9223372036854775808 {
                    case -517:
                        var v40 float32 = (float32(v18) * -2.25)
                        _ = v40
                    case -657:
                        v41 := GenS4{F0: 0, F1: -657}
                        _ = v41
                    case 42:
                        if ((v15 > (^(v15))) && v16) {
                            var v42 uint16 = genFn1(v17)
                            _ = v42
                        } else {
                            var v43 float64 = 0.1
                            v44 := GenS4{F0: uint(v17), F1: v15}
                            _ = len(v1)
                            _ = v43
                            _ = v44
                        }
                    default:
                        if (!true) {
                            var v45 int64 = 0
                            v46 := v2["hello"]
                            switch 128 {
                            case 2147483647:
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            default:
                                v47 := GenS2{F0: 32768, F1: 42}
                                _ = v47
                            }
                            _ = v45
                            _ = v46
                        }
                    }
                    var v48 int16 = v15
                    _ = v48
                }
            } else {
                v49 := GenS6{F0: int16(v5)}
                if true {
                    switch 127 {
                    case -426:
                        var v50 int64 = 399
                        _ = v50
                    case 128:
                        var v51 uint8 = (^(v21))
                        _ = v51
                    default:
                        var v52 uint64 = (^(v17))
                        _ = v52
                    }
                    v21 += (v21 + v21)
                } else {
                    var v53 int32 = 42
                    _ = v53
                }
                _ = v49
            }
            _ = v39
        } else {
            v5 = v5
            v54 := v2["a"]
            var v55 uint32 = v54
            _ = v54
            _ = v55
        }
    case "\n\t":
        fmt.Println(v15)
    default:
        var v56 float64 = (v18 / (v18 + (v18 - float64((-2.25 / (1.5 + (float32(1.5) * 3.14159)))))))
        _ = v56
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
