package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
func genFn0() int16 {
    var v6 uint64 = 766
    fmt.Println(v6)
    v7 := []int{-9223372036854775808, 128}
    v6 += 629
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint8, p1 uint32) uint64 {
    p1 -= (p1 >> ((uint(p0)) & 63))
    v8 := GenS1{F0: 65535, F1: -873, F2: 42}
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: 42, F2: 65535}
        var v11 uint16 = 976
        fmt.Println(fi9, fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    return 354
}

func main() {
    var v1 float64 = (-2.25 * -2.25)
    var v2 int32 = -2147483648
    var v3 int = 127
    var v4 int32 = 0
    v5 := make(map[string]int)
    fmt.Println(v1, v3, v2)
    v4 -= v2
    fmt.Println(v3, v2, v4)
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := []int{(fi12 * 32768), 127, (v3 & (fi12 / ((v3) | 1)))}
        v14 := v5["hello"]
        _ = fi12
        _ = v13
        _ = v14
    }
    switch (v3 % (((-(v3))) | 1)) {
    case 4294967295:
        v1 += v1
    case 757:
        var v15 int16 = -32768
        _ = v15
    default:
        v4 -= (v4 / (((-(v4))) | 1))
    }
    v16 := GenS1{F0: 255, F1: v2, F2: 65535}
    fmt.Println(v3, v4, v1)
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: 2, F1: (v4 / (((v2 % ((-2147483648) | 1))) | 1)), F2: 65535}
        _ = fi17
        _ = v18
    }
    v1 = (v1 + v1)
    v3 -= (v3 << ((4294967295) & 63))
    v1 = v1
    switch (v3 / ((0) | 1)) {
    case -621:
        if (!false) {
            fmt.Println(v3, v3)
        }
    default:
        switch -2147483648 {
        case 32767:
            v1 = (v1 / 1e10)
        }
    }
    var v19 uint = 653
    if false {
        var v20 uint = (v19 / ((0) | 1))
        _ = v20
    }
    v1 -= v1
    v19 += v19
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
