package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint = 1
    var v2 uint8 = 255
    var v3 uint8 = uint8(v1)
    v2 = 42
    var v4 uint = (v1 >> (((v1 >> (((^(v1))) & 63))) & 63))
    var v5 float32 = ((float32(((((1e10 + 0.1) * 0.1) / 1.0) / 3.14159)) - 0.1) + float32((((1e10 / 1e10) - -2.25) / 1e10)))
    var v6 int16 = 0
    v7 := GenS1{F0: 255, F1: v3}
    v2 = 2
    v5 = v5
    fmt.Println(v3, v5)
    if (100 > -19) {
        switch 377 {
        case 32767:
            fmt.Println(v1)
        default:
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 float32 = float32(-2.25)
                _ = fi8
                _ = v9
            }
        }
    }
    v10 := GenS1{F0: 2, F1: v2}
    for fi11 := 0; fi11 < 6; fi11++ {
        switch (("a" + "foo bar") + ((("hello" + ("foo bar" + "\n\t")) + "a") + "a")) {
        case "hello":
            v12 := GenS2{F0: 65535, F1: -692}
            _ = v12
        default:
            var v13 uint16 = 65535
            _ = v13
        }
        v3 = v3
        _ = fi11
    }
    var v14 int8 = 127
    v15 := []int{-9223372036854775808}
    v14 -= v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
