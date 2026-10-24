package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 bool; F1 int64; F2 int8 }
type GenS5 struct { F0 int64 }
func genFn0() int32 {
    var v5 uint32 = 833
    v6 := GenS1{F0: "\n\t", F1: 65535, F2: 1}
    if false {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 string = "foo bar"
            if (v8 != "hello") {
                if ((true || (true && true)) || (!(128 >= 246))) {
                    v9 := GenS2{F0: -32768, F1: 42}
                    var v10 uint = 255
                    switch -582 {
                    case 32768:
                        if true {
                            v5 += 0
                        }
                    case 32767:
                        if (!false) {
                            v11 := []int{(-(fi7)), fi7, (fi7 * int(v5))}
                            var v12 uint32 = v5
                            v5 += v5
                            _ = v11
                            _ = v12
                        } else {
                            var v13 int16 = -1
                            for fi14 := 0; fi14 < 4; fi14++ {
                                if (!true) {
                                    v8 = v8
                                    v15 := []int{(fi7 | (fi14 >> ((42) & 63)))}
                                    var v16 int8 = 1
                                    fmt.Println(v13, v16, v10)
                                    _ = v15
                                    _ = v16
                                } else {
                                    fmt.Println(v8)
                                }
                                v5 = (v5 + 730)
                                var v17 int32 = -812
                                _ = fi14
                                _ = v17
                            }
                            _ = v13
                        }
                    case 2147483647:
                        var v18 int16 = -652
                        _ = v18
                    default:
                        v10 -= v10
                    }
                    _ = v10
                    _ = v9
                } else {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        var v20 int32 = 42
                        _ = fi19
                        _ = v20
                    }
                    if false {
                        var v21 uint8 = 80
                        if false {
                            fmt.Println(v5, v21)
                            var v22 int32 = 2147483647
                            _ = v22
                        } else {
                            var v23 uint64 = 42
                            var v24 uint = 239
                            if false {
                                fmt.Println(v23)
                                var v25 uint = (v24 & v24)
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    fmt.Println(fi7, v21, fi26)
                                    var v27 uint8 = 2
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v25
                            }
                            _ = v23
                            _ = v24
                        }
                        var v28 int16 = 135
                        _ = v21
                        _ = v28
                    } else {
                        var v29 bool = (false && (18446744073709551615 != 0))
                        _ = v29
                    }
                }
                var v30 int = fi7
                v31 := GenS2{F0: -277, F1: -1}
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 11; fi32++ {
                fmt.Println(fi7, fi32)
                v33 := []int{32767, (fi7 % ((fi7) | 1)), -5, (fi7 >> ((1) & 63))}
                v8 = ""
                var v34 int32 = 508
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi7
            _ = v8
        }
    }
    var v35 float64 = 1.5
    switch "hello" {
    case "foo bar":
        var v36 string = "foo bar"
        _ = v36
    }
    _ = v35
    _ = v5
    _ = v6
    return 526
}

func main() {
    v1 := []int{-84}
    var v2 uint8 = 246
    var v3 bool = (2 >= uint64(v2))
    var v4 float32 = (0.0 * (((1e10 - 1.5) * (1.5 / (0.1 + (0.0 / 0.1)))) - 3.14159))
    v4 -= v4
    v37 := GenS3{F0: v3, F1: "\n\t"}
    v3 = (42 > 2147483648)
    switch "\n\t" {
    case "hello":
        fmt.Println(v2, v3, v4)
    case "foo bar":
        v3 = true
    case "\n\t":
        v38 := GenS4{F0: (20 <= 0), F1: 85, F2: 1}
        _ = v38
    default:
        var v39 float32 = v4
        _ = v39
    }
    v40 := GenS5{F0: -9223372036854775808}
    v41 := GenS4{F0: v3, F1: -9223372036854775808, F2: -1}
    v4 += (((0.1 / 3.14159) / v4) + (float32(3.14159) / (-2.25 + v4)))
    v4 += v4
    v42 := GenS4{F0: true, F1: -1, F2: int8(v2)}
    var v43 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
