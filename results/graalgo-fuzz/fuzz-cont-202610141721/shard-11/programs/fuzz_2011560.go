package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
func genFn0() uint64 {
    v5 := GenS1{F0: 494, F1: 43}
    v6 := GenS1{F0: 0, F1: 0}
    var v7 int32 = -2147483648
    var v8 int64 = 42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 297
    var v2 int = 1
    var v3 int = 4294967295
    var v4 int8 = 42
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "foo bar" {
        case "":
            switch v3 {
            case 649:
                fmt.Println(v2)
            default:
                fmt.Println(v2, v3, v1)
            }
        case "hello":
            var v10 int8 = v4
            _ = v10
        case "a":
            switch "hello" {
            case "\n\t":
                var v11 int32 = 42
                _ = v11
            case "hello":
                v12 := []int{(v2 | (-(fi9))), (v2 ^ fi9), 127, fi9}
                _ = v12
            case "":
                v13 := GenS2{F0: 861}
                _ = v13
            }
        }
        _ = fi9
    }
    v14 := GenS2{F0: genFn0()}
    v15 := []int{v2, (v3 / (((v3 & (v2 >> ((747) & 63)))) | 1)), 4294967295, v3}
    var v16 bool = (1 <= 0)
    v4 += v4
    switch "\n\t" {
    case "foo bar":
        switch "" {
        case "":
            fmt.Println(v2, v2, v2)
        case "foo bar":
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        default:
            for fi17 := 0; fi17 < 15; fi17++ {
                fmt.Println(v1, v3, v4)
                _ = fi17
            }
        }
    case "":
        var v18 uint16 = v1
        _ = v18
    }
    v4 = (v4 & (v4 << ((0) & 63)))
    for fi19 := 0; fi19 < 4; fi19++ {
        v16 = (!v16)
        _ = fi19
    }
    var v20 float64 = (0.0 * -2.25)
    v21 := GenS3{F0: int(v1)}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if v16 {
            if v16 {
                if v16 {
                    fmt.Println(v2, v16)
                    var v23 float32 = (((1.5 + ((-2.25 / float32(float64(float32(v20)))) / 1e10)) + 1e10) * 1.5)
                    _ = len(v15)
                    _ = v23
                } else {
                    v2 -= fi22
                }
            }
            if (((((0.0 * 1e10) / 1e10) * 0.1) + 1e10) <= 1e10) {
                v24 := GenS1{F0: 0, F1: 24}
                v4 += (v4 - (v4 + (v4 ^ 3)))
                var v25 float64 = (v20 / (((((((3.14159 * (v20 * v20)) * (v20 / ((0.0 / (((1.5 - v20) / v20) * (v20 * v20))) + float64(0.1)))) * v20) - (((float64(1e10) - float64((((float32(v20) * (1e10 + -2.25)) - float32(1.5)) + 0.0))) + (((0.0 * v20) * float64(0.1)) + v20)) * 1e10)) - v20) + v20) / float64(1.5)))
                _ = v24
                _ = v25
            } else {
                var v26 int32 = 0
                v4 -= v4
                _ = v26
            }
        } else {
            var v27 uint32 = 5
            fmt.Println(v20)
            _ = v27
        }
        _ = fi22
    }
    _ = len(v15)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
