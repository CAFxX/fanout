package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 int = -1
    v2 := make(map[string]uint32)
    var v3 float64 = ((1.5 - ((-2.25 - (3.14159 / 1.0)) - 1e10)) / float64((float32(1e10) - 0.0)))
    var v4 int16 = 820
    var v5 float32 = (3.14159 - 1e10)
    var v6 int16 = (v4 + v4)
    var v7 int8 = int8(v4)
    fmt.Println(v7)
    fmt.Println(v7, v5)
    if true {
        v8 := GenS1{F0: 996}
        v1 = 42
        _ = v8
    } else {
        v1 = (v1 & v1)
        v9 := GenS2{F0: 89, F1: 0}
        var v10 uint8 = 1
        _ = v10
        _ = v9
    }
    switch ("foo bar" + ("\n\t" + "\n\t")) {
    case "":
        switch (v1 - (v1 % (((v1 * v1)) | 1))) {
        case 9223372036854775807:
            if (476 <= -1) {
                fmt.Println(v7, v1, v3)
            } else {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int64 = 1
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v6, v6)
                if true {
                    if (v7 <= v7) {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            v14 := GenS2{F0: 0, F1: 133}
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        var v15 int16 = (v4 ^ 485)
                        _ = v15
                    }
                }
            }
        default:
            fmt.Println(v7, v1, v1)
        }
    default:
        v4 = (^(v4))
    }
    v16 := GenS1{F0: 0}
    v17 := []int{-2147483648, v1, v1}
    if (1 > -726) {
        if (-814 < 0) {
            fmt.Println(v7)
        }
        var v18 int = (v1 & v1)
        v18 += 9223372036854775807
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
