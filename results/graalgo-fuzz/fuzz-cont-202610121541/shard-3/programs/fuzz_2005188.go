package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int8; F2 uint64 }
func genFn0(p0 uint8) int64 {
    v7 := []int{}
    v8 := GenS1{F0: 2147483647}
    _ = len(v7)
    for fi9 := 0; fi9 < 2; fi9++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    v3 := make(map[int]int)
    var v4 uint16 = 255
    var v5 int32 = 1
    var v6 uint8 = v2
    switch ("foo bar" + ("\n\t" + "a")) {
    case "a":
        v10 := GenS1{F0: 0}
        _ = v10
    case "\n\t":
        var v11 int = -979
        _ = v11
    case "hello":
        v4 += v4
    }
    v4 += 0
    v12 := GenS1{F0: -9223372036854775808}
    v5 = (v5 + -2147483648)
    switch "hello" {
    case "hello":
        v5 = -196
    }
    var v13 float64 = 1e10
    v14 := GenS2{F0: -777, F1: -33, F2: 42}
    fmt.Println(v2, v4, v5)
    v15 := []int{128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
