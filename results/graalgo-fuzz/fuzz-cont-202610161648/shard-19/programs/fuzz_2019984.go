package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int16 }
func genFn0(p0 int8) int64 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint8 = 42
        if (float32((0.0 / (1e10 - 3.14159))) <= 3.14159) {
            var v8 int16 = 0
            v9 := []int{(fi6 | 0)}
            p0 += -128
            _ = v8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v10 := GenS1{F0: 42}
    _ = p0
    _ = v10
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    var v3 uint = 100
    var v4 float32 = 1.5
    var v5 int8 = -41
    v11 := GenS2{F0: 1, F1: -2147483648}
    var v12 int16 = 0
    var v13 bool = (798 <= 1)
    switch "hello" {
    case "a":
        var v14 int16 = v12
        _ = v14
    case "":
        if (!v13) {
            if (false || true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v15 string = "foo bar"
                if v13 {
                    fmt.Println(v13)
                    for fi16 := 0; fi16 < 18; fi16++ {
                        if ((fi16 * fi16) < -9223372036854775808) {
                            v3 += v3
                        } else {
                            if v13 {
                                v17 := GenS3{F0: (-(v12)), F1: ((2147483648 != 996) && (!v13)), F2: 42}
                                _ = v17
                            } else {
                                if v13 {
                                    if (!(127 <= (fi16 >> ((v3) & 63)))) {
                                        var v18 uint64 = 1
                                        _ = v18
                                    }
                                } else {
                                    v5 += (v5 | (v5 % ((1) | 1)))
                                    if (0 != 18446744073709551615) {
                                        v12 += 1000
                                        var v19 uint8 = 255
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                        _ = v19
                                    }
                                    v20 := GenS4{F0: 288, F1: 479, F2: (v12 << (((v3 / ((166) | 1))) & 63))}
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 1; fi21++ {
                                    v3 = v3
                                    fmt.Println(fi21, fi16)
                                    v15 = v15
                                    _ = fi21
                                }
                                fmt.Println(v4, v12)
                            }
                            if (100 > v5) {
                                var v22 int32 = 508
                                _ = v22
                            } else {
                                v12 = v12
                            }
                        }
                        _ = fi16
                    }
                    fmt.Println(v4, v15)
                }
                _ = v15
            }
            var v23 uint32 = 1
            _ = v23
        }
    case "foo bar":
        _ = len(v1)
    }
    v4 += -2.25
    var v24 int8 = 127
    if v13 {
        var v25 int64 = 518
        v12 -= (v12 >> ((0) & 63))
        var v26 int8 = v5
        _ = v25
        _ = v26
    } else {
        v12 = v12
        var v27 bool = true
        for fi28 := 0; fi28 < 14; fi28++ {
            v5 = 101
            v29 := GenS2{F0: 484, F1: -524}
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    v4 += 0.0
    for fi30 := 0; fi30 < 6; fi30++ {
        delete(v1, "\n\t")
        v2 = append(v2, 2)
        _ = fi30
    }
    if (v13 && v13) {
        for fi31 := 0; fi31 < 12; fi31++ {
            v32 := v1[("a" + "")]
            _ = fi31
            _ = v32
        }
    }
    if (v13 && (42 == -59)) {
        switch "\n\t" {
        case "":
            _ = len(v2)
        }
        if len(v2) > 0 {
            v2[0] = 0
        }
    } else {
        v24 += v24
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 float32 = 0.0
            _ = fi33
            _ = v34
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
