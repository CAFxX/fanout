package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    var v1 uint64 = 2
    var v2 int32 = 0
    var v3 int32 = 42
    v4 := make(map[string]int)
    fmt.Println(v2, v3, v1)
    v5 := []int{int(v2)}
    var v6 float32 = (((-2.25 - (float32(float64((-2.25 - ((float32(0.0) + 1e10) * 0.0)))) * (float32(-2.25) - 3.14159))) - (((0.0 + 0.1) - (((((float32((1e10 + (((1.5 - ((0.0 - 0.0) - 0.1)) - 1.5) - 3.14159))) + 0.0) / ((float32(float64((-2.25 / (((0.1 + 1.5) + 3.14159) + float32(((((1e10 + ((((-2.25 + -2.25) * 1.5) + (((0.0 - 0.0) - ((-2.25 - float64((1.5 - 0.0))) / float64(3.14159))) * 3.14159)) + -2.25)) + (((float64(0.1) - float64(0.0)) / (3.14159 / 3.14159)) * ((float64(float32(((((-2.25 - (0.1 - (3.14159 * 0.1))) / (-2.25 / float64(((float32((1.5 / 1.0)) - (0.0 / ((0.0 - -2.25) + float32((((0.1 / (-2.25 + float64(0.0))) - 0.1) - 1e10))))) + (float32(-2.25) - (0.1 + -2.25)))))) * 1.5) * 0.1))) * float64(1e10)) / 1.0))) + 1.5) + 0.1)))))) + 1e10) * 0.1)) - -2.25) / 3.14159) * float32((((1e10 - -2.25) + float64(1e10)) / 1.0)))) * float32(-2.25))) + float32(-2.25))
    v3 = (v2 & v2)
    v6 = v6
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            if len(v5) > 0 {
                v5[0] = fi9
            }
            v6 += ((3.14159 / (v6 / (v6 / v6))) + ((v6 + (((v6 - v6) + (3.14159 + float32(1e10))) / (v6 / v6))) - (v6 - v6)))
            _ = fi9
        }
        var v10 uint8 = uint8(v7)
        switch fi8 {
        case -2147483648:
            _ = len(v5)
        case 998:
            if (false && true) {
                var v11 uint64 = (v7 * v7)
                fmt.Println(v7)
                if (v10 <= v10) {
                    var v12 string = ("a" + "")
                    _ = v12
                } else {
                    v13 := GenS1{F0: "hello", F1: -762}
                    _ = len(v5)
                    if false {
                        v2 -= v2
                        switch fi8 {
                        case 2147483647:
                            v7 += (v1 & 4294967295)
                        case 32768:
                            var v14 bool = (418 >= uint16(v1))
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 19; fi15++ {
                            delete(v4, "a")
                            var v16 uint64 = v7
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = v13
                }
                _ = v11
            } else {
                var v17 int8 = -128
                fmt.Println(v10, v1, v7)
                _ = v17
            }
        default:
            fmt.Println(v1, v7)
        }
        _ = fi8
        _ = v10
    }
    if (42 >= 363) {
        var v18 string = "\n\t"
        _ = v18
    } else {
        var v19 uint8 = 2
        fmt.Println(v1)
        _ = v19
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (0.1 <= (v6 / v6)) {
        fmt.Println(v2)
        v20 := GenS2{F0: 227}
        var v21 int = 127
        _ = v20
        _ = v21
    } else {
        if false {
            v4[("a" + (("" + ("a" + "a")) + ""))] = -464
            v22 := GenS3{F0: 42, F1: 567, F2: 2147483647}
            _ = v22
        }
        _ = len(v5)
        if false {
            var v23 int = 9223372036854775807
            var v24 uint32 = 0
            _ = v23
            _ = v24
        } else {
            v25 := GenS1{F0: "hello", F1: -528}
            if (42 > 2) {
                if true {
                    v26 := GenS2{F0: 1}
                    _ = v26
                } else {
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                }
                switch "hello" {
                case "hello":
                    _ = len(v4)
                case "a":
                    v5 = append(v5, 127)
                case "foo bar":
                    v6 += (((((v6 - v6) * 1.5) - ((v6 / (float32((float64(3.14159) * ((3.14159 - 0.0) * float64(((0.1 / v6) / v6))))) - (v6 - (v6 + (0.0 + 3.14159))))) * 1.5)) / v6) - v6)
                }
                var v27 uint64 = v7
                _ = v27
            }
            _ = v25
        }
    }
    v2 += 147
    v2 -= v3
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
