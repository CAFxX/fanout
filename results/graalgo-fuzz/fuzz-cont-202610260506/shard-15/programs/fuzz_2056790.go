package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }
func genFn0() string {
    var v5 bool = (!(!false))
    if v5 {
        v6 := []int{-945}
        v7 := GenS1{F0: -1, F1: 42, F2: 42}
        _ = v6
        _ = v7
    }
    switch ("a" + ((("a" + ("foo bar" + "a")) + "hello") + "")) {
    case "\n\t":
        if true {
            v5 = (!(1 >= -1))
            v5 = (-32768 < 32767)
            v8 := []int{}
            _ = v8
        }
    default:
        v9 := []int{}
        _ = v9
    }
    _ = v5
    return "hello"
}

func main() {
    v1 := []string{"a"}
    v2 := []int{-289}
    v3 := make(map[string]bool)
    var v4 int64 = 1
    v4 += v4
    var v10 string = "foo bar"
    v10 = v10
    v4 = (v4 << ((122) & 63))
    v11 := GenS1{F0: 797, F1: -1, F2: v4}
    v10 = v10
    v10 = v10
    var v12 int64 = (v4 << ((18446744073709551615) & 63))
    var v13 uint64 = 709
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(v4)
        _ = len(v3)
        _ = fi14
    }
    var v15 int64 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
