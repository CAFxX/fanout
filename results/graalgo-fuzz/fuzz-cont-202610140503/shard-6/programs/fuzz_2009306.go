package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int32; F1 string; F2 int8 }

func main() {
    var v1 int8 = -128
    var v2 float64 = 0.0
    var v3 uint = 1
    for fi4 := 0; fi4 < 17; fi4++ {
        v1 -= (v1 << (((^(v3))) & 63))
        var v5 uint8 = 0
        v5 -= (^(v5))
        _ = fi4
        _ = v5
    }
    var v6 float64 = (v2 + (((1e10 * ((0.0 * float64(1e10)) / v2)) * 0.1) * ((v2 / (1.5 - (((v2 - v2) + v2) + v2))) - (v2 + ((float64(0.1) - float64(1e10)) / v2)))))
    v7 := GenS1{F0: 57, F1: 1}
    if true {
        v2 += (-2.25 - (v6 - 3.14159))
    } else {
        var v8 int = 9223372036854775807
        var v9 string = "foo bar"
        _ = v8
        _ = v9
    }
    v2 += (1e10 / -2.25)
    v10 := []int{42, 2147483647, 397}
    v6 = (v6 * 3.14159)
    var v11 float32 = (0.1 / 3.14159)
    for fi12 := 0; fi12 < 2; fi12++ {
        v11 = 1.5
        v11 = v11
        var v13 uint64 = 494
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 42, F1: 446}
    v15 := GenS3{F0: 245, F1: ((("hello" + "\n\t") + "") + "\n\t"), F2: (-(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
