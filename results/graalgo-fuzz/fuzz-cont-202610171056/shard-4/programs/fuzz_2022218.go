package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint16; F1 string; F2 uint64 }
func genFn0() int64 {
    if true {
        var v4 float64 = 1.5
        fmt.Println(v4)
        v5 := []int{4294967295, 0, -457, 0}
        _ = v4
        _ = v5
    } else {
        var v6 uint32 = 65535
        _ = v6
    }
    var v7 int = 2147483647
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 127
    v2 := []string{}
    var v3 int32 = 2147483647
    switch "a" {
    case "":
        var v8 uint16 = 0
        _ = v8
    case "\n\t":
        if ((!(!false)) || (0 < genFn0())) {
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: (v3 % (((-(v3))) | 1)), F1: (v1 >= -28), F2: true}
                v3 += v3
                switch fi9 {
                case 1:
                    var v11 uint64 = 2
                    _ = v11
                case -801:
                    v3 += v3
                default:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        fmt.Println(fi12, v1, v1)
                        _ = fi12
                    }
                }
                _ = fi9
                _ = v10
            }
            var v13 bool = true
            v1 += v1
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 20; fi14++ {
            if (-720 >= -228) {
                var v15 int8 = v1
                _ = v15
            }
            _ = fi14
        }
    }
    v1 += (v1 * 100)
    v3 -= (-(v3))
    var v16 int = 0
    v3 = 0
    v17 := GenS2{F0: v3, F1: 127, F2: 366}
    var v18 uint64 = 1
    if len(v2) > 0 {
        v2[0] = "a"
    }
    switch 549 {
    case -229:
        fmt.Println(v1, v16)
    case 9223372036854775807:
        var v19 uint16 = uint16(v18)
        _ = v19
    case -9223372036854775808:
        v20 := GenS2{F0: 1, F1: (v1 - v1), F2: 0}
        _ = v20
    }
    var v21 uint = 42
    switch (v16 << ((778) & 63)) {
    case 4294967295:
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 float32 = float32(((-2.25 / (((((3.14159 * (3.14159 * (float64(1.5) / 1.0))) / 1.0) - float64((0.1 / (1.5 / float32(1.5))))) * 0.1) + float64(1e10))) / 0.1))
            for fi24 := 0; fi24 < 15; fi24++ {
                v25 := GenS1{F0: (v3 + (v3 & (v3 * -2147483648))), F1: true, F2: (0 <= v21)}
                fmt.Println(v23, v16, fi22)
                _ = fi24
                _ = v25
            }
            switch "foo bar" {
            case "foo bar":
                var v26 uint = v21
                _ = v26
            case "hello":
                switch v16 {
                case 1:
                    var v27 uint32 = 4294967295
                    _ = v27
                }
            case "a":
                var v28 int64 = genFn0()
                _ = v28
            default:
                var v29 int = 9223372036854775807
                _ = v29
            }
            _ = fi22
            _ = v23
        }
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            fmt.Println(v18)
            v3 = v3
            switch "foo bar" {
            case "a":
                var v32 uint = v21
                _ = v32
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            fmt.Println(v18)
            _ = fi31
        }
        switch (v16 + (v16 % ((fi30) | 1))) {
        case 1:
            if (!true) {
                var v33 int16 = 522
                v1 += -128
                _ = v33
            }
        case 2147483647:
            v34 := GenS3{F0: false, F1: v1}
            _ = v34
        case 127:
            for fi35 := 0; fi35 < 8; fi35++ {
                var v36 float64 = 0.1
                v3 = 2147483647
                if (65535 != 0) {
                    var v37 int8 = v1
                    _ = v37
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v38 uint = v21
                        _ = v38
                    case "\n\t":
                        fmt.Println(v36, v1)
                    default:
                        if true {
                            v39 := GenS4{F0: 217, F1: "\n\t", F2: 18446744073709551615}
                            var v40 string = "a"
                            switch v40 {
                            case "a":
                                v2 = append(v2, v40)
                            case "hello":
                                var v41 int16 = 32767
                                _ = v41
                            }
                            v2 = append(v2, "foo bar")
                            var v42 float32 = 1e10
                            _ = v39
                            _ = v40
                            _ = v42
                        } else {
                            fmt.Println(v16, v36, fi35)
                        }
                    }
                    switch "hello" {
                    case "\n\t":
                        var v43 int32 = (-(v3))
                        _ = v43
                    case "hello":
                        switch (fi35 + (fi35 >> (((v21 % (((^(v21))) | 1))) & 63))) {
                        case 433:
                            var v44 string = "a"
                            _ = v44
                        case -9223372036854775808:
                            v45 := GenS3{F0: (1 == 42), F1: v1}
                            _ = v45
                        case 687:
                            v1 = -7
                        default:
                            var v46 int8 = (v1 | v1)
                            _ = v46
                        }
                    default:
                        var v47 uint = 207
                        _ = v47
                    }
                }
                _ = fi35
                _ = v36
            }
        default:
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
        v21 -= (v21 - v21)
        _ = fi30
    }
    v21 = v21
    v1 += (v1 << ((v21) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
