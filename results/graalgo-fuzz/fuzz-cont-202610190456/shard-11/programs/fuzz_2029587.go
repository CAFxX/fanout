package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int; F2 string }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    var v3 int64 = 399
    var v4 uint32 = 4294967295
    var v5 uint64 = 42
    v6 := make(map[uint32]uint32)
    v7 := []int{}
    for fi8 := 0; fi8 < 1; fi8++ {
        if false {
            v4 += v4
            for fi9 := 0; fi9 < 14; fi9++ {
                switch (fi9 ^ (-(fi8))) {
                case -375:
                    fmt.Println(v5, v5, v3)
                case 119:
                    _ = len(v2)
                case 32768:
                    var v10 int16 = int16(v5)
                    _ = v10
                default:
                    switch fi8 {
                    case -773:
                        v1 -= v1
                    }
                }
                _ = fi9
            }
            v5 += (v5 >> ((255) & 63))
        } else {
            var v11 uint32 = (v4 - v4)
            if len(v7) > 0 {
                v7[0] = fi8
            }
            _ = v11
        }
        _ = fi8
    }
    _ = len(v7)
    var v12 int8 = -45
    v13 := GenS1{F0: 925, F1: 1}
    if false {
        var v14 string = ("hello" + "foo bar")
        switch "a" {
        case "":
            v15 := GenS1{F0: 2, F1: 127}
            _ = v15
        }
        _ = v14
    }
    var v16 string = "a"
    v17 := GenS2{F0: 18446744073709551615, F1: 65535, F2: 468}
    if true {
        var v18 float32 = ((1.5 - 1e10) / 1.5)
        _ = v18
    } else {
        var v19 uint16 = 1
        for fi20 := 0; fi20 < 3; fi20++ {
            v12 = v12
            _ = fi20
        }
        _ = v19
    }
    fmt.Println(v5, v3)
    v21 := GenS3{F0: (!((0 > -1) || (!((float32((float64((0.0 - 1e10)) - 1e10)) <= 1e10) || (673 < -1)))))}
    v22 := GenS4{F0: v3, F1: -72, F2: v16}
    _ = len(v2)
    var v23 uint64 = (v5 << ((574) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
