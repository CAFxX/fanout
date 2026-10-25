package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0() uint8 {
    var v5 int = 75
    v6 := []int{-970}
    v5 -= (^(v5))
    var v7 uint8 = 127
    _ = v5
    _ = v6
    _ = v7
    return (^(v7))
}

func main() {
    var v1 uint32 = 734
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := []uint32{719, (v1 >> ((255) & 63)), v1, (v1 & 2147483648)}
    switch "a" {
    case "":
        fmt.Println(v1)
    case "\n\t":
        for fi8 := 0; fi8 < 12; fi8++ {
            v1 = 233
            v9 := GenS1{F0: 1, F1: (v1 & 288), F2: uint16(v1)}
            var v10 int8 = -128
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        var v11 uint8 = 137
        _ = v11
    }
    var v12 float64 = -2.25
    switch 127 {
    case 42:
        fmt.Println(v12, v1)
    case 92:
        switch 1 {
        case -9223372036854775808:
            v12 += v12
        case 4294967295:
            var v13 uint32 = 0
            _ = v13
        case 2147483647:
            v1 = (v1 & (v1 ^ v1))
        default:
            var v14 float32 = float32(v12)
            _ = v14
        }
    default:
        switch "a" {
        case "":
            switch "hello" {
            case "hello":
                var v15 int16 = 759
                _ = v15
            case "foo bar":
                var v16 uint8 = 128
                _ = v16
            case "\n\t":
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 bool = (float64(3.14159) >= (((float64(1.5) - v12) - v12) * v12))
                    _ = fi17
                    _ = v18
                }
            }
        default:
            if (false || true) {
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 int8 = -74
                    v1 += 791
                    fmt.Println(v20, v1)
                    _ = fi19
                    _ = v20
                }
                var v21 uint = 4294967295
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    var v22 uint8 = uint8(v21)
                    _ = v22
                case "hello":
                    for fi23 := 0; fi23 < 9; fi23++ {
                        var v24 int16 = 42
                        fmt.Println(fi23)
                        _ = fi23
                        _ = v24
                    }
                }
                _ = v21
            } else {
                v1 -= v1
                v1 -= v1
                var v25 uint8 = 0
                _ = v25
            }
        }
    }
    v12 = float64(0.0)
    for fi26 := 0; fi26 < 6; fi26++ {
        switch fi26 {
        case 42:
            v27 := v2[(fi26 << ((4294967295) & 63))]
            _ = v27
        case 714:
            var v28 float32 = -2.25
            _ = v28
        }
        v12 += v12
        _ = fi26
    }
    if (!true) {
        v3[9223372036854775807] = -301
        if (!true) {
            delete(v3, -991)
            v12 += (0.0 + v12)
        }
        v29 := GenS2{F0: 255, F1: 1}
        _ = v29
    } else {
        if len(v4) > 0 {
            v4[0] = 0
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            if (!(587 >= 2)) {
                v31 := GenS3{F0: 25, F1: true}
                if false {
                    var v32 uint8 = 42
                    _ = v32
                }
                _ = v31
            }
            var v33 uint32 = (v1 | v1)
            _ = fi30
            _ = v33
        }
        _ = len(v4)
    }
    v1 += v1
    v12 = 1.5
    v12 += (v12 + 0.0)
    var v34 float64 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
