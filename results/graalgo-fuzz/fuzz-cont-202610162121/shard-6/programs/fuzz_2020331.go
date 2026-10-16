package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int64 }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int64; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint64; F1 string }

func main() {
    var v1 int8 = 22
    v2 := make(map[int]bool)
    var v3 float64 = 3.14159
    v4 := make(map[int]int)
    v5 := make(map[uint32]string)
    switch "foo bar" {
    case "":
        if true {
            var v6 int16 = int16(v1)
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 float32 = 1.5
                if false {
                    v9 := []int{(fi7 * -438), fi7, fi7, fi7}
                    var v10 uint8 = 127
                    _ = v10
                    _ = v9
                } else {
                    switch ("foo bar" + ("a" + "\n\t")) {
                    case "\n\t":
                        v11 := GenS1{F0: (0 > 0), F1: 0, F2: 9223372036854775807}
                        _ = v11
                    default:
                        fmt.Println(fi7, v6)
                    }
                    switch 32767 {
                    case 32768:
                        if false {
                            if false {
                                if (v8 > ((v8 / v8) - ((v8 * (v8 + ((v8 - v8) / (v8 - v8)))) - (((0.1 - v8) + (0.0 - (-2.25 + v8))) * v8)))) {
                                    if false {
                                        var v12 int32 = -567
                                        if true {
                                            v13 := GenS2{F0: 4294967295, F1: 1, F2: 1}
                                            if ((fi7 | fi7) >= fi7) {
                                                var v14 int64 = -1
                                                v1 += (v1 << ((65535) & 63))
                                                _ = v14
                                            }
                                            _ = v13
                                        }
                                        _ = v12
                                    }
                                    v8 -= 1e10
                                    v1 -= (v1 + -128)
                                } else {
                                    var v15 int8 = (v1 ^ v1)
                                    _ = v15
                                }
                                var v16 float32 = float32(3.14159)
                                v17 := GenS1{F0: (0 < -955), F1: 0, F2: 0}
                                _ = v16
                                _ = v17
                            } else {
                                v18 := []int{(fi7 & (fi7 | -582)), fi7, fi7}
                                v6 = -1
                                _ = v18
                            }
                            var v19 uint8 = 2
                            _ = v19
                        }
                    case 2147483647:
                        var v20 uint64 = 2
                        _ = v20
                    case 1:
                        if false {
                            v21 := GenS3{F0: -318}
                            fmt.Println(v8, v8)
                            _ = v21
                        } else {
                            if true {
                                v2[32768] = (false || (0 < 0))
                            } else {
                                v22 := GenS4{F0: (!(0 <= 1)), F1: 42, F2: (-(fi7))}
                                _ = v22
                            }
                            fmt.Println(v8, v3, v8)
                            v8 -= 0.0
                        }
                    default:
                        if (false && false) {
                            fmt.Println(fi7, v3, v3)
                            switch "foo bar" {
                            case "foo bar":
                                for fi23 := 0; fi23 < 6; fi23++ {
                                    var v24 float32 = (v8 * (0.1 - v8))
                                    switch (fi7 % (((^(fi23))) | 1)) {
                                    case -551:
                                        var v25 uint = 255
                                        _ = v25
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                            default:
                                for fi26 := 0; fi26 < 2; fi26++ {
                                    v8 = v8
                                    _ = fi26
                                }
                            }
                        }
                    }
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            var v27 int8 = v1
            switch 633 {
            case 2147483647:
                var v28 bool = false
                _ = v28
            default:
                v27 = v27
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 int32 = 0
                    fmt.Println(v3, v3)
                    _ = fi30
                    _ = v31
                }
                if (int64(fi29) >= 882) {
                    if true {
                        v3 = 3.14159
                    } else {
                        v32 := GenS5{F0: 88}
                        var v33 int8 = v1
                        v27 -= v1
                        _ = v32
                        _ = v33
                    }
                }
                v27 = v27
                _ = fi29
            }
            _ = v27
        }
    case "foo bar":
        if (!true) {
            v34 := GenS6{F0: 1000, F1: 0}
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        if (-9223372036854775808 <= -144) {
            var v36 int = (fi35 | fi35)
            if false {
                v37 := []int{4294967295, (v36 | v36), (v36 | (v36 << ((42) & 63))), -48}
                _ = v37
            }
            _ = v36
        } else {
            v38 := []int{4294967295}
            if (((1.5 * 0.0) + 3.14159) >= float64((0.0 / -2.25))) {
                v3 = v3
                var v39 uint32 = 423
                v39 = v39
                _ = v39
            }
            for fi40 := 0; fi40 < 18; fi40++ {
                v5[126] = "foo bar"
                _ = fi40
            }
            _ = v38
        }
        _ = fi35
    }
    var v41 uint8 = 128
    fmt.Println(v3, v1)
    v41 += (v41 | (v41 >> ((2) & 63)))
    v42 := GenS7{F0: 381, F1: "a"}
    fmt.Println(v1, v1, v3)
    var v43 int = 209
    if false {
        v44 := []int{}
        fmt.Println(v1, v43, v1)
        _ = v44
    }
    v5[65535] = "foo bar"
    v4[(v43 >> ((274) & 63))] = v43
    var v45 uint32 = 683
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v5
}
