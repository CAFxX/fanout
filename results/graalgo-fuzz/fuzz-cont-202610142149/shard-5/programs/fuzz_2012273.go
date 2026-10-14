package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0() int {
    var v6 uint16 = 824
    fmt.Println(v6)
    var v7 bool = (4294967295 > 896)
    if true {
        var v8 float64 = 3.14159
        var v9 int64 = 9223372036854775807
        _ = v8
        _ = v9
    }
    fmt.Println(v7, v6)
    _ = v6
    _ = v7
    return -669
}

func main() {
    v1 := []int{}
    var v2 uint = 42
    var v3 uint32 = 42
    v4 := make(map[string]int)
    var v5 bool = false
    for fi10 := 0; fi10 < 11; fi10++ {
        switch fi10 {
        case 425:
            var v11 int8 = -1
            _ = v11
        }
        var v12 int = genFn0()
        _ = fi10
        _ = v12
    }
    if true {
        var v13 int8 = 0
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v3, v3)
        v2 -= v2
        _ = fi14
    }
    v15 := v4[""]
    v3 = (v3 + v3)
    switch genFn0() {
    case -329:
        v16 := GenS1{F0: "a", F1: "a", F2: 9223372036854775807}
        _ = v16
    case 0:
        _ = len(v1)
    default:
        v15 = (v15 ^ genFn0())
    }
    var v17 int = (v15 << ((4294967295) & 63))
    v2 = (v2 + v2)
    v15 -= (^(v17))
    v4["\n\t"] = v17
    switch -1 {
    case 543:
        var v18 int16 = 32767
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if ((128 != 42) || (v5 && ((714 >= 778) || false))) {
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS2{F0: 417, F1: (v2 >> (((v2 & v2)) & 63))}
                _ = fi20
                _ = v21
            }
            var v22 uint32 = v3
            v17 += (fi19 % ((v15) | 1))
            _ = v22
        } else {
            var v23 float32 = ((-2.25 * -2.25) / (0.0 + (1e10 + 1e10)))
            v3 += v3
            var v24 uint8 = 42
            _ = v23
            _ = v24
        }
        _ = fi19
    }
    v3 += v3
    var v25 uint64 = 256
    v26 := GenS2{F0: v2, F1: (v2 ^ v2)}
    v27 := GenS1{F0: ("" + "hello"), F1: "", F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
