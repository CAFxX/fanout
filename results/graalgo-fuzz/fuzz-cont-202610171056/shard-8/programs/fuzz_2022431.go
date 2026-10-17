package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int32 = 61
    v2 := []string{}
    var v3 uint32 = 872
    var v4 int16 = 821
    var v5 int32 = (v1 >> ((18446744073709551615) & 63))
    fmt.Println(v5, v4)
    var v6 float32 = (float32(float64(0.0)) / ((1e10 * 1e10) * ((-2.25 / (-2.25 + (-2.25 * (0.0 / 1e10)))) - 1.5)))
    var v7 float64 = 3.14159
    v6 -= float32(float64(v6))
    v8 := GenS1{F0: v4, F1: 127}
    var v9 string = (("" + ("hello" + "foo bar")) + "\n\t")
    fmt.Println(v5)
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 bool = false
            v13 := GenS1{F0: (-(v4)), F1: int8(v4)}
            fmt.Println(v9, v4, fi10)
            _ = fi11
            _ = v12
            _ = v13
        }
        switch fi10 {
        case 9223372036854775807:
            var v14 uint8 = 196
            _ = v14
        case 42:
            v15 := GenS2{F0: 0}
            _ = v15
        default:
            if len(v2) > 0 {
                v2[0] = ("\n\t" + (v9 + "a"))
            }
        }
        _ = fi10
    }
    v4 = -32768
    v5 = 0
    var v16 uint = 224
    var v17 float64 = float64(1.5)
    var v18 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
