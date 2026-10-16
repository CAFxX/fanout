package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint }
func genFn0(p0 int16) int32 {
    v6 := []int{32767, 1, 2147483647}
    p0 -= 42
    _ = p0
    _ = v6
    return -402
}

func main() {
    v1 := []int{}
    var v2 string = "hello"
    var v3 int = 42
    var v4 int8 = 0
    var v5 uint16 = 0
    switch (v2 + v2) {
    case "\n\t":
        fmt.Println(v2, v5)
    case "foo bar":
        var v7 string = v2
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 6; fi8++ {
            switch v3 {
            case 32768:
                for fi9 := 0; fi9 < 2; fi9++ {
                    if (!false) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v10 float64 = ((((-2.25 * 0.0) + (-2.25 - (1.5 - (1.5 - 1.5)))) - 1e10) / 1.5)
                        _ = v10
                    }
                    var v11 int32 = 42
                    _ = fi9
                    _ = v11
                }
            case 2147483647:
                for fi12 := 0; fi12 < 20; fi12++ {
                    v2 = (v2 + v2)
                    var v13 uint64 = 447
                    _ = fi12
                    _ = v13
                }
            case -9223372036854775808:
                var v14 uint = 592
                _ = v14
            }
            var v15 uint16 = v5
            _ = fi8
            _ = v15
        }
    }
    fmt.Println(v4, v5)
    if len(v1) > 0 {
        v1[0] = 42
    }
    v4 -= v4
    switch v2 {
    case "\n\t":
        v1 = append(v1, 42)
    case "":
        v3 -= -146
    }
    if len(v1) > 0 {
        v1[0] = (v3 / ((2147483647) | 1))
    }
    v16 := GenS1{F0: 288, F1: true}
    v5 += v5
    v17 := GenS1{F0: 577, F1: (1 >= 2)}
    if false {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int16 = 42
            v5 = (v5 | (v5 % ((v5) | 1)))
            v20 := GenS2{F0: (v19 ^ (^(v19))), F1: 42, F2: 268}
            _ = fi18
            _ = v19
            _ = v20
        }
        v4 += v4
        if (!false) {
            var v21 uint64 = 2
            if (v5 != 1) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v23 := GenS3{F0: 18446744073709551615}
                    v2 = v2
                    if (42 <= 2) {
                        v2 = (v2 + (("" + "hello") + v2))
                        switch "a" {
                        case "hello":
                            fmt.Println(v21)
                        case "a":
                            var v24 uint64 = 0
                            _ = v24
                        case "foo bar":
                            v25 := GenS2{F0: -32768, F1: 42, F2: (v21 | (^(v21)))}
                            _ = v25
                        default:
                            v4 -= v4
                        }
                        v26 := GenS4{F0: 127}
                        _ = v26
                    } else {
                        if true {
                            if (89 == 0) {
                                var v27 uint16 = v5
                                _ = v27
                            }
                            var v28 float64 = 1e10
                            var v29 int32 = -611
                            _ = v28
                            _ = v29
                        }
                        switch (fi22 / (((v3 / (((v3 + 32768)) | 1))) | 1)) {
                        case 0:
                            var v30 int16 = -1
                            _ = v30
                        case 42:
                            for fi31 := 0; fi31 < 13; fi31++ {
                                v4 -= -128
                                v21 += (v21 << ((4294967295) & 63))
                                v5 = 65535
                                _ = fi31
                            }
                        case 9223372036854775807:
                            v1 = append(v1, fi22)
                        }
                        _ = len(v1)
                    }
                    _ = fi22
                    _ = v23
                }
            } else {
                v21 -= (v21 >> ((398) & 63))
                v32 := GenS1{F0: -909, F1: ((v4 == v4) || (!false))}
                fmt.Println(v3)
                _ = v32
            }
            _ = v21
        } else {
            v1 = append(v1, v3)
            v2 = v2
            var v33 uint16 = (v5 << ((65535) & 63))
            _ = v33
        }
    }
    v1 = append(v1, -2147483648)
    v34 := GenS5{F0: v3}
    for fi35 := 0; fi35 < 14; fi35++ {
        switch v2 {
        case "\n\t":
            var v36 uint8 = 255
            _ = v36
        case "a":
            v37 := GenS6{F0: ("hello" + (v2 + v2)), F1: 4294967295}
            _ = v37
        default:
            var v38 string = v2
            _ = v38
        }
        var v39 int64 = 46
        fmt.Println(fi35, v39)
        _ = fi35
        _ = v39
    }
    v4 = (^(v4))
    var v40 uint64 = 113
    v3 = v3
    v4 = (v4 << ((2) & 63))
    var v41 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
