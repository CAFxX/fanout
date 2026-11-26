package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 987
    var v2 float32 = (((1e10 / (1e10 / (1e10 / (((0.0 + 0.0) / 3.14159) - ((1e10 - ((((-2.25 + -2.25) + 1.5) - -2.25) + (float32(((1e10 * float64(1.5)) * ((1e10 / 1e10) * float64(1.5)))) - 3.14159))) + (1.5 - 3.14159)))))) - -2.25) - (1.5 - 1.5))
    var v3 uint8 = 143
    v4 := GenS1{F0: 255}
    v5 := GenS1{F0: 32767}
    v6 := GenS1{F0: 1}
    v7 := []int{128, 162, 42, -215}
    v3 = 30
    fmt.Println(v1)
    for fi8 := 0; fi8 < 17; fi8++ {
        v1 -= (v1 ^ (v1 << ((551) & 63)))
        _ = fi8
    }
    v1 += (v1 - v1)
    v1 = v1
    var v9 uint = 0
    switch 0 {
    case -9223372036854775808:
        var v10 string = "hello"
        _ = v10
    default:
        if true {
            switch ("\n\t" + (("" + ((("a" + "hello") + "\n\t") + "\n\t")) + "hello")) {
            case "a":
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int16 = 42
                    _ = fi11
                    _ = v12
                }
            }
            var v13 int32 = -1
            _ = v13
        }
    }
    var v14 float32 = (float32(-2.25) / 1e10)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 string = ("foo bar" + ("" + (("hello" + ("\n\t" + "a")) + "a")))
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
