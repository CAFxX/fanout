package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0(p0 string) int16 {
    var v5 float64 = 1.5
    p0 = p0
    fmt.Println(v5, v5)
    v6 := GenS1{F0: 1000, F1: -17, F2: 1}
    _ = p0
    _ = v5
    _ = v6
    return 353
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 0
    var v3 float64 = (float64(((-2.25 + 3.14159) - (0.1 + 0.0))) + (-2.25 * 1.5))
    var v4 int32 = 42
    fmt.Println(v2, v2, v3)
    fmt.Println(v3, v3)
    v7 := GenS1{F0: (-(v2)), F1: -118, F2: -61}
    var v8 uint8 = 127
    var v9 uint = 458
    v8 -= v8
    v8 = 128
    var v10 uint16 = 65535
    v11 := GenS1{F0: (v2 - (v2 / ((v2) | 1))), F1: 1, F2: 42}
    for fi12 := 0; fi12 < 6; fi12++ {
        delete(v1, "\n\t")
        v1["foo bar"] = false
        switch "" {
        case "foo bar":
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS2{F0: v9, F1: (fi12 % (((^(fi12))) | 1))}
                _ = fi13
                _ = v14
            }
        case "a":
            v4 += (v4 % (((v4 ^ int32(v8))) | 1))
        }
        _ = fi12
    }
    switch 701 {
    case -1:
        v15 := GenS1{F0: (v2 + (^(v2))), F1: 0, F2: 0}
        _ = v15
    case 4294967295:
        v4 = 42
    }
    v4 = -959
    switch "a" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            delete(v1, (("\n\t" + ("" + (("\n\t" + "hello") + ("a" + "")))) + ("" + "foo bar")))
        case "foo bar":
            fmt.Println(v4, v4, v8)
        }
    case "hello":
        v16 := []int{int(v2), 1}
        _ = v16
    case "a":
        v1["hello"] = true
    }
    fmt.Println(v3)
    v17 := GenS1{F0: (v2 & 1), F1: -128, F2: 42}
    fmt.Println(v2, v3, v8)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 float32 = float32(float64(3.14159))
        v20 := GenS3{F0: 363, F1: 1}
        if (fi18 != (fi18 * fi18)) {
            delete(v1, "\n\t")
            v9 = v9
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{-2147483648, 9223372036854775807, -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
