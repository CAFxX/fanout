package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 bool }
func genFn0(p0 uint8, p1 int, p2 uint64) uint8 {
    v6 := GenS1{F0: true, F1: 32767}
    var v7 int8 = -1
    v8 := []int{int(v7), -963}
    for fi9 := 0; fi9 < 15; fi9++ {
        p2 = (p2 + (p2 << ((7) & 63)))
        v10 := GenS1{F0: false, F1: 42}
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "foo bar":
        v11 := GenS1{F0: false, F1: -32768}
        _ = v11
    case "":
        p2 = 1
    case "\n\t":
        p1 = (^(p1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p0 ^ p0)
}
func genFn1(p0 bool, p1 int32) uint {
    p1 = (p1 * p1)
    p1 = 42
    v12 := GenS2{F0: 2, F1: (0 <= 0), F2: (p0 && p0)}
    var v13 uint = 65535
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return v13
}

func main() {
    var v1 uint = 437
    var v2 uint16 = 42
    var v3 int32 = 42
    var v4 uint16 = v2
    var v5 int32 = 258
    v14 := []int{}
    var v15 float64 = -2.25
    fmt.Println(v3)
    var v16 float32 = (0.0 / (0.1 - 1.5))
    v14 = append(v14, -2147483648)
    v17 := GenS1{F0: (!(32767 > -1)), F1: int16(v1)}
    for fi18 := 0; fi18 < 20; fi18++ {
        v3 -= 0
        v5 = v5
        var v19 float64 = -2.25
        _ = fi18
        _ = v19
    }
    var v20 uint8 = 91
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v20 -= (v20 ^ 128)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
