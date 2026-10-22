package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }

func main() {
    var v1 float32 = float32((-2.25 - -2.25))
    var v2 int16 = 1
    var v3 int16 = v2
    v4 := make(map[string]int)
    v1 -= v1
    v3 += -653
    fmt.Println(v1)
    v2 += v3
    v5 := GenS1{F0: 18446744073709551615, F1: 1, F2: 127}
    v6 := []int{-845, 4294967295, 4294967295}
    v3 += v2
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int32 = 991
        _ = fi7
        _ = v8
    }
    fmt.Println(v1)
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint = 255
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(v2)
        for fi12 := 0; fi12 < 15; fi12++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
