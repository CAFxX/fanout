package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool; F1 uint32; F2 uint }

func main() {
    var v1 uint = 153
    var v2 string = ("" + ("\n\t" + "foo bar"))
    var v3 float32 = -2.25
    v4 := GenS1{F0: 0, F1: 1}
    v3 = v3
    v2 = v2
    v5 := []int{165, -9223372036854775808, 128, 42}
    v6 := GenS1{F0: -693, F1: 545}
    if (!true) {
        switch "foo bar" {
        case "hello":
            v1 += v1
        case "":
            var v7 int16 = -32768
            _ = v7
        default:
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 float64 = 0.0
                if (3.14159 >= (v3 - v3)) {
                    var v10 uint8 = 0
                    _ = v10
                } else {
                    var v11 float64 = v9
                    var v12 int = fi8
                    v13 := GenS2{F0: (("a" >= ((v2 + "a") + v2)) || (774 != -392)), F1: 219, F2: (^(v1))}
                    _ = v11
                    _ = v12
                    _ = v13
                }
                fmt.Println(v2)
                _ = fi8
                _ = v9
            }
        }
    }
    var v14 float32 = 1e10
    var v15 float64 = 0.1
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v16 float32 = (v14 * v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
