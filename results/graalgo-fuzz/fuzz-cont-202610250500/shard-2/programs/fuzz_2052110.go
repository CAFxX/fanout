package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    var v1 uint16 = 1
    v2 := []int{}
    var v3 uint32 = 711
    var v4 int = 9223372036854775807
    var v5 uint = 240
    if false {
        v6 := GenS1{F0: 9223372036854775807, F1: 42}
        var v7 float64 = (0.0 * (1.5 * 1e10))
        var v8 uint = (v5 - (v5 % ((v5) | 1)))
        _ = v6
        _ = v7
        _ = v8
    } else {
        if false {
            var v9 int16 = -971
            fmt.Println(v1, v4, v1)
            v3 = (v3 + (v3 ^ (v3 | uint32(v5))))
            _ = v9
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v11 float32 = 3.14159
            _ = fi10
            _ = v11
        }
    }
    v4 = v4
    v12 := GenS1{F0: -1, F1: -1}
    var v13 uint = (v5 / ((v5) | 1))
    v5 = 18446744073709551615
    var v14 uint = (v5 >> ((v13) & 63))
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int = 32767
        if (false && true) {
            if (-1 < 28) {
                var v17 uint32 = (v3 * v3)
                _ = v17
            }
        }
        v18 := GenS1{F0: 9223372036854775807, F1: -327}
        _ = fi15
        _ = v16
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint16 = (v1 | v1)
        var v21 uint32 = (v3 | v3)
        _ = fi19
        _ = v20
        _ = v21
    }
    v13 = (v5 << ((v14) & 63))
    var v22 int16 = 0
    var v23 uint32 = v3
    if false {
        var v24 uint = v14
        v25 := GenS1{F0: 390, F1: 9223372036854775807}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v13 += 4294967295
        for fi27 := 0; fi27 < 18; fi27++ {
            v28 := GenS1{F0: -360, F1: -1}
            fmt.Println(v23, v5)
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 6; fi29++ {
            _ = len(v2)
            _ = fi29
        }
        _ = fi26
    }
    if len(v2) > 0 {
        v2[0] = (^(v4))
    }
    var v30 float32 = 0.1
    var v31 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
