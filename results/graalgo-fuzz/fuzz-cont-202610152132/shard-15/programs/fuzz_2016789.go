package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int8 }
func genFn0() bool {
    var v5 float64 = (3.14159 / (float64(0.0) - 0.1))
    v6 := GenS1{F0: 128, F1: -2147483648}
    v5 += v5
    v5 -= v5
    _ = v5
    _ = v6
    return (1 != 1)
}

func main() {
    v1 := []string{"\n\t", "foo bar", ""}
    var v2 float32 = 1.5
    var v3 uint8 = 1
    v4 := []uint32{931, 0}
    var v7 uint16 = 65535
    v2 -= 3.14159
    v7 = (v7 >> ((255) & 63))
    v8 := GenS2{F0: -128}
    fmt.Println(v3)
    v2 += (float32(1e10) / v2)
    if true {
        for fi9 := 0; fi9 < 12; fi9++ {
            if len(v4) > 0 {
                v4[0] = 42
            }
            v7 = v7
            _ = fi9
        }
    }
    v10 := GenS2{F0: 70}
    switch -886 {
    case 1:
        if true {
            v7 -= 0
        } else {
            v4 = append(v4, 2147483648)
            if len(v1) > 0 {
                v1[0] = ""
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
