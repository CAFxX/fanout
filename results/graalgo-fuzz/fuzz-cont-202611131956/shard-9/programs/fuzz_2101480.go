package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int8 }
func genFn0() string {
    var v4 int8 = -65
    var v5 int8 = 100
    var v6 int32 = 3
    v7 := []int{}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return ("" + "a")
}
func genFn1() int32 {
    if false {
        v8 := GenS1{F0: 42, F1: 65535, F2: 42}
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(fi9)
            var v10 int32 = 2147483647
            _ = fi9
            _ = v10
        }
        var v11 string = ("hello" + "foo bar")
        _ = v11
        _ = v8
    } else {
        var v12 uint64 = 220
        _ = v12
    }
    v13 := []int{128, 32768, 1, 32767}
    var v14 bool = true
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int8 = 103
        var v17 float64 = 1e10
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int64 = int64(v1)
    v19 := GenS1{F0: 128, F1: 2147483647, F2: 0}
    v1 -= 1
    for fi20 := 0; fi20 < 10; fi20++ {
        v1 -= (v1 - fi20)
        v18 += v18
        _ = fi20
    }
    v21 := []int{v1}
    v22 := GenS1{F0: 128, F1: 319, F2: 127}
    if len(v21) > 0 {
        v21[0] = v1
    }
    v23 := GenS1{F0: 231, F1: 4294967295, F2: 0}
    if true {
        v2[(^(v1))] = 2
    } else {
        v1 = v1
    }
    if ((v18 ^ (v18 & (v18 | (-(v18))))) == (v18 / ((v18) | 1))) {
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS1{F0: 1, F1: 65535, F2: 66}
            _ = fi24
            _ = v25
        }
        var v26 int8 = 127
        _ = v26
    } else {
        v18 = v18
        switch v1 {
        case -937:
            v21 = append(v21, -2147483648)
        default:
            v1 = (v1 & (v1 >> ((86) & 63)))
        }
    }
    var v27 uint = 11
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
}
