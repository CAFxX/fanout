package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint }

func main() {
    var v1 bool = ((float64((((0.0 + -2.25) - float32(1.5)) + 0.0)) * (float64(float32(0.1)) / (3.14159 - (0.0 - 1e10)))) <= float64((1e10 - -2.25)))
    v2 := make(map[uint32]int)
    var v3 int = 9223372036854775807
    var v4 string = "a"
    for fi5 := 0; fi5 < 17; fi5++ {
        v3 = (fi5 | fi5)
        _ = fi5
    }
    var v6 string = ""
    var v7 uint16 = 88
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 7; fi9++ {
        v1 = (2 < 2)
        _ = fi9
    }
    if (111 < 42) {
        if false {
            var v10 uint16 = v7
            v10 += v7
            _ = v10
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint32 = 306
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 1, F1: (v1 || ((false || (v4 >= v6)) && (!false))), F2: 0}
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            v6 = v6
            v4 = v6
            _ = fi15
        }
        _ = fi14
    }
    v16 := GenS2{F0: 509, F1: 185, F2: 919}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
