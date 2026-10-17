package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint64 }
func genFn0() int64 {
    var v6 int64 = -358
    var v7 int64 = 0
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    var v3 int8 = 11
    var v4 int8 = (^(v3))
    v5 := make(map[uint32]string)
    if false {
        v8 := []int{128}
        _ = v8
    } else {
        var v9 int = 128
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            v2 = (-(v2))
            _ = fi11
        }
        _ = fi10
    }
    var v12 uint32 = 577
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint = 42
        fmt.Println(v4)
        _ = fi13
        _ = v14
    }
    var v15 uint = 918
    for fi16 := 0; fi16 < 17; fi16++ {
        v4 = (v3 << ((39) & 63))
        _ = fi16
    }
    switch (((("hello" + "a") + "a") + "hello") + ("\n\t" + ("\n\t" + ""))) {
    case "a":
        var v17 bool = (true && (!((-407 > 42) || (0 >= 192))))
        _ = v17
    default:
        var v18 bool = (v15 < uint(v12))
        _ = v18
    }
    if true {
        fmt.Println(v2)
        v19 := []int{727, -9223372036854775808, 951, 341}
        _ = v19
    }
    v4 += (-(v4))
    v20 := GenS1{F0: 42, F1: 32768, F2: 42}
    if (1.5 > ((1e10 + ((0.0 - ((1.5 / 1e10) / -2.25)) / -2.25)) * (0.1 + 1e10))) {
        var v21 uint64 = 18446744073709551615
        v22 := []int{-9223372036854775808, 32767}
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            v2 += 42
            if ((fi23 | 649) > fi23) {
                var v24 uint32 = v12
                _ = v24
            }
            fmt.Println(v3, v15)
            _ = fi23
        }
        if (true || (1 > 795)) {
            var v25 int8 = v4
            if true {
                v26 := GenS1{F0: 32767, F1: -300, F2: 199}
                var v27 uint32 = (v12 % ((631) | 1))
                var v28 uint16 = 272
                _ = v26
                _ = v27
                _ = v28
            }
            switch 4294967295 {
            case -1:
                var v29 bool = false
                _ = v29
            default:
                fmt.Println(v25)
            }
            _ = v25
        }
    }
    var v30 float32 = 0.0
    v31 := GenS2{F0: 0, F1: 1, F2: 212}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
