package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }

func main() {
    var v1 uint8 = 2
    var v2 uint8 = 0
    var v3 bool = false
    var v4 int64 = 0
    fmt.Println(v2, v3)
    fmt.Println(v4, v3, v3)
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 float32 = ((float32((float64(1.5) - (((1.5 / ((-2.25 / 0.1) + ((1e10 + 0.1) - 1e10))) - 1.5) - -2.25))) / (float32(0.1) - ((0.0 + (1e10 + 1.5)) + float32(0.0)))) - -2.25)
        _ = fi5
        _ = v6
    }
    var v7 float32 = 3.14159
    v8 := []int{1}
    v2 -= v1
    if (!(v3 || v3)) {
        fmt.Println(v1, v1)
        fmt.Println(v4, v7, v4)
        v2 -= (v1 * 255)
    }
    switch "foo bar" {
    case "\n\t":
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS1{F0: 32767}
            _ = fi9
            _ = v10
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    switch 2147483647 {
    case 1:
        if v3 {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint64 = 4294967295
                _ = fi11
                _ = v12
            }
            v13 := GenS2{F0: 121, F1: 227, F2: (!(!v3))}
            var v14 int64 = 857
            _ = v13
            _ = v14
        } else {
            v15 := GenS2{F0: 880, F1: 0, F2: v3}
            _ = v15
        }
    case -282:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    case 2147483647:
        switch (("foo bar" + "foo bar") + (("\n\t" + "a") + ("\n\t" + (("\n\t" + ("\n\t" + ("" + ("\n\t" + "hello")))) + ("" + ""))))) {
        case "hello":
            var v16 float32 = -2.25
            _ = v16
        default:
            v2 += v1
        }
    default:
        if len(v8) > 0 {
            v8[0] = -2147483648
        }
    }
    var v17 int16 = 821
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
