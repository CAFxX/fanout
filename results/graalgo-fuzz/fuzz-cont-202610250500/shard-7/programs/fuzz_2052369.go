package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    var v3 int16 = 32767
    v4 := []uint32{543, 350, 2147483647, 65535}
    var v5 uint8 = 128
    var v6 uint16 = 65535
    var v7 uint64 = 854
    if false {
        switch "\n\t" {
        case "":
            v7 = (^(v7))
        case "foo bar":
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 float32 = -2.25
                _ = fi8
                _ = v9
            }
        case "hello":
            var v10 int16 = v3
            _ = v10
        default:
            v2 = v2
        }
    } else {
        var v11 float64 = 0.1
        v12 := GenS1{F0: (v5 + v5), F1: v2}
        v13 := GenS1{F0: v5, F1: (v2 ^ v2)}
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: (^(v5)), F1: v2}
    v2 = v2
    v7 += v7
    var v15 float32 = float32((1e10 / (((-2.25 - (0.1 / (float64(0.0) - -2.25))) / 1e10) * -2.25)))
    if (("" > "foo bar") || false) {
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                if (!(true || true)) {
                    if false {
                        v2 += (v2 % (((v2 / ((0) | 1))) | 1))
                    }
                }
                _ = fi17
            }
            v18 := GenS2{F0: ("\n\t" + ("" + "\n\t"))}
            _ = len(v1)
            _ = fi16
            _ = v18
        }
        v19 := GenS2{F0: "a"}
        _ = v19
    }
    v20 := GenS3{F0: 687}
    v21 := v1[32767]
    v22 := GenS4{F0: v3, F1: (v3 - 1000)}
    var v23 uint = 4294967295
    for fi24 := 0; fi24 < 2; fi24++ {
        switch "" {
        case "hello":
            if v21 {
                v25 := GenS4{F0: v3, F1: v3}
                var v26 int64 = 0
                _ = v25
                _ = v26
            }
        }
        _ = fi24
    }
    var v27 int64 = 9223372036854775807
    v15 = 3.14159
    var v28 bool = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if v28 {
        v21 = true
        v4 = append(v4, 4294967295)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
