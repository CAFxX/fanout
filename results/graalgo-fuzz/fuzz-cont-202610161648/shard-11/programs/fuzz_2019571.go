package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }

func main() {
    var v1 uint8 = 245
    var v2 float64 = ((-2.25 * 0.0) / (float64(1e10) + -2.25))
    var v3 int = -645
    var v4 int16 = -32768
    v4 = v4
    var v5 float32 = ((((1e10 / float32(1.5)) * (1e10 + float32(v2))) * (float32(0.0) * (float32((v2 - v2)) + 1e10))) - 1e10)
    v6 := []int{(v3 * v3), v3}
    v2 -= v2
    v1 -= 128
    var v7 bool = (!(534 != -1))
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float64 = v2
        fmt.Println(v4, v1)
        v2 = float64((0.0 - 1e10))
        _ = fi8
        _ = v9
    }
    v5 += -2.25
    switch ("\n\t" + "a") {
    case "a":
        var v10 int16 = v4
        _ = v10
    case "\n\t":
        if false {
            v11 := GenS1{F0: 122, F1: (("\n\t" + ("\n\t" + "\n\t")) >= ("foo bar" + "hello")), F2: -44}
            _ = v11
        }
    }
    v1 -= 255
    for fi12 := 0; fi12 < 10; fi12++ {
        v1 -= (^(v1))
        v1 = v1
        var v13 uint = 4294967295
        _ = fi12
        _ = v13
    }
    var v14 int64 = 0
    v1 = v1
    v15 := GenS1{F0: 100, F1: true, F2: 100}
    var v16 float64 = 1.5
    var v17 int = (v3 ^ (v3 & v3))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
