package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0() uint64 {
    var v4 int32 = 2147483647
    fmt.Println(v4)
    var v5 string = "foo bar"
    _ = v4
    _ = v5
    return 531
}

func main() {
    var v1 float64 = ((3.14159 + 3.14159) + 0.0)
    v2 := make(map[uint32]string)
    v3 := []uint32{1}
    var v6 string = "\n\t"
    v7 := GenS1{F0: 229, F1: 148}
    switch 4294967295 {
    case -353:
        fmt.Println(v1)
    case -9223372036854775808:
        var v8 int = 80
        _ = v8
    default:
        v9 := GenS2{F0: 42, F1: 0, F2: 127}
        _ = v9
    }
    v6 = (v6 + v6)
    var v10 float32 = ((0.1 - (1e10 * ((-2.25 + 1.5) - ((float32((1e10 * ((0.1 * v1) / ((float64(1e10) - (v1 - float64((1e10 * float32(-2.25))))) * v1)))) * (0.0 + 0.1)) + 0.1)))) * 0.0)
    if (65535 >= 4294967295) {
        var v11 int = -859
        var v12 uint = 651
        switch (((((v6 + (v6 + (v6 + "\n\t"))) + ((("" + v6) + "hello") + "\n\t")) + ("" + "foo bar")) + "\n\t") + v6) {
        case "":
            v6 = v6
        case "a":
            v6 = ""
        case "\n\t":
            _ = len(v3)
        default:
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int8 = -128
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 int16 = 1
                    _ = fi15
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
        }
        _ = v11
        _ = v12
    }
    v17 := GenS3{F0: 100, F1: 431}
    var v18 bool = true
    v19 := GenS1{F0: 253, F1: 193}
    v10 = (v10 + float32(float64(((float32(v1) * (v10 - ((v10 * -2.25) / v10))) * v10))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
