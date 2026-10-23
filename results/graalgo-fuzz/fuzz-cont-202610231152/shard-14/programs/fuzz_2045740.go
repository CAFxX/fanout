package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16; F1 int8; F2 bool }

func main() {
    v1 := []uint32{65535, 1, 633, 2147483647}
    var v2 float32 = (3.14159 + 0.0)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    fmt.Println(v2)
    _ = len(v1)
    fmt.Println(v2)
    switch 1 {
    case 42:
        v5 := GenS1{F0: 4294967295, F1: 42}
        _ = v5
    case -600:
        v6 := GenS2{F0: 1, F1: -100}
        _ = v6
    default:
        v7 := GenS2{F0: 257, F1: -128}
        _ = v7
    }
    if (!((588 == 32767) && false)) {
        var v8 uint8 = 42
        v2 -= float32(float64(0.0))
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        switch fi9 {
        case -213:
            v10 := GenS3{F0: 1000, F1: 0, F2: (309 >= 279)}
            _ = v10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(fi11, v2)
        _ = fi11
    }
    _ = len(v4)
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(fi12, v2)
        var v13 float64 = 3.14159
        var v14 uint16 = 833
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint64 = 18446744073709551615
    switch ("\n\t" + "a") {
    case "\n\t":
        v16 := GenS1{F0: 4294967295, F1: 223}
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(v2, v2, v15)
            _ = fi17
        }
    }
    if true {
        v18 := GenS3{F0: 1000, F1: 1, F2: true}
        var v19 uint32 = 243
        _ = v18
        _ = v19
    }
    _ = len(v1)
    v20 := GenS1{F0: (v15 % (((v15 >> ((65535) & 63))) | 1)), F1: 23}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
