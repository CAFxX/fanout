package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 float32 = 1e10
    var v2 float32 = v1
    var v3 int32 = 2147483647
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: 65535, F1: false, F2: 55}
        _ = fi4
        _ = v5
    }
    v3 -= (v3 - -1)
    var v6 string = "hello"
    fmt.Println(v1, v6)
    v7 := []int{}
    fmt.Println(v2, v1, v1)
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 string = v6
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(v2)
            if (fi8 > (^(fi10))) {
                var v11 uint8 = 170
                _ = v11
            } else {
                v1 -= (v2 - (((-2.25 + (v1 + (0.0 / (float32((float64(v2) - 0.1)) * 1e10)))) + float32(float64((v2 + ((v1 * v2) / v1))))) * -2.25))
                v12 := GenS1{F0: 159, F1: (v3 > v3), F2: 2}
                v13 := GenS2{F0: fi10}
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    switch v6 {
    case "hello":
        fmt.Println(v1, v6, v6)
    case "":
        for fi14 := 0; fi14 < 19; fi14++ {
            v2 += v1
            _ = fi14
        }
    case "foo bar":
        v1 -= (1.5 * -2.25)
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            v1 = (v2 * v2)
            v1 = ((1.5 * v1) - v2)
            v3 -= -1
            _ = fi16
        }
        var v17 bool = (!(100 > 69))
        _ = fi15
        _ = v17
    }
    var v18 int16 = 1000
    switch "a" {
    case "":
        var v19 int = 32768
        _ = v19
    default:
        v20 := GenS2{F0: 129}
        _ = v20
    }
    var v21 uint16 = 255
    fmt.Println(v1, v21, v6)
    fmt.Println(v1, v2, v2)
    fmt.Println(v21)
    v7 = append(v7, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
}
