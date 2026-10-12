package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 uint32, p1 bool) int16 {
    var v5 uint32 = 2
    var v6 bool = true
    for fi7 := 0; fi7 < 19; fi7++ {
        p1 = (!(v5 != 1))
        v8 := GenS1{F0: (fi7 + fi7), F1: 255, F2: 703}
        _ = fi7
        _ = v8
    }
    if (32768 >= 9223372036854775807) {
        v9 := GenS1{F0: 42, F1: 590, F2: 0}
        v10 := GenS2{F0: 18446744073709551615, F1: "\n\t"}
        _ = v10
        _ = v9
    } else {
        var v11 uint32 = 570
        _ = v11
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            p0 = (p0 << ((65535) & 63))
            fmt.Println(v6, v5)
            p0 += (p0 * p0)
        }
    case "hello":
        var v12 bool = (943 < 9223372036854775807)
        _ = v12
    case "foo bar":
        if true {
            var v13 bool = (32767 == 32767)
            _ = v13
        }
    default:
        p0 = 1
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1000
}

func main() {
    v1 := []uint32{2147483647, 221, 65535, 2}
    v2 := make(map[string]int)
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 20; fi15++ {
        delete(v4, 2147483647)
        fmt.Println(v14)
        _ = fi15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: 42, F1: "a"}
        fmt.Println(v14, fi16)
        if true {
            if true {
                switch (fi16 + 1) {
                case 127:
                    v14 = v14
                case 32767:
                    v18 := GenS1{F0: fi16, F1: v14, F2: 1}
                    _ = v18
                }
            } else {
                v14 += (v14 - 0)
                if (!(false || (427 <= 42))) {
                    v19 := GenS1{F0: (fi16 ^ (fi16 + -9223372036854775808)), F1: v14, F2: 42}
                    _ = v19
                }
                var v20 string = "hello"
                _ = v20
            }
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v14)
    v14 += v14
    var v21 uint8 = 38
    var v22 uint = 255
    fmt.Println(v14)
    var v23 int16 = 1
    if (v21 < (v21 & (v21 % (((v21 >> (((v22 - v22)) & 63))) | 1)))) {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int16 = v23
            if (uint64(v14) >= 4294967295) {
                v26 := GenS1{F0: (fi24 + (^(fi24))), F1: v14, F2: 18446744073709551615}
                for fi27 := 0; fi27 < 18; fi27++ {
                    _ = len(v3)
                    switch "" {
                    case "":
                        v3 = append(v3, 476)
                    case "\n\t":
                        delete(v2, "a")
                    case "hello":
                        switch "" {
                        case "":
                            v22 = (v22 >> ((v22) & 63))
                        case "\n\t":
                            v23 -= (v23 << ((18446744073709551615) & 63))
                        case "hello":
                            fmt.Println(v23, v22)
                            var v28 int16 = (v25 % (((v25 << (((v22 * v22)) & 63))) | 1))
                            v22 = (v22 >> ((18446744073709551615) & 63))
                            _ = v28
                        default:
                            if true {
                                if (499 == -2147483648) {
                                    var v29 int = fi24
                                    v29 -= fi27
                                    _ = v29
                                } else {
                                    v23 = genFn0(52, ((int8(v23) <= -1) || (!true)))
                                    var v30 float32 = 0.0
                                    _ = v30
                                }
                                var v31 float32 = (0.1 * 0.1)
                                var v32 uint16 = v14
                                fmt.Println(fi27)
                                _ = v31
                                _ = v32
                            }
                        }
                        var v33 uint64 = 676
                        switch int(v33) {
                        case 42:
                            if true {
                                var v34 uint16 = v14
                                fmt.Println(fi24, v23, fi27)
                                if ((v25 + (v25 / ((v25) | 1))) != v23) {
                                    var v35 int = int(v33)
                                    v4[1] = (v35 / ((42) | 1))
                                    v4[65535] = (v35 << (((v22 / (((v22 % (((^(v22))) | 1))) | 1))) & 63))
                                    _ = v35
                                } else {
                                    if false {
                                        var v36 bool = false
                                        var v37 string = ""
                                        _ = len(v1)
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 uint8 = (v21 * (v21 ^ v21))
                                        _ = v38
                                    }
                                    if (fi27 == fi24) {
                                        v22 -= v22
                                    } else {
                                        v39 := GenS3{F0: (^(fi24)), F1: (^(v33))}
                                        var v40 int = (fi24 | -2147483648)
                                        var v41 uint16 = (v14 << ((v22) & 63))
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                _ = v34
                            }
                        case -1:
                            v42 := GenS1{F0: (-(fi27)), F1: 32767, F2: (^(v33))}
                            _ = v42
                        }
                        _ = v33
                    default:
                        var v43 bool = false
                        _ = v43
                    }
                    _ = fi27
                }
                v3 = append(v3, 1)
                _ = v26
            } else {
                v4[65535] = (^(fi24))
                v44 := GenS1{F0: (-(fi24)), F1: (v14 * (v14 + 65535)), F2: 1}
                _ = v44
            }
            v14 -= v14
            _ = fi24
            _ = v25
        }
    } else {
        if ((v22 != v22) && ((^(v21)) == v21)) {
            v14 = (v14 | (v14 / (((^(v14))) | 1)))
        } else {
            fmt.Println(v14, v22, v22)
            fmt.Println(v23, v23)
            var v45 uint8 = v21
            _ = v45
        }
        for fi46 := 0; fi46 < 1; fi46++ {
            if true {
                var v47 float32 = (((0.0 / 1.0) / 1.5) / 1.0)
                fmt.Println(v47, v21, fi46)
                _ = v47
            } else {
                v14 = (^(v14))
            }
            fmt.Println(v14, v21)
            _ = fi46
        }
    }
    v22 -= (v22 ^ 255)
    if (!(!("\n\t" >= ("\n\t" + "a")))) {
        fmt.Println(v22, v14, v21)
    }
    var v48 int = int(v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v48
}
