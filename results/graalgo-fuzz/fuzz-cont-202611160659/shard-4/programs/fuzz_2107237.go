package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }

func main() {
    var v1 int8 = 0
    v2 := []uint32{}
    var v3 uint16 = 655
    v4 := make(map[string]int)
    var v5 bool = false
    v6 := make(map[uint32]bool)
    v5 = ((!(v5 || (545 != 0))) && (!false))
    v3 += v3
    v5 = ((255 <= v3) || (((3.14159 * ((0.0 - 0.0) / 1.5)) + (3.14159 + 1.5)) != (1.5 / 1.5)))
    if false {
        var v7 float32 = float32(((3.14159 + 3.14159) / 0.1))
        v8 := GenS1{F0: v3, F1: -888, F2: 42}
        _ = v7
        _ = v8
    } else {
        var v9 int16 = -1
        for fi10 := 0; fi10 < 13; fi10++ {
            v3 = 255
            v11 := GenS1{F0: 442, F1: fi10, F2: 2147483648}
            for fi12 := 0; fi12 < 1; fi12++ {
                if (0 >= 0) {
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v2)
                    default:
                        var v13 float64 = (1e10 - 0.0)
                        _ = v13
                    }
                    v14 := v4[""]
                    _ = v14
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    fmt.Println(v3, v5)
    fmt.Println(v5, v5)
    v15 := GenS1{F0: 898, F1: 68, F2: 65535}
    var v16 bool = true
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int32 = -633
        var v19 int16 = -1
        if true {
            v20 := GenS1{F0: v3, F1: (fi17 % (((-(fi17))) | 1)), F2: 0}
            _ = v20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    if v5 {
        var v21 int8 = v1
        _ = v21
    } else {
        v5 = v16
        v3 -= 607
        var v22 bool = false
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
