package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0() uint8 {
    var v5 int8 = -1
    v5 -= v5
    for fi6 := 0; fi6 < 1; fi6++ {
        v5 += 0
        _ = fi6
    }
    v5 -= v5
    _ = v5
    return 0
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = 0
    var v3 uint32 = 1
    var v4 int32 = (^(v2))
    v7 := GenS1{F0: 729, F1: 2, F2: (!false)}
    v8 := []int{}
    _ = len(v8)
    v4 = -2147483648
    v4 += (v4 + (v2 | (^(v4))))
    for fi9 := 0; fi9 < 16; fi9++ {
        _ = len(v8)
        _ = fi9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        switch 460 {
        case 1:
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 float64 = (((1.5 - 0.0) + 1.5) - 0.1)
                v3 += v3
                v2 = (v4 - (v2 + (v2 ^ v4)))
                v13 := GenS1{F0: 2, F1: (v1 % (((v1 << ((uint(v3)) & 63))) | 1)), F2: true}
                fmt.Println(v1, v2, v3)
                _ = fi11
                _ = v12
                _ = v13
            }
        case -523:
            var v14 int8 = -54
            _ = v14
        }
        v15 := GenS2{F0: (int16(v3) < 1000), F1: 32}
        if false {
            v1 = 4294967295
        }
        _ = fi10
        _ = v15
    }
    v3 = (v3 % ((v3) | 1))
    if len(v8) > 0 {
        v8[0] = -2147483648
    }
    v16 := GenS2{F0: true, F1: (v3 >> ((585) & 63))}
    switch "\n\t" {
    case "":
        if false {
            var v17 string = "hello"
            _ = v17
        } else {
            var v18 uint32 = v3
            if len(v8) > 0 {
                v8[0] = 285
            }
            _ = v18
        }
    case "a":
        v3 = v3
    }
    v3 = (v3 - (v3 / ((42) | 1)))
    if ((!(4294967295 < 19)) || (false || false)) {
        var v19 uint8 = 2
        _ = v19
    } else {
        var v20 bool = (-2147483648 != 32768)
        switch "\n\t" {
        case "a":
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 uint16 = 746
                _ = fi21
                _ = v22
            }
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
