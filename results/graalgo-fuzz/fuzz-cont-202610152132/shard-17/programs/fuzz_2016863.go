package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []int{-2147483648, -1, 0}
    v2 := []int{-747, -9223372036854775808, -1}
    var v3 float32 = 0.1
    for fi4 := 0; fi4 < 5; fi4++ {
        v3 += ((-2.25 * v3) + float32(0.0))
        v3 += 0.1
        if len(v1) > 0 {
            v1[0] = fi4
        }
        _ = fi4
    }
    v3 = v3
    if false {
        var v5 int8 = 93
        var v6 uint32 = 874
        _ = v5
        _ = v6
    }
    var v7 uint = 1
    var v8 string = "a"
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint32 = 377
        if false {
            v10 = v10
            v1 = append(v1, (fi9 & (^(fi9))))
        }
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v8 = (v8 + v8)
        var v12 int16 = 1
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 497, F1: (v7 + 4294967295)}
    if false {
        for fi14 := 0; fi14 < 2; fi14++ {
            v7 = (v7 | (v7 >> ((255) & 63)))
            _ = fi14
        }
        fmt.Println(v3, v3)
    }
    switch v8 {
    case "hello":
        v3 -= 0.1
    case "\n\t":
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS2{F0: 255, F1: -1}
            v17 := GenS2{F0: 2, F1: 42}
            var v18 float64 = (3.14159 / -2.25)
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    case "":
        switch -2147483648 {
        case 32767:
            var v19 float64 = (1e10 * (3.14159 / 1e10))
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        _ = len(v1)
        _ = fi20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 int = 42
            var v24 float32 = v3
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    if len(v2) > 0 {
        v2[0] = -526
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
