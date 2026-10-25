package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 bool }
func genFn0() int8 {
    var v7 uint = 0
    v7 += 383
    if false {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 120}
            _ = fi8
            _ = v9
        }
        v10 := []int{0, 42, -9223372036854775808, 1}
        _ = v10
    } else {
        v7 = (v7 + (v7 | (v7 + v7)))
        var v11 float64 = -2.25
        var v12 uint16 = 255
        _ = v11
        _ = v12
    }
    _ = v7
    return -1
}

func main() {
    v1 := []string{"foo bar", "foo bar", "foo bar", "a"}
    v2 := make(map[uint32]string)
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 65535
    var v5 int16 = 32767
    var v6 float32 = 1e10
    fmt.Println(v3, v6)
    v13 := GenS1{F0: 164}
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(v4, v3, fi14)
        var v15 uint64 = 18446744073709551615
        v6 = (((v6 / 0.1) / (v6 + v6)) / (float32(((3.14159 + (0.1 - -2.25)) / 0.1)) - v6))
        _ = fi14
        _ = v15
    }
    if true {
        v16 := GenS1{F0: 42}
        fmt.Println(v3, v3)
        _ = v16
    }
    v17 := GenS1{F0: 2}
    var v18 int32 = -1
    _ = len(v1)
    v4 = v4
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int64 = (v3 % (((v3 % ((v3) | 1))) | 1))
        var v21 int8 = -1
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        fmt.Println(v4)
        var v22 float64 = float64(v6)
        var v23 uint = 2
        _ = v22
        _ = v23
    } else {
        v6 += (v6 - -2.25)
        for fi24 := 0; fi24 < 7; fi24++ {
            v25 := GenS2{F0: 968, F1: (^(v4)), F2: true}
            _ = fi24
            _ = v25
        }
        fmt.Println(v18)
    }
    v26 := GenS2{F0: 18446744073709551615, F1: (^(v4)), F2: ((genFn0() > 49) || (1e10 >= (-2.25 / 1e10)))}
    var v27 string = ""
    v18 = v18
    v28 := GenS2{F0: 2, F1: 32767, F2: (3.14159 != 0.0)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
