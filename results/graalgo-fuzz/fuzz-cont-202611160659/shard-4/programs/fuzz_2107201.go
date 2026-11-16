package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
func genFn0() uint8 {
    v6 := GenS1{F0: 9223372036854775807, F1: 1}
    var v7 string = ("foo bar" + "foo bar")
    var v8 uint64 = 700
    _ = v6
    _ = v7
    _ = v8
    return 180
}

func main() {
    v1 := []string{}
    var v2 int64 = 9223372036854775807
    v3 := []string{"\n\t", ((("hello" + "a") + (((("hello" + "foo bar") + "") + "a") + "")) + ("hello" + ("foo bar" + ""))), "", "foo bar"}
    v4 := []uint32{1}
    v5 := []int{-2147483648, 32767, 32768}
    v9 := GenS1{F0: v2, F1: 42}
    switch "" {
    case "a":
        v2 -= (v2 % (((v2 / (((-(v2))) | 1))) | 1))
    case "hello":
        var v10 int16 = 1
        _ = v10
    case "":
        var v11 int32 = 42
        _ = v11
    }
    v12 := GenS1{F0: (^(v2)), F1: 42}
    var v13 float64 = 3.14159
    v14 := GenS2{F0: false, F1: 127, F2: 42}
    var v15 int = -2147483648
    v16 := GenS1{F0: 1, F1: 976}
    v17 := GenS2{F0: (!false), F1: 0, F2: 556}
    v13 += float64((0.1 + 0.0))
    v15 += v15
    v13 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
