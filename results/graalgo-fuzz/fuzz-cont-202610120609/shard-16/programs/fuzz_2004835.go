package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 string; F2 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint16 = 65535
    var v2 int64 = 9223372036854775807
    var v3 uint8 = 2
    var v4 float64 = 1.5
    var v5 uint = 42
    var v6 int32 = 481
    v7 := GenS1{F0: 313}
    var v8 float32 = ((1e10 + float32(v4)) + -2.25)
    v6 = int32(v5)
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int16 = 32767
        fmt.Println(v4, v5, v1)
        for fi11 := 0; fi11 < 4; fi11++ {
            v8 = ((v8 / (v8 - (((((v8 + v8) * v8) + ((((1.5 - (v8 / v8)) + float32((v4 / ((0.0 + (v4 + float64((v8 + v8)))) + (3.14159 * v4))))) / 3.14159) - v8)) / -2.25) / v8))) / v8)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    switch ("hello" + "\n\t") {
    case "":
        var v12 uint32 = 391
        _ = v12
    case "\n\t":
        fmt.Println(v3, v5)
    default:
        v13 := GenS2{F0: 765, F1: "\n\t", F2: v5}
        _ = v13
    }
    var v14 int64 = (v2 ^ 1)
    fmt.Println(v6)
    v2 = v2
    v14 = v2
    v3 = v3
    v1 -= 0
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := []int{(fi15 * (fi15 % ((fi16) | 1))), (fi15 << ((v5) & 63))}
            _ = len(v17)
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            v18 := GenS3{F0: (^(v6))}
            v14 += v14
            v14 += -1
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
