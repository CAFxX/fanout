package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint16 = 78
    v3 := []uint32{v1, v1, v1}
    var v4 float32 = 1.5
    if true {
        v4 -= float32(0.1)
        var v5 uint32 = 259
        _ = v5
    }
    v2 += v2
    var v6 float64 = (1.5 - 3.14159)
    var v7 int16 = 457
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint16 = (v2 * (^(v2)))
        _ = fi8
        _ = v9
    }
    v4 += (v4 + (v4 * 1e10))
    fmt.Println(v4, v6)
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int16 = -32768
        v3 = append(v3, (v1 << ((1) & 63)))
        switch ("hello" + "a") {
        case "foo bar":
            var v12 int64 = 9223372036854775807
            _ = v12
        case "\n\t":
            v7 += v7
        }
        _ = fi10
        _ = v11
    }
    switch ("\n\t" + "hello") {
    case "hello":
        v7 += (-(v7))
    default:
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    if (!true) {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: (fi14 % ((fi14) | 1)), F1: 1, F2: int8(v7)}
            var v16 int8 = 100
            _ = fi14
            _ = v15
            _ = v16
        }
        if ((v4 < float32(v6)) || true) {
            for fi17 := 0; fi17 < 13; fi17++ {
                fmt.Println(v6, v1)
                fmt.Println(fi17, fi17, fi17)
                v1 = v1
                _ = fi17
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 float64 = v6
                _ = fi18
                _ = v19
            }
        }
    }
    v3 = append(v3, v1)
    var v20 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
