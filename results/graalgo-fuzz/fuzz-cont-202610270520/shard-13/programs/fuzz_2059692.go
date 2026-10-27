package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
func genFn0() string {
    var v7 float64 = float64((float32((0.1 * -2.25)) + 3.14159))
    if (!(!true)) {
        switch 128 {
        case 4294967295:
            fmt.Println(v7)
        default:
            v7 = v7
        }
    }
    if (-512 < -9223372036854775808) {
        var v8 bool = (208 < 255)
        _ = v8
    }
    _ = v7
    return "foo bar"
}

func main() {
    var v1 bool = false
    v2 := []string{"a", ("a" + "")}
    v3 := []int{4294967295, 701}
    var v4 int32 = 42
    v5 := []string{("\n\t" + "a"), ("foo bar" + "hello")}
    v6 := []uint32{}
    v9 := GenS1{F0: -1, F1: -546}
    switch "a" {
    case "a":
        var v10 uint8 = 2
        _ = v10
    case "hello":
        var v11 string = ("" + "\n\t")
        _ = v11
    case "foo bar":
        _ = len(v3)
    }
    v1 = v1
    var v12 int32 = 239
    v1 = false
    fmt.Println(v12, v1)
    var v13 string = "a"
    v3 = append(v3, 1)
    v14 := GenS1{F0: 769, F1: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
