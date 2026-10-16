package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int16; F1 uint16; F2 int8 }

func main() {
    var v1 float64 = (0.1 + (-2.25 + 0.1))
    v2 := make(map[string]string)
    var v3 bool = false
    v4 := make(map[uint32]string)
    v5 := []string{}
    v1 = (1e10 - (v1 + 0.1))
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = 255
        _ = len(v2)
        switch "foo bar" {
        case "hello":
            for fi8 := 0; fi8 < 7; fi8++ {
                v1 = v1
                _ = fi8
            }
        }
        _ = fi6
        _ = v7
    }
    delete(v4, 340)
    fmt.Println(v1)
    v9 := GenS1{F0: 0}
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if v3 {
                v12 := GenS1{F0: -1}
                fmt.Println(v3, fi11, fi10)
                _ = v12
            } else {
                if (255 < 42) {
                    v13 := GenS2{F0: 10}
                    v14 := GenS3{F0: 18446744073709551615, F1: 215}
                    var v15 int32 = 0
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    v1 -= v1
                    v1 -= float64(0.0)
                    v4[42] = "\n\t"
                    var v16 int16 = -32768
                    v1 = ((0.0 + (1.5 - 1.5)) / 1.0)
                    _ = v16
                }
                switch ((((("a" + "a") + "a") + "foo bar") + "foo bar") + "a") {
                case "foo bar":
                    var v17 float64 = float64(0.1)
                    _ = v17
                case "hello":
                    var v18 float32 = -2.25
                    _ = v18
                case "":
                    v3 = true
                default:
                    v5 = append(v5, "hello")
                }
                switch fi11 {
                case 150:
                    v19 := GenS4{F0: 44, F1: 32767, F2: 127}
                    _ = v19
                case 32767:
                    v20 := v4[1]
                    _ = v20
                case -476:
                    var v21 int8 = -128
                    _ = v21
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    v1 = (v1 / (1.5 * (0.1 + v1)))
    v1 -= v1
    var v22 int8 = 1
    var v23 float64 = v1
    v24 := GenS4{F0: 1000, F1: 42, F2: (v22 << ((164) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
