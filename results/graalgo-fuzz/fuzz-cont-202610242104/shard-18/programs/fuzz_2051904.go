package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint8 }
func genFn0() uint32 {
    v5 := GenS1{F0: 0, F1: 42, F2: 293}
    var v6 float64 = -2.25
    if false {
        fmt.Println(v6)
    } else {
        if true {
            var v7 float32 = 0.1
            fmt.Println(v7)
            _ = v7
        } else {
            v8 := GenS1{F0: 121, F1: 2, F2: 255}
            _ = v8
        }
    }
    var v9 float64 = float64(-2.25)
    _ = v5
    _ = v6
    _ = v9
    return 4294967295
}
func genFn1() int16 {
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint64 = 18446744073709551615
            var v13 uint32 = 2147483647
            v13 += (^(v13))
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 int16 = 32767
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS1{F0: -9223372036854775808, F1: 255, F2: 685}
        _ = fi15
        _ = v16
    }
    fmt.Println(v14)
    _ = v14
    return (-(v14))
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = 4294967295
    v3 := make(map[string]bool)
    var v4 int16 = -1
    var v17 int16 = -144
    var v18 float32 = 0.0
    v4 = (v4 << (((^(v1))) & 63))
    var v19 int8 = 0
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS2{F0: 42}
        v22 := GenS2{F0: 26}
        v19 += -128
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint32 = 920
    v4 += v17
    for fi24 := 0; fi24 < 6; fi24++ {
        v18 = (v18 - (((v18 + v18) * v18) + (3.14159 / 0.1)))
        switch "hello" {
        case "a":
            fmt.Println(v17, v4, fi24)
        }
        _ = fi24
    }
    var v25 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
