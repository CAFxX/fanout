package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 string; F2 int16 }

func main() {
    var v1 int = 758
    var v2 bool = (65535 != 0)
    var v3 uint64 = 4294967295
    v4 := GenS1{F0: 0, F1: 1}
    v5 := []int{(v1 + 11), v1, v1, -1}
    v6 := GenS2{F0: 128, F1: 32767, F2: 2147483647}
    v3 -= (^(v3))
    switch "hello" {
    case "hello":
        var v7 int = v1
        _ = v7
    case "\n\t":
        v1 += 593
    case "a":
        if len(v5) > 0 {
            v5[0] = 15
        }
    default:
        var v8 int8 = -23
        _ = v8
    }
    fmt.Println(v3)
    var v9 bool = ((0.0 - 1.5) > (0.0 * 1e10))
    if v2 {
        switch v1 {
        case 2147483647:
            var v10 int8 = 33
            _ = v10
        case 9223372036854775807:
            v2 = ((!(v2 && (!(!((0.0 >= (-2.25 / 1.5)) || v9))))) || true)
        case -991:
            var v11 int64 = -1
            _ = v11
        }
        if (898 < 1) {
            var v12 uint = 4294967295
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        _ = fi13
    }
    var v14 string = "\n\t"
    if len(v5) > 0 {
        v5[0] = v1
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v15 := GenS3{F0: 793}
    v5 = append(v5, 239)
    v16 := GenS1{F0: 2, F1: 0}
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            v5 = append(v5, fi18)
            if true {
                v19 := GenS1{F0: 133, F1: 42}
                _ = v19
            } else {
                var v20 int32 = -644
                v5 = append(v5, (v1 / (((fi17 >> ((317) & 63))) | 1)))
                _ = v20
            }
            v21 := GenS4{F0: (!(42 > 2)), F1: (v14 + (v14 + "\n\t")), F2: -961}
            _ = fi18
            _ = v21
        }
        var v22 float32 = (float32((1e10 - 1.5)) - -2.25)
        fmt.Println(fi17, v3, v3)
        _ = fi17
        _ = v22
    }
    v9 = ((1e10 < ((0.1 + 0.1) - (3.14159 + 3.14159))) || (!true))
    v5 = append(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
