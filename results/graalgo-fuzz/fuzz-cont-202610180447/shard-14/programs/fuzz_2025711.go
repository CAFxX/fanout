package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int64; F1 string }

func main() {
    var v1 int32 = -124
    var v2 uint16 = 0
    var v3 float64 = 0.0
    var v4 uint16 = (v2 ^ 330)
    var v5 int64 = 42
    for fi6 := 0; fi6 < 20; fi6++ {
        if true {
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(v3)
                var v8 float64 = 1.5
                _ = fi7
                _ = v8
            }
        }
        _ = fi6
    }
    if true {
        v9 := []int{128}
        _ = v9
    } else {
        v10 := GenS1{F0: 32767}
        v11 := []int{32768, -422}
        v12 := GenS2{F0: -43, F1: 128}
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: 9223372036854775807, F1: ("a" + "foo bar")}
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int16 = -655
        v16 := GenS1{F0: v15}
        v4 = (v4 ^ 32767)
        _ = fi14
        _ = v15
        _ = v16
    }
    v4 = (v2 / ((v2) | 1))
    v2 += (v4 & 32767)
    v5 += 545
    v17 := GenS2{F0: 127, F1: -490}
    var v18 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
