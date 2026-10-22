package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint64; F2 bool }
type GenS6 struct { F0 int64; F1 int32; F2 bool }
func genFn0(p0 int) int {
    var v5 uint8 = 128
    if false {
        v6 := GenS1{F0: -1, F1: 493, F2: 870}
        _ = v6
    } else {
        v7 := []int{(p0 << ((437) & 63)), p0, (p0 ^ (p0 + (-(p0)))), (p0 >> ((2) & 63))}
        for fi8 := 0; fi8 < 17; fi8++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi8
        }
        _ = v7
    }
    _ = p0
    _ = v5
    return p0
}
func genFn1(p0 int32) uint16 {
    if (false || false) {
        p0 = (p0 | (p0 / ((0) | 1)))
        var v9 uint16 = 1
        v9 -= v9
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS2{F0: 65535, F1: 752, F2: 32767}
        p0 = (-(p0))
        _ = fi10
        _ = v11
    }
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 float32 = 0.0
    var v2 uint = 255
    v3 := []string{("\n\t" + ""), "foo bar", "\n\t", "hello"}
    var v4 int32 = 42
    var v12 int8 = 58
    v1 -= v1
    var v13 int8 = v12
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (1000 < 1000) {
        var v14 uint8 = 0
        v15 := GenS3{F0: (v14 / ((v14) | 1)), F1: false}
        v16 := GenS4{F0: 279}
        _ = v14
        _ = v15
        _ = v16
    } else {
        switch 2147483647 {
        case 2147483647:
            var v17 uint64 = 1
            _ = v17
        }
        var v18 int16 = 0
        _ = v18
    }
    v19 := GenS5{F0: (v4 + (-(v4))), F1: 483, F2: (988 > 409)}
    v2 = (v2 >> (((v2 ^ (^(v2)))) & 63))
    for fi20 := 0; fi20 < 13; fi20++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS5{F0: (v4 * (v4 / (((v4 / ((1) | 1))) | 1))), F1: 18446744073709551615, F2: false}
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 uint32 = 1
    for fi24 := 0; fi24 < 4; fi24++ {
        _ = len(v3)
        var v25 float32 = v1
        var v26 int = (fi24 & fi24)
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(v12, v13)
        _ = fi27
    }
    v28 := GenS6{F0: 42, F1: (v4 * -1), F2: true}
    if (float32(float64(float32(3.14159))) < float32(0.1)) {
        v2 += v2
        _ = len(v3)
        v12 += v13
    } else {
        fmt.Println(v4)
    }
    var v29 int = 128
    var v30 string = "a"
    if (v30 != ((v30 + v30) + (("\n\t" + (v30 + (("" + "") + v30))) + (v30 + "")))) {
        for fi31 := 0; fi31 < 4; fi31++ {
            v1 += v1
            for fi32 := 0; fi32 < 1; fi32++ {
                if len(v3) > 0 {
                    v3[0] = v30
                }
                if false {
                    if true {
                        switch v30 {
                        case "":
                            fmt.Println(fi31, v4, fi31)
                        case "foo bar":
                            switch (("\n\t" + v30) + "foo bar") {
                            case "\n\t":
                                var v33 int = (fi31 << (((v2 & uint(v23))) & 63))
                                _ = v33
                            default:
                                _ = len(v3)
                            }
                        }
                        v2 = (v2 >> (((v2 | v2)) & 63))
                    } else {
                        v4 += 0
                    }
                }
                _ = fi32
            }
            _ = fi31
        }
    }
    v13 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
