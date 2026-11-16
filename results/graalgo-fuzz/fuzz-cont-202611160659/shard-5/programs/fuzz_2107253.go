package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
func genFn0(p0 uint8, p1 int) int16 {
    p1 = (p1 << ((871) & 63))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return -32768
}
func genFn1(p0 uint16, p1 int16) uint32 {
    var v5 float32 = 1e10
    var v6 float32 = v5
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{-2147483648, (fi7 / (((-(fi7))) | 1)), (fi7 - fi7), 929}
        var v9 float32 = v5
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 -= ((v6 / (1.5 - float32(-2.25))) * v5)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 string = "hello"
    var v2 bool = true
    var v3 int8 = 100
    v4 := make(map[string]uint32)
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v1, v1, v2)
        _ = fi10
    }
    if true {
        v3 -= (v3 % ((v3) | 1))
        v3 = (-(v3))
    } else {
        v2 = ((994 != 42) && false)
    }
    v11 := GenS1{F0: 1, F1: 123, F2: false}
    v2 = ((((("hello" + ("\n\t" + ("\n\t" + v1))) + v1) + v1) + v1) <= v1)
    var v12 uint64 = 42
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint16 = 42
        var v15 uint32 = 42
        _ = fi13
        _ = v14
        _ = v15
    }
    switch "" {
    case "":
        var v16 uint32 = 2
        _ = v16
    case "\n\t":
        _ = len(v4)
    case "foo bar":
        var v17 uint = 1
        _ = v17
    default:
        if v2 {
            var v18 int16 = 32767
            v4["\n\t"] = 0
            _ = v18
        } else {
            var v19 uint8 = 30
            v20 := GenS1{F0: 516, F1: (v19 - v19), F2: ((v12 << ((0) & 63)) <= (v12 << ((0) & 63)))}
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := []int{(fi21 & (fi21 & -1)), fi21}
                fmt.Println(v2)
                fmt.Println(v19, v19)
                _ = fi21
                _ = v22
            }
            _ = v19
            _ = v20
        }
    }
    var v23 int16 = -839
    switch (v1 + v1) {
    case "\n\t":
        var v24 bool = v2
        _ = v24
    case "hello":
        v25 := GenS2{F0: v3, F1: -9223372036854775808, F2: 0}
        _ = v25
    case "a":
        fmt.Println(v12)
    default:
        switch (v1 + ("\n\t" + (((v1 + v1) + v1) + (v1 + ("" + v1))))) {
        case "a":
            v26 := GenS2{F0: v3, F1: int64(v3), F2: 1}
            _ = v26
        }
    }
    if (v2 || v2) {
        v1 = "a"
    } else {
        v27 := []int{-857, 32767}
        fmt.Println(v3, v23)
        v2 = (-2.25 <= 1.5)
        _ = v27
    }
    v28 := GenS1{F0: 4294967295, F1: 128, F2: ((true || v2) && true)}
    v1 = ("" + v1)
    v29 := GenS2{F0: v3, F1: -9223372036854775808, F2: 58}
    v30 := GenS2{F0: -124, F1: 42, F2: 127}
    v4[v1] = 249
    switch v1 {
    case "foo bar":
        v31 := GenS1{F0: 2, F1: 255, F2: v2}
        _ = v31
    case "":
        var v32 int32 = 42
        _ = v32
    case "hello":
        var v33 int64 = 0
        _ = v33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        if v2 {
            var v35 uint16 = 65535
            v2 = v2
            _ = v35
        } else {
            var v36 uint64 = 0
            _ = v36
        }
        _ = fi34
    }
    v37 := GenS1{F0: 255, F1: 128, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
}
