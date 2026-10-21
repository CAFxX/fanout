package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    var v1 string = "\n\t"
    var v2 int64 = 1
    var v3 uint8 = 192
    var v4 float32 = float32(-2.25)
    v5 := make(map[int]string)
    v6 := []int{1, 42, 42, 42}
    var v7 int8 = 100
    var v8 int64 = (^(v2))
    var v9 uint8 = 127
    v10 := GenS1{F0: 2147483647, F1: ("foo bar" + v1), F2: 1}
    if false {
        var v11 float32 = v4
        var v12 uint = 2
        _ = v11
        _ = v12
    }
    v6 = append(v6, 1)
    var v13 int32 = 751
    if false {
        v14 := GenS2{F0: 2, F1: v7}
        _ = v14
    } else {
        switch (v1 + "hello") {
        case "":
            v15 := GenS1{F0: 2147483647, F1: "a", F2: v9}
            _ = v15
        case "\n\t":
            v16 := GenS2{F0: (v9 / ((v3) | 1)), F1: (v7 / (((v7 % ((v7) | 1))) | 1))}
            _ = v16
        case "hello":
            if (v1 >= v1) {
                var v17 float32 = -2.25
                _ = v17
            } else {
                if true {
                    v18 := GenS1{F0: 365, F1: v1, F2: v3}
                    _ = len(v5)
                    var v19 uint16 = 32767
                    _ = v18
                    _ = v19
                } else {
                    v8 = (-(v2))
                    var v20 float32 = v4
                    _ = v20
                }
                v6 = append(v6, -680)
            }
        default:
            switch 127 {
            case -305:
                var v21 float32 = (v4 * ((((((float32(float64(v4)) / 1.0) - float32(float64(3.14159))) + v4) - (0.0 + 0.0)) - float32(1.5)) * (v4 + 3.14159)))
                _ = v21
            default:
                switch 32767 {
                case 2147483647:
                    if (2 <= 18446744073709551615) {
                        v8 = -9223372036854775808
                        if true {
                            v22 := GenS2{F0: (^(v9)), F1: v7}
                            if false {
                                v23 := GenS2{F0: 2, F1: v7}
                                _ = v23
                            }
                            _ = v22
                        }
                    }
                }
            }
        }
        v6 = append(v6, 128)
        var v24 int64 = (-(v8))
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
