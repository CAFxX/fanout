package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
func genFn0() uint32 {
    var v6 bool = true
    v7 := GenS1{F0: 42, F1: 42}
    var v8 uint32 = 4294967295
    v9 := []int{42, 1, 127}
    v8 += v8
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1() int16 {
    v10 := []int{}
    switch 4294967295 {
    case 32768:
        var v11 float32 = 1.5
        _ = v11
    default:
        _ = len(v10)
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = v10
    return 925
}

func main() {
    var v1 uint32 = 489
    var v2 uint = 18446744073709551615
    var v3 uint16 = 0
    v4 := []int{729, -624}
    var v5 uint32 = (v1 - (v1 >> ((v2) & 63)))
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: (v2 + (v2 - v2)), F1: 1}
        for fi14 := 0; fi14 < 6; fi14++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v3 = v3
    v3 = v3
    v4 = append(v4, 2147483647)
    fmt.Println(v5)
    if false {
        fmt.Println(v2, v1)
        var v15 string = ("a" + "foo bar")
        _ = v15
    }
    if false {
        var v16 uint32 = 65535
        v2 = (v2 | (^(v2)))
        v4 = append(v4, 32768)
        _ = v16
    } else {
        v17 := GenS1{F0: v2, F1: (v3 ^ (v3 + v3))}
        if false {
            _ = len(v4)
            for fi18 := 0; fi18 < 17; fi18++ {
                v5 += (v1 / (((v5 * (v5 + (v5 - genFn0())))) | 1))
                _ = fi18
            }
            v3 = (v3 ^ v3)
        }
        _ = len(v4)
        _ = v17
    }
    v4 = append(v4, 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
