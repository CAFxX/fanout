package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint16; F2 int }
func genFn0() uint64 {
    var v6 uint32 = 325
    v7 := []int{15}
    var v8 bool = false
    v8 = (18446744073709551615 < 654)
    for fi9 := 0; fi9 < 1; fi9++ {
        if v8 {
            fmt.Println(fi9)
            switch "foo bar" {
            case "a":
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 uint16 = 65535
                    var v12 int32 = -2147483648
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        }
        _ = fi9
    }
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 string, p1 int8) uint {
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS1{F0: 362, F1: 42, F2: 255}
            v16 := GenS2{F0: -887, F1: 1, F2: 213}
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := []int{(fi13 * fi13), 1}
        _ = fi13
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            p0 = "a"
            fmt.Println(p1)
            var v20 bool = true
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        p0 = "hello"
        var v22 int8 = p1
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    return 821
}

func main() {
    var v1 uint16 = 378
    v2 := []uint32{42, 219}
    var v3 uint16 = (v1 / (((v1 & (^(v1)))) | 1))
    v4 := make(map[string]int)
    var v5 int8 = 1
    var v23 int8 = (^(v5))
    if true {
        fmt.Println(v3)
        switch "a" {
        case "hello":
            if (true || true) {
                v23 = 42
                v2 = append(v2, 137)
            }
        case "a":
            v23 += v23
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v24 int32 = 42
            _ = v24
        }
        if ((!false) && false) {
            v25 := GenS3{F0: 327, F1: 0}
            v4[("a" + "hello")] = 9223372036854775807
            var v26 int32 = 0
            _ = v25
            _ = v26
        }
    }
    var v27 int = -27
    switch "\n\t" {
    case "":
        var v28 int = v27
        _ = v28
    case "a":
        var v29 int32 = 61
        _ = v29
    case "foo bar":
        v30 := GenS4{F0: 715}
        _ = v30
    }
    v31 := GenS5{F0: -787, F1: -1}
    if false {
        v5 += (-(v23))
        var v32 float64 = 1.5
        _ = v32
    } else {
        v3 += v1
        v27 -= (v27 * (v27 / ((v27) | 1)))
    }
    var v33 uint32 = 42
    v33 -= v33
    if true {
        fmt.Println(v27)
    } else {
        var v34 int = (v27 >> ((197) & 63))
        var v35 uint32 = v33
        switch v27 {
        case 973:
            var v36 bool = false
            _ = v36
        case -342:
            switch (v27 * v34) {
            case -692:
                if false {
                    var v37 string = "foo bar"
                    v23 -= (v5 % (((v23 ^ 0)) | 1))
                    _ = v37
                }
            case 32767:
                var v38 float64 = 1e10
                _ = v38
            case 9223372036854775807:
                var v39 int32 = 0
                _ = v39
            default:
                if true {
                    var v40 string = "hello"
                    for fi41 := 0; fi41 < 10; fi41++ {
                        if len(v2) > 0 {
                            v2[0] = (v35 & (v35 - (v35 << ((1) & 63))))
                        }
                        v27 = (v34 | fi41)
                        for fi42 := 0; fi42 < 11; fi42++ {
                            v35 = (v35 ^ v35)
                            _ = fi42
                        }
                        _ = fi41
                    }
                    if true {
                        var v43 int8 = (v5 | (v5 * (v23 ^ 127)))
                        _ = v43
                    } else {
                        v44 := GenS6{F0: v1, F1: 32767, F2: v34}
                        _ = v44
                    }
                    _ = v40
                }
            }
        default:
            for fi45 := 0; fi45 < 16; fi45++ {
                _ = len(v4)
                _ = fi45
            }
        }
        _ = v34
        _ = v35
    }
    fmt.Println(v27, v33)
    var v46 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v46
    _ = v5
}
