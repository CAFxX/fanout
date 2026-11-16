package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 bool }
func genFn0(p0 uint8) uint32 {
    var v5 int8 = -1
    if ((float64(0.0) / -2.25) > (0.0 + float64(1.5))) {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: uint64(p0), F1: ("" + "")}
            if (!("foo bar" <= "foo bar")) {
                for fi8 := 0; fi8 < 18; fi8++ {
                    if true {
                        var v9 uint16 = 396
                        _ = v9
                    } else {
                        var v10 int16 = -321
                        if false {
                            fmt.Println(v10, fi6)
                        }
                        _ = v10
                    }
                    if true {
                        v11 := []int{}
                        v5 -= v5
                        _ = v11
                    } else {
                        if (!(fi8 >= 127)) {
                            switch "hello" {
                            case "hello":
                                p0 += 2
                            case "a":
                                p0 = 2
                            default:
                                v12 := []int{(fi8 * (fi6 + (fi8 ^ (fi8 + -1)))), fi8, (fi8 / ((fi6) | 1))}
                                _ = v12
                            }
                            var v13 int32 = 0
                            var v14 int32 = v13
                            _ = v13
                            _ = v14
                        } else {
                            v15 := []int{fi8}
                            _ = len(v15)
                            _ = v15
                        }
                        p0 = (p0 & p0)
                        v16 := []int{fi8, (-(fi8))}
                        _ = v16
                    }
                    _ = fi8
                }
                var v17 uint = 119
                _ = v17
            }
            _ = fi6
            _ = v7
        }
        v5 += (^(v5))
        var v18 float32 = (0.0 - 0.1)
        _ = v18
    } else {
        var v19 float64 = float64(((1.5 - 1.5) - ((((0.0 / (-2.25 * 1.5)) / 1.0) - (((0.1 - 0.1) / (0.1 + ((0.1 * 0.1) / float32(((1e10 + (3.14159 * -2.25)) + 0.0))))) / (3.14159 / 1.0))) - 3.14159)))
        switch "" {
        case "a":
            v20 := []int{9223372036854775807, 127}
            _ = v20
        default:
            var v21 float32 = 3.14159
            _ = v21
        }
        _ = v19
    }
    _ = p0
    _ = v5
    return 14
}
func genFn1() uint16 {
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 int8 = -1
        _ = fi22
        _ = v23
    }
    var v24 bool = (!true)
    _ = v24
    return 0
}

func main() {
    var v1 bool = ((float64((1.5 + 0.1)) * -2.25) != 1e10)
    var v2 int16 = 0
    var v3 uint32 = 764
    v4 := make(map[int]string)
    v25 := GenS1{F0: 267, F1: ("a" + "foo bar")}
    v3 = genFn0(1)
    v2 = 0
    v26 := GenS1{F0: 2, F1: "hello"}
    var v27 int32 = -1
    switch 9223372036854775807 {
    case 128:
        var v28 float64 = ((0.0 * 0.1) / 3.14159)
        _ = v28
    case 127:
        v3 = (v3 & (^(v3)))
    default:
        switch 2147483647 {
        case 1:
            v29 := GenS2{F0: 32767}
            _ = v29
        case 128:
            var v30 int16 = 1000
            _ = v30
        default:
            v31 := GenS3{F0: -591, F1: 255, F2: "\n\t"}
            _ = v31
        }
    }
    fmt.Println(v1)
    switch "a" {
    case "a":
        delete(v4, 742)
    case "":
        v32 := GenS4{F0: 324}
        _ = v32
    case "foo bar":
        var v33 int = -2147483648
        _ = v33
    default:
        var v34 float32 = ((1e10 + float32(1.5)) / (1e10 / -2.25))
        _ = v34
    }
    _ = len(v4)
    v35 := v4[0]
    if false {
        v2 = v2
    } else {
        switch 4294967295 {
        case 1:
            fmt.Println(v3, v3, v27)
        case 127:
            v35 = "a"
        }
    }
    var v36 uint8 = 1
    fmt.Println(v27)
    if v1 {
        if (9223372036854775807 < 9223372036854775807) {
            v3 += (v3 % (((v3 * 852)) | 1))
            fmt.Println(v36, v35)
            var v37 uint16 = 65535
            _ = v37
        }
    } else {
        v3 = 962
        fmt.Println(v1)
    }
    var v38 float32 = 1e10
    v39 := GenS5{F0: "foo bar", F1: (!(478 != 18446744073709551615))}
    for fi40 := 0; fi40 < 14; fi40++ {
        v41 := GenS3{F0: 9223372036854775807, F1: 42, F2: ""}
        _ = len(v4)
        _ = fi40
        _ = v41
    }
    v3 -= (v3 | (v3 ^ 65535))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
}
