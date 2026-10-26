package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 uint8; F1 int; F2 bool }
type GenS5 struct { F0 int16 }

func main() {
    var v1 float32 = (0.0 + -2.25)
    v2 := make(map[string]uint32)
    var v3 uint8 = 23
    v4 := make(map[string]int)
    var v5 uint16 = 2
    fmt.Println(v3, v5)
    var v6 bool = (-161 <= 520)
    v2["foo bar"] = 384
    v6 = (325 == 66)
    var v7 uint8 = (v3 ^ 2)
    var v8 bool = true
    var v9 bool = true
    if v6 {
        switch "\n\t" {
        case "foo bar":
            switch ("a" + "\n\t") {
            case "a":
                fmt.Println(v1, v9, v1)
            case "\n\t":
                if (v1 == (((((v1 * 0.0) / (((v1 * -2.25) - 0.0) / 1e10)) * (3.14159 - float32(1.5))) / 1.0) + ((v1 / 1e10) / v1))) {
                    fmt.Println(v5)
                    v10 := GenS1{F0: 492}
                    v11 := GenS2{F0: 64}
                    _ = v10
                    _ = v11
                }
            default:
                var v12 uint32 = 2147483647
                _ = v12
            }
        default:
            if true {
                switch -9223372036854775808 {
                case 32767:
                    v13 := GenS3{F0: 582, F1: 0, F2: "foo bar"}
                    _ = v13
                case 905:
                    fmt.Println(v8)
                case -9223372036854775808:
                    delete(v2, "\n\t")
                }
            } else {
                var v14 uint8 = v7
                if v9 {
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v16 := GenS4{F0: v3, F1: fi15, F2: v8}
                        fmt.Println(v5)
                        _ = fi15
                        _ = v16
                    }
                }
                _ = v14
            }
        }
        if false {
            fmt.Println(v1)
        }
    } else {
        var v17 uint8 = (^(v7))
        fmt.Println(v3, v5, v3)
        var v18 float64 = (((1e10 / (-2.25 + 0.0)) + ((0.0 + 1.5) + (((1e10 / (0.1 - (-2.25 + 0.0))) + 3.14159) / 1.5))) * -2.25)
        _ = v17
        _ = v18
    }
    v19 := GenS5{F0: -92}
    var v20 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
