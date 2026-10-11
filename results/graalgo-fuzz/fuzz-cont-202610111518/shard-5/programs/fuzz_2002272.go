package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint = 18446744073709551615
    var v2 string = "\n\t"
    var v3 uint8 = 206
    v4 := []int{}
    var v5 int32 = -679
    v3 = v3
    for fi6 := 0; fi6 < 10; fi6++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v7 := GenS1{F0: v1}
        var v8 int16 = -673
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 bool = (!(1.5 <= ((((0.0 + float32(3.14159)) / -2.25) - float32(1.5)) * ((float32(1e10) + (0.0 / ((float32(float64(-2.25)) + 0.0) - ((0.0 * (-2.25 + 3.14159)) - float32(1e10))))) * 3.14159))))
        _ = fi9
        _ = v10
    }
    if len(v4) > 0 {
        v4[0] = -1
    }
    if false {
        var v11 uint16 = 42
        _ = v11
    } else {
        for fi12 := 0; fi12 < 6; fi12++ {
            _ = len(v4)
            _ = fi12
        }
        v5 = 0
        v3 -= (v3 / ((177) | 1))
    }
    var v13 float64 = 1.5
    for fi14 := 0; fi14 < 5; fi14++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
    }
    v13 += float64(float32(-2.25))
    v2 = ("a" + v2)
    v15 := GenS1{F0: uint(v3)}
    var v16 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
