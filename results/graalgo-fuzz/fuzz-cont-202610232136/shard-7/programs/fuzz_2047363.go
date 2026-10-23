package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 65535
    var v3 uint32 = 65535
    v4 := []uint32{v3}
    v5 := make(map[string]string)
    v6 := []uint32{}
    delete(v5, "\n\t")
    if ((false || ((true && true) && (18446744073709551615 < 255))) && (1 < 85)) {
        v2 += v2
        var v7 int16 = 0
        switch 9223372036854775807 {
        case -249:
            v1 -= (v1 * v1)
        case -2147483648:
            var v8 uint32 = 2
            _ = v8
        default:
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint8 = 127
                v1 += (^(v1))
                _ = fi9
                _ = v10
            }
        }
        _ = v7
    } else {
        var v11 uint16 = (^(v2))
        var v12 uint64 = 0
        _ = v11
        _ = v12
    }
    v13 := v5["foo bar"]
    fmt.Println(v2)
    var v14 uint8 = 152
    v15 := GenS1{F0: 0, F1: v1, F2: 18446744073709551615}
    v16 := GenS1{F0: 42, F1: (v1 >> ((1) & 63)), F2: 815}
    v17 := GenS1{F0: 18446744073709551615, F1: v1, F2: 2}
    _ = len(v5)
    for fi18 := 0; fi18 < 9; fi18++ {
        v1 = v1
        fmt.Println(v13, v3, v13)
        _ = fi18
    }
    var v19 float64 = -2.25
    for fi20 := 0; fi20 < 1; fi20++ {
        if false {
            v19 += 0.0
        } else {
            var v21 int32 = 2147483647
            fmt.Println(v1, v1, v2)
            _ = v21
        }
        var v22 int8 = 42
        _ = fi20
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
