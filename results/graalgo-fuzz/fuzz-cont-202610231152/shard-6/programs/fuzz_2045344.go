package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
type GenS3 struct { F0 int8; F1 bool; F2 int32 }

func main() {
    var v1 int16 = 1
    var v2 uint16 = 65535
    v3 := []int{-920, 127, 42, 0}
    var v4 uint64 = 1
    v5 := GenS1{F0: 152}
    v2 -= v2
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int = (^(fi6))
        var v8 float64 = -2.25
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 119
    if false {
        var v10 int32 = 2147483647
        var v11 int8 = 100
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            if ((959 != 32767) && true) {
                if false {
                    if true {
                        v13 := GenS1{F0: 2147483648}
                        var v14 float32 = 3.14159
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        _ = v13
                        _ = v14
                    } else {
                        v1 = (v1 << ((171) & 63))
                        v2 = 42
                    }
                }
            } else {
                if (((((0.0 - (1.5 * 3.14159)) / 3.14159) * -2.25) / 1e10) > (float64(0.0) * -2.25)) {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        var v16 string = "foo bar"
                        _ = len(v3)
                        _ = fi15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 17; fi17++ {
                        fmt.Println(v9, v2, v9)
                        _ = fi17
                    }
                } else {
                    if false {
                        _ = len(v3)
                    } else {
                        var v18 float64 = (((((((1.5 * (0.1 - -2.25)) / (((3.14159 + float64(float32((((1e10 * 1.5) - (-2.25 - 3.14159)) / (float64((-2.25 / (1e10 / 0.1))) - 0.1))))) - float64(3.14159)) - 1e10)) + 0.0) * 1e10) + -2.25) * (1.5 / 0.1)) + 1e10)
                        for fi19 := 0; fi19 < 9; fi19++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi19
                        }
                        _ = v18
                    }
                    v9 += v9
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v2 -= v2
                        fmt.Println(v2)
                        _ = fi20
                    }
                }
                var v21 uint8 = (^(v9))
                _ = v21
            }
            var v22 uint16 = (^(v2))
            v23 := GenS2{F0: "\n\t", F1: 189, F2: 0}
            _ = fi12
            _ = v22
            _ = v23
        }
        switch "foo bar" {
        case "foo bar":
            switch "\n\t" {
            case "":
                v24 := GenS3{F0: 42, F1: true, F2: 1}
                _ = v24
            case "\n\t":
                fmt.Println(v2)
            case "a":
                fmt.Println(v1, v9)
            }
        }
        var v25 uint8 = v9
        _ = v25
    }
    v1 -= v1
    v2 += (^(v2))
    v26 := GenS1{F0: 104}
    v9 += v9
    var v27 uint64 = v4
    v27 = (v4 / ((18446744073709551615) | 1))
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 float32 = 3.14159
        for fi30 := 0; fi30 < 17; fi30++ {
            if (128 != v9) {
                v9 = (^(v9))
                v29 -= float32((1e10 + (0.0 + (-2.25 * -2.25))))
                v27 = 1
            }
            var v31 int64 = 42
            fmt.Println(fi30)
            var v32 float32 = v29
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    switch "\n\t" {
    case "\n\t":
        switch -41 {
        case 32768:
            fmt.Println(v4, v9)
        default:
            var v33 float32 = 0.1
            _ = v33
        }
    case "hello":
        fmt.Println(v1, v2, v27)
    case "foo bar":
        var v34 int8 = -124
        _ = v34
    }
    v3 = append(v3, -9223372036854775808)
    var v35 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
