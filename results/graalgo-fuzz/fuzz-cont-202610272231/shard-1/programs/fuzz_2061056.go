package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }
func genFn0() uint16 {
    var v6 string = ""
    var v7 int8 = 42
    var v8 int16 = -153
    if true {
        v9 := []int{}
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(v7)
            _ = fi10
        }
        _ = v9
    }
    v11 := GenS1{F0: v8, F1: 18446744073709551615, F2: -1}
    _ = v11
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := []int{-686}
    var v2 bool = (!false)
    var v3 uint8 = 2
    var v4 int64 = 42
    var v5 uint8 = 84
    v12 := GenS2{F0: 0, F1: -1}
    if v2 {
        v13 := GenS2{F0: 472, F1: 2147483647}
        for fi14 := 0; fi14 < 18; fi14++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v16 bool = v2
                _ = fi15
                _ = v16
            }
            var v17 uint = 0
            _ = fi14
            _ = v17
        }
        _ = v13
    } else {
        v1 = append(v1, 676)
        v18 := GenS3{F0: 499, F1: false, F2: false}
        _ = v18
    }
    v4 += v4
    if v2 {
        var v19 float64 = (((-2.25 * 3.14159) - (1.5 * ((1e10 - float64(0.0)) - (1.5 - 1.5)))) * 0.1)
        var v20 int = int(v5)
        var v21 uint32 = 369
        _ = v19
        _ = v20
        _ = v21
    } else {
        var v22 uint16 = 255
        _ = v22
    }
    if v2 {
        v3 = (v3 - 100)
    }
    v1 = append(v1, 576)
    var v23 bool = false
    var v24 uint = 4294967295
    fmt.Println(v23, v4, v24)
    switch 32767 {
    case 4294967295:
        fmt.Println(v3, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
