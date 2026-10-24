package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0() uint16 {
    switch 181 {
    case 32768:
        var v5 int8 = -90
        _ = v5
    case 9223372036854775807:
        var v6 uint16 = 313
        _ = v6
    case 127:
        v7 := []int{42, 9223372036854775807}
        _ = v7
    }
    var v8 int8 = 42
    _ = v8
    return 70
}
func genFn1() uint16 {
    v9 := GenS1{F0: 42, F1: ((("hello" + "foo bar") + "\n\t") + "hello")}
    var v10 int = 42
    switch (("\n\t" + "foo bar") + "hello") {
    case "foo bar":
        v11 := []int{-2147483648}
        _ = v11
    case "hello":
        fmt.Println(v10)
    case "\n\t":
        fmt.Println(v10)
    default:
        var v12 float64 = 1e10
        _ = v12
    }
    switch "a" {
    case "hello":
        var v13 int = v10
        _ = v13
    default:
        v14 := []int{v10, v10, v10, v10}
        _ = v14
    }
    _ = v10
    _ = v9
    return 255
}

func main() {
    var v1 int16 = -314
    var v2 uint8 = 127
    var v3 uint8 = v2
    var v4 float64 = 1e10
    var v15 float32 = 0.1
    v4 -= (0.1 * v4)
    var v16 uint8 = v2
    v17 := []int{}
    var v18 bool = true
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println(v15)
    v19 := GenS1{F0: 40, F1: ("\n\t" + "\n\t")}
    var v20 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
