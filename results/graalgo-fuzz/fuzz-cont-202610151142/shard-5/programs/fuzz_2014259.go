package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 string = "hello"
    var v2 float32 = (1.5 + (1.5 + -2.25))
    var v3 float32 = 1e10
    var v4 int32 = -1
    v5 := GenS1{F0: 0}
    var v6 float64 = 0.0
    v1 = "\n\t"
    v6 -= float64(0.0)
    v7 := GenS2{F0: 597, F1: int64(v4)}
    if ((true && (true && true)) || (1 == 251)) {
        v8 := []int{32768}
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v2 -= 3.14159
        if (v6 > v6) {
            var v10 uint32 = 2147483648
            v11 := []int{127, fi9}
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := []int{-106, fi12, (fi12 / ((fi12) | 1))}
                _ = len(v13)
                fmt.Println(fi12, v2, fi12)
                _ = fi12
                _ = v13
            }
        }
        _ = fi9
    }
    var v14 bool = (false && (!true))
    fmt.Println(v14, v1)
    v2 -= ((1.5 - (3.14159 / v3)) - (float32(v6) * v2))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
