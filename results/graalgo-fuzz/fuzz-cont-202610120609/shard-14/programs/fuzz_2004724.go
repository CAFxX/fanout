package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool }

func main() {
    v1 := []uint32{}
    v2 := []uint32{8, 2147483648, 875, 2147483648}
    var v3 int8 = -88
    v3 -= (v3 | v3)
    v4 := GenS1{F0: -1}
    var v5 float64 = float64(float32(-2.25))
    switch "a" {
    case "a":
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: -370}
            _ = fi6
            _ = v7
        }
    case "":
        v8 := GenS2{F0: 4294967295, F1: 42, F2: 206}
        _ = v8
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            switch "hello" {
            case "a":
                v10 := GenS1{F0: 32767}
                _ = v10
            case "hello":
                v3 += (v3 % ((-1) | 1))
            }
            v5 = ((v5 - v5) + (3.14159 * v5))
            _ = fi9
        }
    }
    fmt.Println(v5, v5)
    v11 := GenS3{F0: 255}
    _ = len(v1)
    switch "\n\t" {
    case "":
        switch -2147483648 {
        case -543:
            var v12 int32 = 2147483647
            _ = v12
        case 32767:
            var v13 int8 = v3
            _ = v13
        case 42:
            v14 := GenS4{F0: ((v5 - float64(1e10)) != (3.14159 + v5))}
            _ = v14
        }
    case "a":
        v3 = 100
    case "\n\t":
        v15 := GenS2{F0: 18446744073709551615, F1: 172, F2: 1}
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        switch "\n\t" {
        case "\n\t":
            var v17 int64 = -607
            _ = v17
        case "":
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS5{F0: 42, F1: false}
                fmt.Println(v3, v5)
                _ = fi18
                _ = v19
            }
        }
        _ = fi16
    }
    v2 = append(v2, 42)
    v3 -= -34
    if true {
        var v20 string = "hello"
        fmt.Println(v3, v3)
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
