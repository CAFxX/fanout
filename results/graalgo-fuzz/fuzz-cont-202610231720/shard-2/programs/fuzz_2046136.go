package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int32 = 1
    var v2 int16 = 305
    v3 := make(map[string]string)
    var v4 int = -729
    v5 := make(map[string]int)
    var v6 int16 = v2
    fmt.Println(v6, v2, v6)
    switch -30 {
    case 4294967295:
        var v7 uint16 = 42
        _ = v7
    case 127:
        var v8 float64 = (1e10 + (-2.25 / -2.25))
        _ = v8
    case 128:
        v4 = (v4 % (((v4 << ((541) & 63))) | 1))
    default:
        v4 -= 29
    }
    var v9 uint64 = 2
    for fi10 := 0; fi10 < 6; fi10++ {
        v1 = v1
        _ = fi10
    }
    v4 = v4
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int16 = int16(v1)
        var v13 int8 = 116
        v4 = v4
        _ = fi11
        _ = v12
        _ = v13
    }
    v4 -= v4
    var v14 uint8 = 206
    v1 -= v1
    for fi15 := 0; fi15 < 9; fi15++ {
        v14 = 255
        if ((1 < 366) && true) {
            v16 := GenS1{F0: (v9 ^ 1)}
            _ = v16
        } else {
            var v17 int = (fi15 % (((^(fi15))) | 1))
            _ = v17
        }
        delete(v3, ("\n\t" + "hello"))
        _ = fi15
    }
    if ((0.1 < ((((-2.25 * (1.5 / 1.0)) + -2.25) * (0.1 - -2.25)) + 1.5)) && false) {
        v18 := GenS1{F0: 4294967295}
        var v19 uint32 = 430
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
