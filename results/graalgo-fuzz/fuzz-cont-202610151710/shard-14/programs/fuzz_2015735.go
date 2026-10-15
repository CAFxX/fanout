package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 uint64 = 42
    var v2 float32 = (1.5 / 0.1)
    v3 := []string{(("foo bar" + "foo bar") + ("hello" + (("\n\t" + "") + "foo bar"))), (((("foo bar" + "foo bar") + "hello") + ("foo bar" + "hello")) + "hello")}
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: -128, F1: (((("hello" + "hello") + ((("\n\t" + "a") + "\n\t") + "")) + (("\n\t" + "") + ("foo bar" + ("\n\t" + "")))) + (("a" + "hello") + "hello"))}
    for fi6 := 0; fi6 < 11; fi6++ {
        v2 -= (float32(((-2.25 - (1e10 * (-2.25 / 3.14159))) - -2.25)) + (3.14159 - (float32((1.5 / ((1e10 + (float64(((((v2 - v2) / 1e10) * ((v2 + (((-2.25 + ((1.5 * 0.0) + float32((0.1 - (-2.25 / 0.1))))) * v2) + ((v2 - v2) / (0.1 + 3.14159)))) - v2)) - ((v2 - (v2 - v2)) * -2.25))) + 1.5)) - (1.5 / 1e10)))) + v2)))
        _ = fi6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int16 = 32767
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 128
    var v10 float32 = 0.0
    if (1 <= 533) {
        v4[1] = 2
        v10 = ((v10 / (float32(0.0) - 1.5)) * v2)
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: -123, F1: "foo bar"}
            _ = len(v3)
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v9 -= (v9 | v9)
        var v14 uint8 = (v9 << ((454) & 63))
        _ = len(v4)
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 127, F1: (((("hello" + ("foo bar" + (("foo bar" + "foo bar") + ("foo bar" + ((("\n\t" + ("\n\t" + "foo bar")) + "hello") + "foo bar"))))) + "\n\t") + "foo bar") + ("hello" + ""))}
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint = 0
        _ = fi16
        _ = v17
    }
    v10 += (v10 / v2)
    v2 -= 1.5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v1)
    fmt.Println(v9, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
