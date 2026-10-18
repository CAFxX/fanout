package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int8 }

func main() {
    var v1 int = 32767
    var v2 uint64 = 982
    v3 := []int{(v1 ^ 42), v1}
    var v4 float32 = (((0.1 / ((float32((((1e10 * (1e10 * (1e10 / 1.0))) * ((1.5 - float64(float32(-2.25))) / 3.14159)) / 0.1)) - (-2.25 - (0.1 + 1.5))) / 1.0)) - (0.1 * 1e10)) / 0.1)
    v5 := []int{-2147483648, (v1 & (v1 / (((v1 / (((v1 + 32768)) | 1))) | 1))), (v1 - (v1 | v1))}
    var v6 int32 = -1
    fmt.Println(v6, v2)
    v5 = append(v5, (v1 << ((65535) & 63)))
    if (2 >= (v2 | v2)) {
        if ((v2 / ((v2) | 1)) <= (v2 >> ((255) & 63))) {
            v4 += v4
            var v7 int64 = -587
            _ = v7
        } else {
            v6 = v6
            if false {
                var v8 int64 = 570
                for fi9 := 0; fi9 < 14; fi9++ {
                    if true {
                        _ = len(v3)
                    }
                    _ = fi9
                }
                _ = v8
            } else {
                var v10 uint16 = 1
                var v11 int32 = -1
                v6 -= 519
                _ = v10
                _ = v11
            }
            fmt.Println(v4, v1, v2)
        }
        v12 := GenS1{F0: 128, F1: (true || (((v4 / (v4 + float32(float64(v4)))) + v4) >= v4)), F2: -128}
        _ = v12
    }
    var v13 string = "\n\t"
    if (!((42 < 0) && true)) {
        if (673 < 42) {
            if (0 > (v6 << ((763) & 63))) {
                switch 611 {
                case 32767:
                    v14 := GenS1{F0: 9, F1: (-32768 <= -1), F2: -128}
                    _ = v14
                case -2147483648:
                    if false {
                        var v15 uint64 = v2
                        v1 = -9223372036854775808
                        _ = v15
                    }
                case 127:
                    v16 := GenS2{F0: 422}
                    _ = v16
                default:
                    v4 -= v4
                }
                fmt.Println(v4)
                if true {
                    var v17 uint8 = 255
                    var v18 uint32 = 962
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint32 = 1
                    fmt.Println(v4, v1)
                    _ = v19
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch ((v13 + v13) + "hello") {
            case "\n\t":
                v2 += 708
            case "foo bar":
                v20 := GenS3{F0: v13, F1: 56}
                _ = v20
            case "hello":
                fmt.Println(v4)
            default:
                if true {
                    v4 = 1.5
                    v1 -= (v1 << ((0) & 63))
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v21 uint64 = 42
                    _ = v21
                }
            }
        }
        fmt.Println(v13)
    }
    v2 = (v2 | (v2 * 42))
    var v22 uint64 = (v2 + v2)
    v23 := GenS1{F0: 197, F1: (false || (0 != 1)), F2: 4}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
