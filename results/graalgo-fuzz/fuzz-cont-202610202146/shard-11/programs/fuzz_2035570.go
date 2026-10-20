package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0() int32 {
    switch ("a" + "hello") {
    case "\n\t":
        v6 := GenS1{F0: 4294967295, F1: 4294967295}
        _ = v6
    default:
        var v7 uint = 255
        _ = v7
    }
    v8 := []int{32768, 310}
    _ = v8
    return 42
}

func main() {
    v1 := []int{405, -442}
    var v2 int8 = 1
    var v3 float64 = float64((0.0 - 3.14159))
    v4 := []uint32{0}
    v5 := []int{2147483647}
    v9 := GenS2{F0: -353, F1: 127}
    switch "" {
    case "":
        v2 -= v2
    case "a":
        v2 = v2
    }
    v3 += float64(0.0)
    if true {
        v3 = v3
    } else {
        switch ("foo bar" + ("foo bar" + ("hello" + (("hello" + "a") + "foo bar")))) {
        case "a":
            var v10 uint = 18446744073709551615
            _ = v10
        }
    }
    if false {
        fmt.Println(v3, v3)
    }
    var v11 int16 = 0
    var v12 bool = true
    var v13 float64 = v3
    v14 := GenS1{F0: 2, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
