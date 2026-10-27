package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32; F1 uint; F2 int32 }
func genFn0() int {
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 int16 = 1000
    v7 := GenS1{F0: 0, F1: 1, F2: 2}
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v6, fi8)
        _ = fi8
    }
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int8) string {
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int16 = -32768
            p0 -= (p0 - (p0 + p0))
            v12 := []int{32768, fi9, fi9, (fi9 ^ fi10)}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    var v13 uint16 = 2
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 uint16 = (v13 | 42)
            switch "" {
            case "a":
                fmt.Println(v13, p0)
            case "foo bar":
                v16 += (^(v16))
            case "\n\t":
                v13 = v16
            default:
                v17 := GenS1{F0: v16, F1: p0, F2: 18446744073709551615}
                _ = v17
            }
            var v18 uint64 = 18446744073709551615
            _ = fi15
            _ = v16
            _ = v18
        }
        fmt.Println(v13, p0)
        _ = fi14
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := []int{}
        if (-2.25 >= 1.5) {
            fmt.Println(v13, fi19, p0)
            p0 = p0
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = v13
    return ("\n\t" + "hello")
}

func main() {
    var v1 uint8 = 42
    var v2 bool = true
    var v3 float64 = -2.25
    var v4 float32 = 0.1
    v3 = v3
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4, v3, v1)
    case "a":
        fmt.Println(v3, v4)
    }
    v2 = (false || (1 <= 850))
    v21 := GenS2{F0: v1, F1: 65535, F2: 255}
    fmt.Println(v1, v3)
    if (v1 >= (^(v1))) {
        v22 := GenS3{F0: -899, F1: 1}
        _ = v22
    } else {
        v23 := []int{0, 9223372036854775807}
        v24 := GenS4{F0: 0, F1: 2, F2: -1}
        _ = v23
        _ = v24
    }
    if (!(v3 <= float64((-2.25 + ((v4 - ((float32(1e10) * v4) * (v4 * (v4 / ((v4 - float32(((((1e10 + 0.0) + 1.5) + v3) * float64(-2.25)))) * ((-2.25 * ((v4 - (-2.25 - v4)) - v4)) / 1.0)))))) + 0.1))))) {
        v4 += (v4 * 1.5)
        v4 = ((((v4 - -2.25) * v4) - (v4 + (v4 * ((v4 * (v4 + v4)) * float32(float64((v4 * 0.0))))))) - (float32((v3 - v3)) * float32(v3)))
        v25 := []int{}
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := []int{(fi26 << ((393) & 63))}
        _ = fi26
        _ = v27
    }
    v28 := []int{9223372036854775807, 128, 855}
    fmt.Println(v2, v3)
    var v29 int16 = 32767
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 uint64 = 216
        _ = fi30
        _ = v31
    }
    v32 := GenS1{F0: 42, F1: -69, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
