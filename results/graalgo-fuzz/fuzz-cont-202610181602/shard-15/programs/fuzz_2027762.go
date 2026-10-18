package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
func genFn0() uint8 {
    if true {
        var v6 int = -290
        v6 = 1
        v6 = v6
        _ = v6
    }
    var v7 int = 128
    v7 += v7
    _ = v7
    return 171
}
func genFn1() int16 {
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 15; fi9++ {
        switch fi9 {
        case 615:
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(fi9, v8, fi10)
                _ = fi10
            }
        }
        _ = fi9
    }
    v11 := []int{-9223372036854775808, 2147483647}
    _ = v11
    _ = v8
    return 32767
}

func main() {
    var v1 int = 179
    var v2 int8 = -124
    var v3 uint = 42
    var v4 float32 = 3.14159
    var v5 int = -1
    var v12 int = (v1 * int(v3))
    switch 431 {
    case -752:
        if true {
            v13 := []int{(-(v12))}
            _ = v13
        }
    case 893:
        v14 := []int{2147483647, v12, (v5 - 1), -983}
        _ = v14
    case -1:
        v15 := GenS1{F0: 220, F1: 1, F2: (v3 & (v3 + (v3 << ((v3) & 63))))}
        _ = v15
    }
    switch v12 {
    case -9223372036854775808:
        var v16 string = ("\n\t" + ("a" + ""))
        _ = v16
    }
    v17 := []int{v5, 4294967295, (v1 + v5), 42}
    fmt.Println(v2, v1, v4)
    v12 += 2147483647
    fmt.Println(v1, v1, v3)
    var v18 int16 = -1
    var v19 float64 = float64(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
