package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    switch ((("foo bar" + ("hello" + ("a" + (("" + ("a" + "\n\t")) + "hello")))) + "hello") + "") {
    case "":
        v4 := []int{128, 2147483647, -1}
        _ = v4
    case "hello":
        if false {
            fmt.Println(v1)
        }
    }
    if false {
        v1 = (v1 - v1)
    }
    v1 -= 255
    if false {
        var v5 bool = ((-2.25 - float64(((((3.14159 / 1e10) - 1e10) - 0.0) * (((-2.25 - ((0.1 - 0.0) - float32(1e10))) * float32((((-2.25 / (3.14159 - 1.5)) + -2.25) - 0.0))) + ((-2.25 / 1.5) / 1.0))))) <= float64(1.5))
        v6 := []int{1}
        _ = v5
        _ = v6
    }
    fmt.Println(v1)
    v7 := GenS1{F0: 255}
    v8 := GenS1{F0: 18446744073709551615}
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v1, v1)
        var v10 int32 = 2147483647
        var v11 uint = 225
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "\n\t" {
    case "\n\t":
        v1 = (v1 % ((v1) | 1))
    }
    v2[42] = ("a" + ("foo bar" + "\n\t"))
    if true {
        fmt.Println(v1)
    } else {
        v1 -= v1
    }
    fmt.Println(v1)
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v1, fi13, v1)
            var v14 int16 = -203
            fmt.Println(v14, fi13, fi13)
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: 18446744073709551615}
            v1 = v1
            var v17 bool = false
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := GenS2{F0: 9223372036854775807, F1: 107}
        _ = fi12
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
