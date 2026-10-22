package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint8) uint {
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 13; fi8++ {
        p0 = (p0 | 255)
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int32 = 1
            if ((((-2.25 / -2.25) * 0.1) - 0.1) <= 3.14159) {
                v11 := []int{fi9, 42, (fi9 << ((1) & 63))}
                var v12 int = fi9
                _ = v11
                _ = v12
            } else {
                v10 = v10
                v13 := GenS1{F0: 51}
                var v14 uint64 = v7
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v7 -= v7
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint8 = p0
            v18 := GenS2{F0: (v17 << ((42) & 63)), F1: 42, F2: 42}
            fmt.Println(fi15)
            if ((!false) && false) {
                fmt.Println(p0)
                v7 = (^(v7))
                v7 -= v7
            } else {
                v19 := []int{fi15, (fi16 % ((128) | 1))}
                _ = v19
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        switch -2147483648 {
        case 48:
            v20 := GenS2{F0: (p0 + (p0 / (((p0 % ((p0) | 1))) | 1))), F1: 42, F2: (v7 - (v7 / ((v7) | 1)))}
            _ = v20
        case 32768:
            var v21 uint32 = 129
            _ = v21
        }
        _ = fi15
    }
    _ = p0
    _ = v7
    return 72
}
func genFn1() uint {
    var v22 int8 = 0
    v23 := GenS3{F0: 0, F1: 194}
    fmt.Println(v22)
    _ = v22
    _ = v23
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -1
    var v3 uint = 7
    v4 := make(map[int]int)
    var v5 uint32 = 0
    v6 := []uint32{}
    switch "" {
    case "foo bar":
        _ = len(v6)
    }
    v2 += int8(v5)
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 uint32 = (v5 % ((975) | 1))
        fmt.Println(v2)
        _ = fi24
        _ = v25
    }
    v2 = v2
    v26 := GenS1{F0: (v2 >> (((v3 - v3)) & 63))}
    switch 1 {
    case 127:
        fmt.Println(v3, v2, v3)
    }
    if (18446744073709551615 != 4294967295) {
        var v27 float32 = 0.1
        _ = v27
    } else {
        var v28 float32 = ((float32(((1.5 * -2.25) + 3.14159)) / 1.5) / 1e10)
        _ = v28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        v5 = (^(v5))
        v30 := GenS3{F0: 18446744073709551615, F1: 79}
        _ = fi29
        _ = v30
    }
    delete(v1, "a")
    var v31 int16 = -817
    if false {
        if ((true || (true && false)) && ((v31 >= v31) || false)) {
            if false {
                for fi32 := 0; fi32 < 4; fi32++ {
                    for fi33 := 0; fi33 < 20; fi33++ {
                        var v34 int32 = 2147483647
                        v35 := GenS4{F0: 694, F1: 255, F2: -32768}
                        v36 := GenS1{F0: (v2 * (v2 ^ v2))}
                        v37 := GenS1{F0: (-(v2))}
                        var v38 int = fi33
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = fi32
                }
            } else {
                v5 += v5
            }
        } else {
            v5 = 1
            _ = len(v1)
            v39 := GenS3{F0: 255, F1: 42}
            _ = v39
        }
        fmt.Println(v3, v31, v5)
        fmt.Println(v5)
    }
    v40 := GenS3{F0: (v3 * (^(v3))), F1: 178}
    var v41 bool = (313 != 42)
    var v42 int32 = 679
    v43 := GenS5{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
