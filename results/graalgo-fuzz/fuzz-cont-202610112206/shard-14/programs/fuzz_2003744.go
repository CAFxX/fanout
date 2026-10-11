package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int16 }

func main() {
    var v1 int = -2147483648
    v2 := []int{2147483647, v1, -631}
    var v3 int = (v1 & v1)
    var v4 float64 = (float64((3.14159 + 3.14159)) / ((0.1 * (((-2.25 + 1e10) * 3.14159) - ((3.14159 - (((((0.0 * 3.14159) - ((1e10 * 1e10) / float64((1e10 + 0.0)))) - float64(1e10)) + 0.0) / 1.0)) * float64((float32(-2.25) - 0.0))))) - -2.25))
    var v5 uint = 42
    var v6 bool = true
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!((v1 ^ -9223372036854775808) > v3)) {
        v2 = append(v2, (-(v1)))
    } else {
        var v7 uint32 = 2
        v8 := GenS1{F0: (v3 & (^(v1))), F1: v1, F2: 907}
        _ = v7
        _ = v8
    }
    if (false && ("\n\t" == "foo bar")) {
        fmt.Println(v1, v3, v6)
        switch "\n\t" {
        case "\n\t":
            var v9 int16 = 1
            _ = v9
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = 840
            }
        }
    } else {
        if v6 {
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS2{F0: 344, F1: 1}
                _ = fi10
                _ = v11
            }
        } else {
            var v12 int64 = -9223372036854775808
            v2 = append(v2, 32767)
            _ = v12
        }
    }
    var v13 float32 = (3.14159 * float32(((v4 / (((((v4 / float64(1.5)) * 0.1) / v4) / v4) * v4)) - -2.25)))
    var v14 uint = (v5 - (v5 + 897))
    v15 := GenS3{F0: 1000}
    v4 += float64(v13)
    var v16 uint8 = 0
    switch ("a" + "a") {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        v17 := GenS4{F0: -128, F1: 42}
        _ = v17
    case "foo bar":
        fmt.Println(v1, v6, v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
