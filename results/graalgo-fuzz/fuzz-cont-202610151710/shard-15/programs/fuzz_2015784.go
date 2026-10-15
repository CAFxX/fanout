package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int }

func main() {
    var v1 float64 = (float64(((-2.25 * 0.0) * 3.14159)) - 1e10)
    v2 := []string{("\n\t" + ""), "hello", "\n\t"}
    var v3 int16 = -1
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 float32 = ((1e10 / (-2.25 / 1.5)) * -2.25)
        var v6 float64 = v1
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += v3
    var v7 string = ""
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 127, F1: 2}
            _ = fi9
            _ = v10
        }
        fmt.Println(fi8, v1, v1)
        _ = fi8
    }
    if len(v2) > 0 {
        v2[0] = (v7 + v7)
    }
    v11 := GenS1{F0: -116, F1: 42}
    switch -9223372036854775808 {
    case 127:
        fmt.Println(v3, v1)
    case 42:
        v12 := GenS1{F0: 100, F1: 1}
        _ = v12
    case -878:
        v1 -= v1
    }
    v13 := GenS1{F0: 0, F1: 42}
    v14 := GenS2{F0: 127}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        _ = len(v2)
        var v16 uint32 = 1
        fmt.Println(v3)
        _ = fi15
        _ = v16
    }
    var v17 int = 4294967295
    var v18 int = v17
    var v19 int16 = -65
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
}
