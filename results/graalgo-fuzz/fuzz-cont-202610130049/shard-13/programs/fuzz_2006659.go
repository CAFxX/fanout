package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    v5 := GenS1{F0: -9223372036854775808}
    v6 := []int{1, 4294967295}
    _ = v5
    _ = v6
    return 424
}
func genFn1(p0 uint8, p1 uint32, p2 uint) string {
    switch "\n\t" {
    case "foo bar":
        v7 := GenS1{F0: 42}
        _ = v7
    default:
        if true {
            v8 := []int{-790}
            fmt.Println(p1)
            p0 += 16
            _ = v8
        }
    }
    p0 = (^(p0))
    for fi9 := 0; fi9 < 5; fi9++ {
        switch "foo bar" {
        case "a":
            if false {
                v10 := GenS1{F0: 938}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    if ((^(fi9)) < fi11) {
                        var v12 int8 = 0
                        _ = v12
                    }
                    var v13 float64 = (-2.25 - (1.5 / float64(3.14159)))
                    _ = fi11
                    _ = v13
                }
            }
        default:
            p1 -= uint32(p0)
        }
        switch (-(fi9)) {
        case 769:
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 int64 = -9223372036854775808
                var v16 float32 = 0.1
                _ = fi14
                _ = v15
                _ = v16
            }
        case -1:
            if ((p0 << ((p2) & 63)) != p0) {
                p2 = p2
                p2 = p2
                if false {
                    v17 := GenS1{F0: 9223372036854775807}
                    for fi18 := 0; fi18 < 12; fi18++ {
                        var v19 uint = (p2 - uint(p1))
                        _ = fi18
                        _ = v19
                    }
                    if (32767 == -32768) {
                        var v20 float32 = (((1e10 + -2.25) * (((3.14159 + 1e10) * 3.14159) / (0.0 + 0.1))) * (((0.1 * (-2.25 / 1e10)) + 0.0) * 1.5))
                        _ = v20
                    }
                    _ = v17
                } else {
                    p2 -= (p2 | (p2 | p2))
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 int32 = -456
                        _ = fi21
                        _ = v22
                    }
                }
            }
        case 32768:
            v23 := GenS1{F0: genFn0()}
            _ = v23
        default:
            p2 += (p2 * (p2 / ((uint(p1)) | 1)))
        }
        var v24 uint32 = p1
        _ = fi9
        _ = v24
    }
    p2 += p2
    p0 -= 1
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    v1 := []uint32{2, 65535, 412}
    var v2 float64 = 0.0
    var v3 uint32 = 4294967295
    var v4 float64 = 0.1
    v25 := GenS1{F0: -1}
    switch -850 {
    case -484:
        if len(v1) > 0 {
            v1[0] = v3
        }
    case 128:
        var v26 int64 = 9223372036854775807
        _ = v26
    case 9223372036854775807:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 = (0.0 / v2)
    v27 := GenS2{F0: 255, F1: 0}
    for fi28 := 0; fi28 < 15; fi28++ {
        switch (fi28 & (fi28 / (((fi28 * (fi28 << ((18446744073709551615) & 63)))) | 1))) {
        case 332:
            v29 := GenS2{F0: 42, F1: 4294967295}
            _ = v29
        case 438:
            for fi30 := 0; fi30 < 7; fi30++ {
                v3 -= 147
                _ = fi30
            }
        }
        _ = fi28
    }
    if false {
        switch ("" + "a") {
        case "foo bar":
            var v31 uint16 = 32767
            _ = v31
        case "hello":
            if false {
                v1 = append(v1, 42)
                var v32 uint8 = 2
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v32
            }
        case "":
            fmt.Println(v3)
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            _ = len(v1)
            v34 := GenS3{F0: v3}
            _ = fi33
            _ = v34
        }
    } else {
        fmt.Println(v4, v3, v3)
        var v35 int32 = -564
        _ = v35
    }
    fmt.Println(v3, v2)
    v4 = (0.1 / 1.5)
    for fi36 := 0; fi36 < 8; fi36++ {
        switch "foo bar" {
        case "hello":
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 bool = (500 >= 42)
                var v39 uint = 2
                switch fi36 {
                case 4294967295:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    fmt.Println(v38, v38, fi37)
                }
                _ = fi37
                _ = v38
                _ = v39
            }
        case "foo bar":
            var v40 int = (fi36 & fi36)
            _ = v40
        case "":
            if (1.5 >= (((1.5 * 0.1) * 1.5) - (0.1 / float32(v2)))) {
                var v41 int32 = 54
                var v42 uint64 = 440
                _ = v41
                _ = v42
            }
        }
        v4 -= ((1e10 - v2) / (v4 / v4))
        _ = fi36
    }
    var v43 int64 = 42
    if len(v1) > 0 {
        v1[0] = v3
    }
    v4 = ((float64((0.1 / float32((v4 - 3.14159)))) * v2) + v4)
    var v44 string = "foo bar"
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
