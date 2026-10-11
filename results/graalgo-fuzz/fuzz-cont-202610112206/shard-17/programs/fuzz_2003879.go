package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 42
    var v3 int32 = 1
    v4 := GenS1{F0: (!("foo bar" >= ""))}
    v1 = 1e10
    if true {
        if true {
            var v5 uint16 = 0
            fmt.Println(v3)
            _ = v5
        }
        for fi6 := 0; fi6 < 20; fi6++ {
            fmt.Println(v3)
            v7 := []int{fi6, 32767}
            _ = fi6
            _ = v7
        }
        v8 := []int{32768, 0}
        _ = v8
    }
    switch "a" {
    case "hello":
        v9 := GenS1{F0: false}
        _ = v9
    case "":
        v10 := []int{-9223372036854775808, 0, 2147483647}
        _ = v10
    case "\n\t":
        v2 = (v2 << ((42) & 63))
    }
    v11 := GenS2{F0: -1, F1: -9223372036854775808}
    var v12 uint8 = v2
    v3 -= v3
    v1 = ((((float64(0.0) - ((0.1 / -2.25) - 0.1)) / -2.25) * 0.0) + float64(-2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
