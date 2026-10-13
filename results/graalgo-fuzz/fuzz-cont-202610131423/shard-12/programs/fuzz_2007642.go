package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int32; F2 string }
type GenS5 struct { F0 bool }
func genFn0(p0 uint8, p1 int16, p2 uint16) uint8 {
    var v7 uint = 489
    fmt.Println(p1, p1, p2)
    var v8 float64 = 0.0
    var v9 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 int32, p1 int16, p2 int64) uint8 {
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(p2, p2)
        v11 := []int{32767, (fi10 * (fi10 * (fi10 & 42))), 128}
        fmt.Println(p0, p1)
        _ = fi10
        _ = v11
    }
    v12 := []int{int(p0), 145}
    p1 -= 0
    for fi13 := 0; fi13 < 3; fi13++ {
        p2 -= (p2 - -9223372036854775808)
        if true {
            p0 -= p0
            v12 = append(v12, (fi13 << ((18446744073709551615) & 63)))
            p1 -= int16(fi13)
        } else {
            var v14 int32 = (p0 & (p0 % ((p0) | 1)))
            fmt.Println(p1)
            switch fi13 {
            case 128:
                var v15 int64 = 585
                _ = v15
            case -826:
                fmt.Println(v14, v14, fi13)
            case 42:
                v16 := GenS1{F0: (p2 * (p2 << ((0) & 63))), F1: 240, F2: v14}
                _ = v16
            }
            _ = v14
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 50
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 int64 = 541
    var v5 string = "\n\t"
    var v6 int64 = v1
    var v17 int64 = (^(v4))
    v18 := GenS1{F0: v1, F1: 1, F2: 710}
    if (32768 == 825) {
        v3 = ((1 < -452) || (3.14159 != (0.0 * 0.0)))
    }
    if v3 {
        v3 = (100 > 45)
        switch -9223372036854775808 {
        case -80:
            switch -267 {
            case 128:
                var v19 float32 = -2.25
                _ = v19
            case 4294967295:
                for fi20 := 0; fi20 < 18; fi20++ {
                    v17 -= v1
                    if ("hello" == "foo bar") {
                        v21 := GenS2{F0: v6}
                        v22 := GenS3{F0: 0, F1: 1}
                        _ = v21
                        _ = v22
                    }
                    var v23 int = 32768
                    _ = fi20
                    _ = v23
                }
            default:
                if v3 {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if (v3 && true) {
                            v25 := []int{}
                            _ = v25
                        }
                        switch (fi24 & fi24) {
                        case -2147483648:
                            var v26 uint32 = 1000
                            _ = v26
                        case 32768:
                            v27 := GenS3{F0: 2, F1: 127}
                            _ = v27
                        }
                        v17 += (v1 & (v4 + -9223372036854775808))
                        _ = fi24
                    }
                    switch v5 {
                    case "\n\t":
                        fmt.Println(v1)
                    case "a":
                        v28 := GenS2{F0: v6}
                        _ = v28
                    case "hello":
                        switch 347 {
                        case -156:
                            fmt.Println(v17)
                        case 9223372036854775807:
                            if true {
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    v30 := []int{(fi29 << ((642) & 63)), fi29, 9223372036854775807}
                                    _ = fi29
                                    _ = v30
                                }
                                v31 := v2[2147483647]
                                _ = v31
                            } else {
                                if v3 {
                                    delete(v2, 2147483647)
                                    for fi32 := 0; fi32 < 3; fi32++ {
                                        v33 := []int{(fi32 ^ -376), (fi32 % ((fi32) | 1)), 42}
                                        for fi34 := 0; fi34 < 20; fi34++ {
                                            fmt.Println(fi34)
                                            _ = len(v33)
                                            _ = fi34
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                }
                            }
                        }
                    default:
                        var v35 int64 = (^(v17))
                        _ = v35
                    }
                }
            }
        case -9223372036854775808:
            for fi36 := 0; fi36 < 9; fi36++ {
                v1 = v17
                _ = fi36
            }
        }
        v37 := GenS4{F0: 255, F1: 42, F2: "hello"}
        _ = v37
    } else {
        v1 += (v17 << ((700) & 63))
        var v38 uint16 = 710
        v2[2147483647] = 2147483647
        _ = v38
    }
    var v39 int8 = 83
    v40 := GenS5{F0: (65535 > 255)}
    v6 -= (v1 << ((18446744073709551615) & 63))
    v4 -= 42
    switch -965 {
    case 44:
        v41 := GenS4{F0: 524, F1: 992, F2: v5}
        _ = v41
    case -2147483648:
        var v42 uint = 4294967295
        _ = v42
    case 1:
        v1 = v1
    }
    var v43 uint16 = 0
    fmt.Println(v43, v43)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
