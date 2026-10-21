package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int16) uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 -= p0
        _ = fi4
    }
    var v5 int16 = -392
    p0 -= (v5 | p0)
    v5 -= p0
    _ = p0
    _ = v5
    return 13
}
func genFn1(p0 uint64, p1 uint32) int8 {
    p0 -= p0
    var v6 uint32 = (p1 << ((1) & 63))
    if false {
        v7 := []int{-169, 32768, -1}
        fmt.Println(v6)
        if false {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            fmt.Println(p1, p0)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -937
    v3 := make(map[string]uint32)
    var v8 int16 = -1
    for fi9 := 0; fi9 < 3; fi9++ {
        if true {
            var v10 int32 = -2147483648
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(v8)
                fmt.Println(v8, fi9)
                _ = fi11
            }
            _ = v10
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: 1, F1: 1}
            v14 := GenS1{F0: 255, F1: 85}
            fmt.Println(fi12, v2, fi12)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        if false {
            var v16 int8 = 43
            var v17 int8 = v16
            _ = v16
            _ = v17
        }
        v8 = v8
        _ = fi15
    }
    v18 := GenS1{F0: 2, F1: 826}
    var v19 string = "hello"
    fmt.Println(v19)
    v20 := v1[867]
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v19, v19)
        var v22 uint32 = v20
        _ = fi21
        _ = v22
    }
    fmt.Println(v20)
    switch (v2 + -9223372036854775808) {
    case 42:
        v8 = (v8 ^ v8)
    default:
        if true {
            v19 = (v19 + v19)
            v20 -= v20
        }
    }
    v20 = (v20 - v20)
    var v23 uint8 = 0
    v19 = v19
    switch 1 {
    case -1:
        v19 = ("hello" + "a")
    default:
        for fi24 := 0; fi24 < 1; fi24++ {
            v2 += -1
            _ = fi24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v8
}
