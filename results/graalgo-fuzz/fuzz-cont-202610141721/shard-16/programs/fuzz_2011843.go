package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int16; F2 uint }
func genFn0() int {
    v4 := GenS1{F0: 435}
    if false {
        var v5 int8 = 127
        _ = v5
    }
    switch 0 {
    case 1:
        var v6 float64 = float64(0.1)
        _ = v6
    default:
        var v7 uint64 = 367
        _ = v7
    }
    var v8 float64 = 3.14159
    _ = v4
    _ = v8
    return 317
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -485
    v3 := []int{9223372036854775807, -2147483648}
    _ = len(v1)
    v9 := GenS2{F0: 398, F1: (v2 - v2), F2: 0}
    switch ("\n\t" + ("a" + ((("foo bar" + "") + (("a" + "") + ("" + (("" + "\n\t") + ("hello" + (("\n\t" + ("hello" + "a")) + "\n\t")))))) + "foo bar"))) {
    case "foo bar":
        switch "" {
        case "":
            fmt.Println(v2)
        case "\n\t":
            fmt.Println(v2)
        }
    case "hello":
        var v10 int = 42
        _ = v10
    }
    if false {
        v2 -= (v2 >> ((1) & 63))
        v3 = append(v3, 9223372036854775807)
        var v11 int16 = (-(v2))
        _ = v11
    }
    v12 := GenS1{F0: 18446744073709551615}
    var v13 uint16 = 32767
    fmt.Println(v13, v13)
    var v14 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
