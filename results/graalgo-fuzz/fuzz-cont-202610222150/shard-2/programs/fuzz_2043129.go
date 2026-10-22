package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint8 = 1
    var v2 uint = 2
    var v3 float32 = 0.1
    v4 := GenS1{F0: -118, F1: 202, F2: -349}
    v3 = v3
    fmt.Println(v3)
    if false {
        switch "" {
        case "hello":
            var v5 int = 128
            _ = v5
        default:
            v2 = (v2 | (v2 << ((uint(v1)) & 63)))
        }
        fmt.Println(v1)
    } else {
        v6 := GenS1{F0: 127, F1: 0, F2: -526}
        _ = v6
    }
    v2 = (v2 % ((793) | 1))
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{fi7, (fi7 ^ fi7), fi7, (fi7 % (((fi7 << ((922) & 63))) | 1))}
        v9 := GenS2{F0: true}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = 42
    v3 = -2.25
    v10 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
