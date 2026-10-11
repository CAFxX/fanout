package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint8 = 1
    var v2 uint32 = 65535
    var v3 float64 = 1e10
    v4 := []uint32{v2, v2, 4294967295}
    v5 := []uint32{}
    var v6 uint = 432
    var v7 uint16 = 1
    var v8 int32 = -1
    switch 128 {
    case -1:
        v9 := GenS1{F0: 1}
        _ = v9
    case 993:
        v10 := GenS2{F0: v1, F1: -128}
        _ = v10
    }
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v5)
    var v11 int32 = 303
    switch 2147483647 {
    case 2147483647:
        for fi12 := 0; fi12 < 14; fi12++ {
            switch (fi12 * (fi12 - fi12)) {
            case -1:
                var v13 uint = (v6 / ((v6) | 1))
                _ = v13
            case 128:
                for fi14 := 0; fi14 < 6; fi14++ {
                    _ = len(v5)
                    var v15 uint = (^(v6))
                    _ = fi14
                    _ = v15
                }
            case 32768:
                var v16 int8 = -30
                _ = v16
            default:
                var v17 float32 = 0.0
                _ = v17
            }
            var v18 string = ("foo bar" + "hello")
            v8 = -2147483648
            _ = fi12
            _ = v18
        }
    }
    v19 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
