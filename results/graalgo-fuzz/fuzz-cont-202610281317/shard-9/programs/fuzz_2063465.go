package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 string }

func main() {
    var v1 float64 = 1.5
    var v2 float32 = 1e10
    v3 := []string{"\n\t", ""}
    var v4 float32 = -2.25
    var v5 int = -9223372036854775808
    v5 += -118
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v2)
        _ = fi6
    }
    var v7 int = 4294967295
    var v8 int32 = 49
    switch "\n\t" {
    case "\n\t":
        v8 = v8
    case "":
        if (!(!((v2 - ((float32(v1) / (v2 * float32(float64((v2 - ((((float32((-2.25 - v1)) + 1e10) + (v4 * float32(v1))) + v4) * 1.5)))))) / v4)) >= (((float32(-2.25) / (v2 + v4)) - (v2 + v4)) * 1e10)))) {
            fmt.Println(v8, v2, v7)
            if true {
                v5 = (v7 * 32768)
            }
        }
    default:
        if (-1 >= 9223372036854775807) {
            switch ("hello" + "foo bar") {
            case "a":
                var v9 int = v7
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                v1 = 1.5
                _ = fi10
            }
            v7 = (-(v7))
        } else {
            fmt.Println(v8)
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 uint8 = 2
                _ = fi11
                _ = v12
            }
        }
    }
    fmt.Println(v8, v1, v1)
    if false {
        fmt.Println(v2, v7)
    } else {
        v13 := GenS1{F0: 128, F1: (("" + "\n\t") + ""), F2: "foo bar"}
        var v14 bool = (1 <= 1)
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v2 -= 1.5
        _ = fi15
    }
    var v16 uint16 = 0
    var v17 uint32 = uint32(v16)
    fmt.Println(v17, v2, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
