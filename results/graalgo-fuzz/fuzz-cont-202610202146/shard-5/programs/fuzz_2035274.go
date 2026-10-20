package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float32 = -2.25
    var v2 int = 529
    var v3 uint8 = 157
    var v4 int8 = 0
    var v5 int32 = 136
    v6 := make(map[uint32]uint32)
    v3 = (v3 - v3)
    var v7 uint16 = 0
    v4 = v4
    v5 = v5
    if true {
        v8 := GenS1{F0: 0, F1: 42}
        v4 += v4
        v9 := GenS1{F0: 0, F1: 1}
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = (v3 & 36)
        var v11 float64 = -2.25
        switch ("a" + "foo bar") {
        case "hello":
            v10 = v10
        }
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: 1, F1: uint64(v3)}
        if (v5 == (v5 / (((-(v5))) | 1))) {
            if false {
                v14 := GenS1{F0: 1, F1: 18446744073709551615}
                var v15 uint64 = 883
                _ = v14
                _ = v15
            } else {
                var v16 uint32 = 4294967295
                _ = v16
            }
            v4 = v4
        } else {
            v17 := GenS1{F0: 0, F1: 1}
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    var v18 bool = ((1 >= 0) && true)
    fmt.Println(v1, v7)
    switch (v2 - v2) {
    case 2147483647:
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := []int{-246, (v2 - fi19)}
            v5 -= v5
            if ((((((((1e10 / v1) * 0.1) / (v1 + v1)) / (v1 * (v1 + -2.25))) + (3.14159 - v1)) - v1) <= 3.14159) || v18) {
                v4 = v4
            }
            _ = fi19
            _ = v20
        }
    case 42:
        v21 := GenS1{F0: -1, F1: 4294967295}
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 uint32 = 42
        var v24 uint8 = (v3 >> ((594) & 63))
        var v25 float32 = 0.0
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v2 -= v2
    var v26 uint16 = uint16(v3)
    var v27 uint8 = v3
    v4 = (^(v4))
    v28 := []int{v2, v2, int(v5), (-(v2))}
    var v29 uint64 = 931
    if ((v4 + -128) >= 110) {
        v29 -= 2
        fmt.Println(v26, v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
