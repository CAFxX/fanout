package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int; F1 int; F2 int16 }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }
func genFn0() int {
    v4 := []int{32767, -2147483648, 4294967295, 2147483647}
    var v5 uint32 = 353
    switch "" {
    case "hello":
        var v6 uint64 = 42
        _ = v6
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 uint8 = 128
    v2 := []string{""}
    var v3 uint64 = 477
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 := GenS1{F0: uint8(v3), F1: 18446744073709551615, F2: 1}
    _ = len(v2)
    v8 := GenS2{F0: genFn0(), F1: 746, F2: 32767}
    v9 := GenS3{F0: false, F1: 42, F2: -1}
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint = 1
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v1 += v1
    if ((float32(float64(1.5)) + ((1e10 + -2.25) / 3.14159)) >= (float32(((-2.25 + float64((1.5 / 0.1))) - 0.1)) + 0.0)) {
        fmt.Println(v1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 += v3
    } else {
        switch 32768 {
        case -162:
            fmt.Println(v1, v1)
        default:
            var v13 int16 = -131
            _ = v13
        }
        v3 = v3
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v1 += (v1 - v1)
        v1 -= 128
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
