package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint64 = 2
    var v2 uint32 = 65535
    var v3 int8 = 127
    var v4 float64 = 0.0
    v1 += v1
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint64 = v1
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: ((((!true) || true) || (-32768 >= 32767)) || true), F1: -924}
    v8 := []int{128, -215, 175}
    var v9 int64 = -235
    fmt.Println(v9, v2)
    v10 := GenS2{F0: 161, F1: (v1 - v1), F2: -236}
    switch 50 {
    case 184:
        v11 := GenS3{F0: 963, F1: -9223372036854775808}
        _ = v11
    case 9223372036854775807:
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS2{F0: 2, F1: v1, F2: 42}
            _ = fi12
            _ = v13
        }
    case -192:
        fmt.Println(v1, v4)
    }
    if ("foo bar" >= (("" + ("a" + ("\n\t" + ("hello" + "")))) + "foo bar")) {
        fmt.Println(v2)
        v4 += ((1.5 / v4) * ((float64(float32((((((v4 + ((float64(0.0) * 0.0) * ((v4 / v4) + v4))) + -2.25) / ((v4 / (0.1 * float64(1.5))) * (v4 / v4))) * float64(0.0)) - (float64(3.14159) / 1.5)))) * v4) - 0.0))
        var v14 int64 = v9
        _ = v14
    } else {
        var v15 bool = (true || false)
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS4{F0: 1, F1: v9}
        v1 = (^(v1))
        v3 = 100
        _ = fi16
        _ = v17
    }
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
