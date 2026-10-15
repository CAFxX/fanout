package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 int32 }
func genFn0() int {
    if (65535 < 0) {
        v5 := GenS1{F0: 2147483647, F1: -956, F2: 976}
        v6 := GenS1{F0: 2147483647, F1: 32767, F2: 618}
        var v7 int32 = -728
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 uint = 42
        v9 := []int{-631, -9223372036854775808, -437}
        var v10 int = 9223372036854775807
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int64 = 1
    v12 := GenS1{F0: 42, F1: 944, F2: 255}
    var v13 int8 = -1
    _ = v11
    _ = v12
    _ = v13
    return -2147483648
}

func main() {
    var v1 float64 = 0.1
    var v2 float64 = 0.1
    var v3 bool = true
    var v4 uint32 = 2147483648
    var v14 string = "\n\t"
    fmt.Println(v1, v4, v3)
    switch "foo bar" {
    case "a":
        v1 = 1e10
    }
    var v15 int8 = 127
    v16 := []int{32768, 32767, 4294967295}
    var v17 bool = v3
    v17 = v17
    var v18 bool = false
    v19 := GenS2{F0: 1}
    var v20 uint32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
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
