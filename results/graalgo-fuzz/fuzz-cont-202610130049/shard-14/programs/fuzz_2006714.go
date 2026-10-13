package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string; F2 int }
func genFn0(p0 uint8) uint16 {
    switch "" {
    case "a":
        v6 := GenS1{F0: 935}
        _ = v6
    default:
        var v7 bool = false
        _ = v7
    }
    var v8 uint16 = 2
    switch ("" + ("" + "\n\t")) {
    case "foo bar":
        v9 := GenS1{F0: 128}
        _ = v9
    case "\n\t":
        v8 -= 945
    }
    v8 = (^(v8))
    _ = p0
    _ = v8
    return 32767
}
func genFn1() string {
    var v10 bool = false
    v11 := []int{456, 680}
    _ = v10
    _ = v11
    return "a"
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 128
    var v3 bool = false
    var v4 int = 2147483647
    v5 := make(map[int]bool)
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 bool = ((false || ((((3.14159 * 3.14159) - (0.0 * -2.25)) / 1.0) != -2.25)) && false)
        switch "a" {
        case "a":
            var v14 uint64 = uint64(v2)
            _ = v14
        }
        v5[(^(v4))] = v13
        _ = fi12
        _ = v13
    }
    fmt.Println(v3, v3, v4)
    v2 -= 1
    var v15 string = (("hello" + ("foo bar" + "hello")) + "foo bar")
    fmt.Println(v3)
    v4 -= v4
    v16 := GenS2{F0: (v4 << ((441) & 63)), F1: v15, F2: (v4 | v4)}
    fmt.Println(v2, v2)
    v3 = false
    v1 = append(v1, 4294967295)
    v3 = v3
    var v17 uint16 = 1
    var v18 uint = 255
    v19 := GenS2{F0: v4, F1: v15, F2: (v4 >> ((v18) & 63))}
    var v20 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
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
    _ = v5
}
