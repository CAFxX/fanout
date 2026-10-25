package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    var v1 int8 = 42
    var v2 uint32 = 111
    v3 := []int{-2147483648, -381, 1, -1}
    v4 := GenS1{F0: 2147483647, F1: 1000}
    v2 = v2
    v1 += 52
    if false {
        var v5 int32 = 0
        if false {
            var v6 float64 = -2.25
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(v5, v2)
                _ = fi7
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v6
        }
        _ = v5
    }
    var v8 bool = false
    if v8 {
        if ((!(0 > 255)) && v8) {
            v9 := GenS1{F0: 2147483647, F1: 32767}
            if len(v3) > 0 {
                v3[0] = -9223372036854775808
            }
            _ = v9
        }
    }
    var v10 int = -1
    v11 := GenS1{F0: 811, F1: 0}
    var v12 int16 = 42
    v12 += v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
