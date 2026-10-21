package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 56
    v3 := []uint32{65535, 2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v4 int64 = -1
    fmt.Println(v1, v4)
    var v5 int16 = v1
    fmt.Println(v4, v4, v5)
    v6 := GenS1{F0: 0}
    if false {
        v7 := GenS2{F0: 0, F1: 682, F2: 2}
        _ = v7
    }
    var v8 float64 = float64(float32(0.1))
    v8 -= v8
    switch 42 {
    case 0:
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            if (100 == -1) {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                fmt.Println(v1, v4)
                v8 += ((v8 * v8) / 3.14159)
            }
        }
    }
    v1 = v5
    v9 := GenS2{F0: 2147483647, F1: -910, F2: 541}
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 string = "a"
        for fi12 := 0; fi12 < 7; fi12++ {
            v2 -= 87
            fmt.Println(v2, fi10)
            fmt.Println(v4)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
