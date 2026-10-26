package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 string }

func main() {
    var v1 float64 = ((0.1 + 0.1) - 3.14159)
    v2 := []uint32{}
    var v3 string = ((("foo bar" + "\n\t") + "\n\t") + "hello")
    var v4 string = v3
    if (!true) {
        v4 = v4
        var v5 float32 = ((float32((((float64(-2.25) * 0.0) - v1) * -2.25)) - 0.1) + 1.5)
        _ = v5
    }
    var v6 bool = (!true)
    fmt.Println(v1)
    v4 = ((v4 + (("" + v3) + (("foo bar" + ((v4 + "\n\t") + v3)) + "foo bar"))) + "")
    _ = len(v2)
    v6 = true
    if (0 >= 42) {
        v1 += ((-2.25 - v1) / v1)
        if false {
            v4 = ("foo bar" + v3)
            var v7 int32 = -164
            _ = v7
        } else {
            var v8 int = 32767
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint16 = 1
                var v12 float32 = -2.25
                var v13 int8 = 21
                if (!((fi10 * (fi10 % ((fi9) | 1))) != (fi10 / ((1) | 1)))) {
                    var v14 int32 = int32(v11)
                    _ = v14
                } else {
                    v13 -= (v13 & (^(v13)))
                    v15 := GenS1{F0: 1, F1: 1000, F2: -2147483648}
                    var v16 bool = v6
                    _ = v15
                    _ = v16
                }
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            var v17 float32 = float32(1.5)
            var v18 float64 = (-2.25 / (v1 / float64(v17)))
            _ = fi9
            _ = v17
            _ = v18
        }
    } else {
        fmt.Println(v3, v4)
        v19 := GenS1{F0: -1, F1: -94, F2: 1}
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 int16 = 42
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v23 := GenS2{F0: 0, F1: 83, F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
