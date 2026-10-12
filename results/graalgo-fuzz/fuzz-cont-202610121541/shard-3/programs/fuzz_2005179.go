package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0() int32 {
    switch "a" {
    case "foo bar":
        if ((70 >= -1) || false) {
            v7 := []int{106, 128, 966, -2147483648}
            _ = v7
        }
    case "\n\t":
        var v8 string = ("a" + "")
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 0}
    v10 := GenS1{F0: 643, F1: 373}
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 1
    var v3 int32 = 1
    var v4 string = "foo bar"
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    v5[42] = 1
    if true {
        if true {
            v11 := []int{2147483647, -764, 32768, 42}
            _ = v11
        }
    }
    var v12 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: -9223372036854775808, F1: -2147483648}
        var v15 uint64 = 4294967295
        var v16 uint64 = 2
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    fmt.Println(v4)
    fmt.Println(v3)
    fmt.Println(v4, v3, v3)
    fmt.Println(v3, v1, v3)
    v2 = v2
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
