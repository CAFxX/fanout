package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    var v1 int16 = 968
    var v2 uint16 = 42
    var v3 uint16 = v2
    var v4 float64 = ((float64((((1.5 / 3.14159) - (3.14159 / 1e10)) + 0.1)) - 0.0) - float64(1.5))
    var v5 float32 = 0.0
    v4 += v4
    fmt.Println(v3)
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 9223372036854775807, F1: false}
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 1
    if (!true) {
        var v10 uint16 = 65535
        v9 -= 2
        _ = v10
    } else {
        v11 := GenS1{F0: -9223372036854775808, F1: false}
        v9 += 255
        for fi12 := 0; fi12 < 17; fi12++ {
            v2 -= 469
            if true {
                for fi13 := 0; fi13 < 1; fi13++ {
                    v3 -= (v3 / ((0) | 1))
                    var v14 int8 = -1
                    _ = fi13
                    _ = v14
                }
                v15 := []int{(fi12 & 2147483647), fi12, 556, 274}
                fmt.Println(v5)
                _ = v15
            }
            var v16 string = ("hello" + ((("" + (("" + "foo bar") + ("a" + "a"))) + "a") + ("foo bar" + "hello")))
            _ = fi12
            _ = v16
        }
        _ = v11
    }
    v3 += (v3 / (((v2 & (v3 / ((v3) | 1)))) | 1))
    fmt.Println(v9)
    var v17 uint32 = 1
    var v18 uint16 = 255
    v19 := []int{}
    v20 := GenS2{F0: -1, F1: 1000}
    v18 += v2
    switch "hello" {
    case "\n\t":
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    case "a":
        switch "foo bar" {
        case "":
            v1 += 875
        }
    default:
        v21 := GenS1{F0: 1, F1: true}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
