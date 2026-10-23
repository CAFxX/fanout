package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 string }
func genFn0() uint64 {
    var v5 string = ("hello" + "\n\t")
    v5 = v5
    _ = v5
    return 2
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 646
    v3 := []uint32{}
    var v4 string = "a"
    switch "a" {
    case "foo bar":
        v1[-2147483648] = int(v2)
    case "hello":
        v6 := GenS1{F0: 1}
        _ = v6
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v7 int8 = 99
        _ = v7
    }
    v8 := GenS1{F0: 18446744073709551615}
    if false {
        var v9 float64 = 3.14159
        v10 := GenS2{F0: 2147483647, F1: v4}
        v2 -= v2
        _ = v10
        _ = v9
    }
    if (!false) {
        var v11 int = 9223372036854775807
        _ = v11
    }
    var v12 int16 = (v2 - (v2 >> ((927) & 63)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint = 4294967295
    var v14 uint16 = uint16(v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
