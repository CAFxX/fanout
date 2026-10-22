package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0() uint {
    var v4 int32 = 42
    var v5 int = 4294967295
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 float64 = -2.25
        v8 := GenS1{F0: (-1 < -1), F1: "foo bar"}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v4
    _ = v5
    return 269
}
func genFn1() uint32 {
    var v9 int64 = 1
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int16 = -858
        var v12 float32 = (1e10 * 3.14159)
        switch ("" + "hello") {
        case "":
            var v13 int64 = (v9 * v9)
            _ = v13
        case "\n\t":
            var v14 uint16 = 1
            _ = v14
        case "foo bar":
            fmt.Println(fi10, v9)
        default:
            var v15 int8 = 127
            _ = v15
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    var v16 int8 = 0
    _ = v16
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v17 int = 32767
    fmt.Println(v17)
    switch "" {
    case "foo bar":
        v17 += (^(v17))
    default:
        var v18 uint32 = 0
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        switch "" {
        case "":
            fmt.Println(fi19, fi19)
        case "hello":
            switch "\n\t" {
            case "hello":
                var v20 int32 = 15
                _ = v20
            case "a":
                switch fi19 {
                case -2147483648:
                    var v21 float32 = (float32((0.0 * 1.5)) - 1.5)
                    _ = v21
                default:
                    var v22 uint16 = 1
                    _ = v22
                }
            case "\n\t":
                for fi23 := 0; fi23 < 20; fi23++ {
                    if (float64(float32(float64(1e10))) <= (-2.25 - float64((((1.5 / (1.5 + (1e10 / float32(((0.0 + (-2.25 + (float64(float32(3.14159)) / (0.1 * (0.0 - (-2.25 + 1.5)))))) * (0.0 - -2.25)))))) / 1.0) - 0.1)))) {
                        v17 += (fi23 / ((178) | 1))
                    }
                    _ = fi23
                }
            }
        }
        _ = fi19
    }
    var v24 uint16 = 32767
    v25 := GenS1{F0: false, F1: ""}
    for fi26 := 0; fi26 < 11; fi26++ {
        v17 += (v17 - 32768)
        _ = fi26
    }
    v24 = 42
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 bool = true
        _ = fi27
        _ = v28
    }
    v17 -= v17
    var v29 int = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
}
