package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int = -1
    var v2 string = ("hello" + (((("" + ("a" + "\n\t")) + "foo bar") + "hello") + ((("foo bar" + "\n\t") + "\n\t") + (("" + "") + ("" + ("" + "foo bar"))))))
    var v3 float64 = (float64(float32(-2.25)) / (float64(1e10) / 1.0))
    var v4 uint32 = 2147483647
    v5 := GenS1{F0: 4294967295, F1: (v1 & (v1 / ((v1) | 1)))}
    v6 := []int{v1, v1, (v1 & (-(v1)))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 float64 = float64((0.0 / -2.25))
    for fi8 := 0; fi8 < 3; fi8++ {
        v4 -= 90
        _ = fi8
    }
    var v9 int8 = 46
    for fi10 := 0; fi10 < 7; fi10++ {
        if len(v6) > 0 {
            v6[0] = v1
        }
        _ = fi10
    }
    v9 = -1
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int16 = 42
        v6 = append(v6, v1)
        _ = fi11
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS2{F0: -1, F1: 32767}
        v7 -= 3.14159
        _ = fi14
        _ = v15
    }
    if len(v6) > 0 {
        v6[0] = (v1 - v1)
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        switch (^(fi16)) {
        case 347:
            v17 := GenS2{F0: -1, F1: 497}
            _ = v17
        case 32768:
            v1 -= int(v9)
        }
        var v18 uint16 = 980
        _ = fi16
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 float64 = (v3 * (float64(0.0) / v3))
        _ = fi19
        _ = v20
    }
    var v21 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
