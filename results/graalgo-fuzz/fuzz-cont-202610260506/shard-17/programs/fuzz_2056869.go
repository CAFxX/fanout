package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []int{128, 2147483647}
    var v2 float32 = 3.14159
    var v3 float64 = 3.14159
    var v4 uint32 = 300
    var v5 uint32 = (v4 | v4)
    var v6 float32 = (float32((float64(3.14159) - v3)) - float32((v3 * 1e10)))
    v4 += v5
    switch 32768 {
    case 32767:
        fmt.Println(v4)
    case 327:
        _ = len(v1)
    default:
        switch -333 {
        case 4294967295:
            v5 -= 0
        default:
            if (-9223372036854775808 > -710) {
                if false {
                    v5 += v5
                }
            }
        }
    }
    v4 -= v5
    v2 = float32(((float64(v2) * -2.25) + (((float64(v6) * float64(v2)) * v3) / (v3 / float64(1e10)))))
    if true {
        if (true && ((float64(((v2 - v2) + float32(0.0))) - (1e10 * v3)) != (v3 * float64(v6)))) {
            v5 -= v5
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 0
        }
    }
    var v7 int8 = -128
    v5 = 42
    var v8 int8 = 79
    for fi9 := 0; fi9 < 14; fi9++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi9
    }
    fmt.Println(v2, v7)
    v10 := GenS1{F0: 2, F1: 124}
    switch "" {
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = 127
        }
    case "hello":
        v11 := GenS1{F0: 255, F1: 117}
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 9; fi12++ {
            v7 -= v8
            var v13 string = "hello"
            _ = fi12
            _ = v13
        }
    default:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
