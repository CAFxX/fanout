package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint8) string {
    switch "\n\t" {
    case "":
        v6 := GenS1{F0: 52}
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: 2147483647}
        _ = v7
    case "hello":
        fmt.Println(p0)
    default:
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := []int{43, (fi8 & fi8), (-(fi8)), (fi8 | (-(fi8)))}
            switch fi8 {
            case 477:
                if true {
                    v10 := GenS2{F0: 731, F1: 1, F2: 100}
                    v11 := GenS2{F0: 42, F1: -2147483648, F2: -128}
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS2{F0: 0, F1: 1, F2: -128}
                    var v13 int8 = 0
                    _ = v12
                    _ = v13
                }
            case -318:
                var v14 uint64 = 4294967295
                _ = v14
            }
            p0 = (p0 >> ((309) & 63))
            _ = fi8
            _ = v9
        }
    }
    var v15 int16 = 1
    for fi16 := 0; fi16 < 17; fi16++ {
        v15 += (v15 % (((v15 % ((v15) | 1))) | 1))
        p0 += (p0 ^ p0)
        v17 := GenS3{F0: 874, F1: (2147483647 >= 616)}
        _ = fi16
        _ = v17
    }
    v18 := []int{422}
    _ = p0
    _ = v15
    _ = v18
    return ""
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 2147483647
    v3 := []string{}
    var v4 int32 = -1
    v5 := make(map[int]bool)
    v19 := GenS2{F0: v4, F1: (v4 >> ((87) & 63)), F2: -128}
    if false {
        var v20 float32 = 1.5
        v4 += (v4 ^ (-(v4)))
        var v21 bool = (((((1e10 / v20) * float32((((float64((((1.5 - (v20 * v20)) * ((1.5 * 1.5) + 0.0)) * 1.5)) * (((-2.25 - 1.5) + (float64(0.0) + 3.14159)) * 1e10)) / float64(v20)) + float64((v20 - (-2.25 * (((v20 - float32(3.14159)) + 3.14159) - float32(((0.0 * 0.0) * (0.0 * 3.14159)))))))))) * v20) / ((((1.5 * -2.25) * ((float32((1e10 - (float64(v20) / (((((0.0 - -2.25) / (float64(-2.25) * float64((v20 - (v20 - v20))))) - 0.1) + -2.25) * 0.0)))) + (3.14159 - (v20 - float32((0.1 * 0.0))))) / v20)) * ((((((v20 - float32(-2.25)) + v20) * ((1.5 + ((v20 / v20) / v20)) + v20)) / ((v20 - (v20 / v20)) + (v20 * v20))) * 1.5) - v20)) + v20)) >= 3.14159)
        _ = v20
        _ = v21
    } else {
        v2 = v2
    }
    var v22 uint32 = 63
    v23 := GenS4{F0: 711}
    var v24 uint32 = (v22 & (^(v22)))
    if len(v3) > 0 {
        v3[0] = ""
    }
    switch "\n\t" {
    case "foo bar":
        var v25 uint16 = 1
        _ = v25
    case "":
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := GenS5{F0: 42, F1: (v4 / ((v4) | 1))}
            _ = fi26
            _ = v27
        }
    case "hello":
        if (!(int16(v4) > 32767)) {
            switch "foo bar" {
            case "a":
                var v28 uint32 = 0
                _ = v28
            }
        }
    default:
        var v29 int8 = -1
        _ = v29
    }
    fmt.Println(v22, v4)
    if false {
        var v30 uint = 255
        v4 += (v4 / ((-2147483648) | 1))
        var v31 int64 = 9223372036854775807
        _ = v30
        _ = v31
    } else {
        var v32 int32 = v4
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 int = (fi33 - (v2 | fi33))
        if true {
            v24 = (v24 / (((v22 | (v24 ^ v24))) | 1))
            var v35 int16 = 1000
            _ = v35
        } else {
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 float64 = 0.1
                fmt.Println(v34)
                _ = fi36
                _ = v37
            }
        }
        _ = fi33
        _ = v34
    }
    var v38 int32 = 2147483647
    if true {
        v39 := GenS2{F0: (v38 % (((v38 << ((65535) & 63))) | 1)), F1: 0, F2: 42}
        if (false && (42 != 9223372036854775807)) {
            var v40 int16 = 42
            _ = v40
        } else {
            var v41 uint16 = 969
            v2 += v2
            _ = v41
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            _ = len(v3)
            _ = fi42
        }
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
