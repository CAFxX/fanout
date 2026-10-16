package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint; F2 int64 }
type GenS5 struct { F0 uint8; F1 uint; F2 bool }
type GenS6 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]string)
    var v3 float32 = (-2.25 / 1.0)
    var v4 float64 = 1.5
    switch "a" {
    case "foo bar":
        v5 := []int{4294967295}
        _ = v5
    case "\n\t":
        switch "" {
        case "a":
            switch "a" {
            case "":
                v4 -= v4
            case "hello":
                v6 := v2[169]
                _ = v6
            case "\n\t":
                for fi7 := 0; fi7 < 6; fi7++ {
                    v8 := []int{}
                    _ = fi7
                    _ = v8
                }
            }
        case "\n\t":
            switch 128 {
            case 32768:
                v4 = ((v4 + v4) - float64(v3))
            default:
                var v9 float64 = float64((0.1 - 3.14159))
                _ = v9
            }
        case "hello":
            v4 += v4
        }
    case "":
        v1 = v1
    }
    fmt.Println(v1, v3, v3)
    var v10 float32 = (1e10 - (float32(v4) - (-2.25 * v3)))
    var v11 int = -158
    v12 := GenS1{F0: 4294967295, F1: 1, F2: (v1 | (v1 & (v1 ^ -551)))}
    fmt.Println(v10, v1, v11)
    v13 := GenS2{F0: 0, F1: (915 >= 0)}
    var v14 int16 = 32767
    v14 += int16(v11)
    fmt.Println(v14, v4, v11)
    var v15 float32 = ((v10 / 1e10) + (v3 + ((v10 * v3) * (1.5 * -2.25))))
    switch ("\n\t" + "hello") {
    case "":
        if false {
            v16 := GenS3{F0: -104}
            _ = v16
        }
    case "a":
        switch ("" + ("foo bar" + "foo bar")) {
        case "":
            v1 = (^(v1))
        case "a":
            var v17 bool = false
            _ = v17
        default:
            switch "hello" {
            case "foo bar":
                switch "" {
                case "":
                    v1 = -1
                case "foo bar":
                    v18 := GenS1{F0: 40, F1: 4294967295, F2: v1}
                    _ = v18
                case "a":
                    var v19 int16 = 0
                    _ = v19
                }
            default:
                if false {
                    v15 += (v15 * (1.5 / -2.25))
                    v10 += v15
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if (2 <= 2) {
                            var v21 uint16 = 386
                            v22 := GenS1{F0: 42, F1: 1, F2: v1}
                            v11 += v11
                            _ = v21
                            _ = v22
                        }
                        v10 += 1.5
                        _ = fi20
                    }
                } else {
                    if (9223372036854775807 != -1) {
                        var v23 int64 = 42
                        v24 := GenS4{F0: 2147483647, F1: 65535, F2: (v23 >> ((0) & 63))}
                        v2[4294967295] = "\n\t"
                        _ = v23
                        _ = v24
                    } else {
                        if true {
                            v15 = v3
                            v25 := GenS3{F0: -1}
                            _ = v25
                        }
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    }
                    if true {
                        v27 := GenS1{F0: 633, F1: 0, F2: (v1 * (v1 - -423))}
                        v3 -= (float32(v4) * (float32(-2.25) + (0.1 / v3)))
                        _ = v27
                    } else {
                        var v28 float64 = 1e10
                        _ = v28
                    }
                }
            }
        }
    default:
        switch v11 {
        case -574:
            fmt.Println(v15, v14, v4)
        default:
            if true {
                v29 := GenS3{F0: 127}
                for fi30 := 0; fi30 < 10; fi30++ {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        fmt.Println(v14, v11)
                        fmt.Println(v14, v1)
                        _ = fi31
                    }
                    v1 = (v1 & -1)
                    _ = fi30
                }
                _ = v29
            } else {
                v1 += v1
            }
        }
    }
    switch "a" {
    case "a":
        if (1 > 42) {
            var v32 float64 = v4
            v33 := []int{v11, 127, v11, v11}
            v14 = (v14 >> ((0) & 63))
            _ = v32
            _ = v33
        } else {
            var v34 float32 = 0.1
            _ = v34
        }
    case "\n\t":
        v35 := GenS5{F0: 0, F1: 1, F2: (!false)}
        _ = v35
    default:
        if ((true && true) && true) {
            var v36 uint32 = 2147483647
            fmt.Println(v1, v10)
            fmt.Println(v15)
            _ = v36
        } else {
            v15 -= 1e10
            fmt.Println(v15, v3)
            v37 := GenS6{F0: v1, F1: 2}
            _ = v37
        }
    }
    if false {
        v1 = (^(v1))
    } else {
        var v38 int64 = 9223372036854775807
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
