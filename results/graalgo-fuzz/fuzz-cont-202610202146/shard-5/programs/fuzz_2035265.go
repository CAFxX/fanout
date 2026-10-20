package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int64; F2 int64 }

func main() {
    var v1 int = 2147483647
    var v2 float32 = (float32(float64(0.1)) / ((1e10 + 0.1) / 1.0))
    var v3 float32 = 0.1
    if ((3.14159 * (-2.25 / ((((0.0 + 1e10) / ((((3.14159 - (-2.25 - 3.14159)) / (((1.5 + (3.14159 - 0.0)) / 1e10) + ((3.14159 - -2.25) + (3.14159 / 1e10)))) / 3.14159) / 3.14159)) / (0.1 - -2.25)) + (((3.14159 + float64(1.5)) - 1e10) - 0.1)))) >= 0.1) {
        for fi4 := 0; fi4 < 9; fi4++ {
            fmt.Println(v2)
            v5 := []int{(v1 % ((128) | 1))}
            _ = fi4
            _ = v5
        }
        var v6 int32 = 42
        _ = v6
    }
    fmt.Println(v3)
    if (-1 < 830) {
        v3 = v3
        switch v1 {
        case 1:
            if (v1 < -9223372036854775808) {
                switch "\n\t" {
                case "hello":
                    v7 := GenS1{F0: 1000, F1: -1, F2: ""}
                    _ = v7
                case "\n\t":
                    if false {
                        v1 = v1
                    }
                default:
                    v8 := GenS2{F0: "hello"}
                    _ = v8
                }
            } else {
                if false {
                    if false {
                        v9 := GenS2{F0: "a"}
                        v10 := GenS3{F0: -768, F1: 993, F2: 0}
                        _ = v10
                        _ = v9
                    }
                } else {
                    var v11 int8 = 1
                    var v12 int = v1
                    var v13 int16 = 1000
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v1 = v1
            }
        default:
            var v14 int64 = 0
            _ = v14
        }
    }
    if false {
        v15 := []int{v1}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        fmt.Println(v1, v1)
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    v16 -= (v16 + (-(v16)))
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 float64 = 1.5
        v1 -= (fi17 & 316)
        var v19 float64 = (v18 + 0.1)
        _ = fi17
        _ = v18
        _ = v19
    }
    switch ("" + "a") {
    case "foo bar":
        v20 := []int{v1, 544, v1, -637}
        _ = v20
    case "":
        var v21 bool = true
        _ = v21
    default:
        var v22 int16 = 215
        _ = v22
    }
    var v23 uint8 = 128
    if (255 >= 18446744073709551615) {
        for fi24 := 0; fi24 < 12; fi24++ {
            v25 := []int{(fi24 >> ((272) & 63))}
            v16 += v16
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        fmt.Println(v1, v2, v23)
        v27 := []int{fi26, (-(fi26))}
        v16 -= (v16 / ((42) | 1))
        _ = fi26
        _ = v27
    }
    v16 -= (v16 >> ((2) & 63))
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
