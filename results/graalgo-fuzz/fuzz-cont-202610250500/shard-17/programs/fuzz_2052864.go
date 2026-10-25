package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 string = ("" + "a")
    v2 := make(map[string]bool)
    v3 := []uint32{}
    v4 := GenS1{F0: 0}
    _ = len(v2)
    v5 := GenS1{F0: 100}
    var v6 float64 = -2.25
    var v7 uint = 18446744073709551615
    v1 = v1
    v8 := GenS2{F0: 631, F1: 0}
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            v2[("" + "")] = (false || true)
            var v11 float64 = (float64((0.1 / 1.0)) * float64(1.5))
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS3{F0: ((!(-1 == 32767)) && true), F1: 1}
            v3 = append(v3, 1)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int32 = 42
            var v16 int64 = 9223372036854775807
            v16 -= v16
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
